package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

func main() {
	robot, boxes, walls, moves, yLen, xLen, grid := readInput("input.txt")

	boxes = move(robot, moves, walls, boxes, xLen, yLen)

	sum := 0
	for box := range boxes {
		sum += 100*box.Y + box.X
	}

	fmt.Printf("Part 1 Sum: %d\n", sum)

	largeBoxes := make(map[Point]Point)
	walls, largeBoxes, robot, xLen, yLen = expandGrid(grid)

	largeBoxes = moveExpanded(robot, moves, walls, largeBoxes, xLen, yLen)

	sum = 0
	counted := make(map[Point]bool)

	for box, pair := range largeBoxes {
		if !counted[box] && !counted[pair] {
			if box.X < pair.X {
				sum += 100*box.Y + box.X
			} else {
				sum += 100*pair.Y + pair.X
			}
			counted[box] = true
			counted[pair] = true
		}
	}

	fmt.Printf("Part 2 Sum: %d\n", sum)

}

func moveExpanded(robot Point, moves []string, walls map[Point]bool, boxes map[Point]Point, xLen int, yLen int) map[Point]Point {
	for i := 0; i < len(moves); i++ {
		delta := getDelta(moves[i])
		nextPos := Point{robot.X + delta.X, robot.Y + delta.Y}

		if walls[nextPos] {
			continue
		}

		if !walls[nextPos] && boxes[nextPos] == (Point{}) {
			robot = nextPos
			continue
		}

		if _, ok := boxes[nextPos]; ok {
			foundEmpty := false
			foundWall := false
			positions := []Point{}
			positions = append(positions, nextPos)

			// Handle horizontal moves
			for !foundEmpty && !foundWall && delta.X != 0 {
				nextPos = Point{nextPos.X + delta.X, nextPos.Y}

				// If wall move nothing
				if walls[nextPos] {
					foundWall = true
					positions = []Point{}
					break
				}

				if boxes[nextPos] != (Point{}) {
					positions = append(positions, nextPos)
					continue
				}

				foundEmpty = true
			}

			if foundEmpty && delta.X != 0 {
				tempBoxes := make(map[Point]Point)
				for i := 0; i < len(positions); i += 2 {
					b := positions[i]
					curBox := boxes[b]

					tempBoxes[Point{b.X + delta.X, b.Y}] = Point{curBox.X + delta.X, curBox.Y}
					tempBoxes[Point{curBox.X + delta.X, curBox.Y}] = Point{b.X + delta.X, b.Y}

					delete(boxes, b)
					delete(boxes, curBox)
				}

				//fmt.Println(tempBoxes)

				for b, p := range tempBoxes {
					boxes[b] = p
					boxes[p] = b
				}

				//fmt.Println(boxes)

				robot.X = positions[0].X
				robot.Y = positions[0].Y
			}

			// Handle Vertical
			foundEmpty = false
			foundWall = false
			foundBoxes := []Point{nextPos, boxes[nextPos]}
			for !foundEmpty && !foundWall && delta.Y != 0 {
				if walls[nextPos] {
					foundWall = true
					break
				}
				if boxes[nextPos] == (Point{}) {
					foundEmpty = true
					break
				}

				var newBoxes []Point
				newBoxes, foundWall = checkForBox(boxes, walls, nextPos, boxes[nextPos], delta.Y)
				if foundWall {
					foundBoxes = []Point{}
					break
				}
				foundBoxes = append(foundBoxes, newBoxes...)
				nextPos = Point{nextPos.X, nextPos.Y + delta.Y}
			}

			if len(foundBoxes) > 0 && delta.Y != 0 && !foundWall {
				tempBoxes := make(map[Point]Point)
				for i := len(foundBoxes) - 1; i >= 0; i-- {
					b := foundBoxes[i]
					if b == (Point{}) || boxes[b] == (Point{}) {
						continue
					}

					curBox := boxes[b]
					tempBoxes[Point{b.X, b.Y + delta.Y}] = Point{curBox.X, curBox.Y + delta.Y}
					tempBoxes[Point{curBox.X, curBox.Y + delta.Y}] = Point{b.X, b.Y + delta.Y}

					delete(boxes, curBox)
					delete(boxes, b)
				}

				for b, p := range tempBoxes {
					boxes[b] = p
					boxes[p] = b
				}
				robot.Y = robot.Y + delta.Y
			}
		}
	}

	return boxes
}

func checkForBox(boxes map[Point]Point, walls map[Point]bool, box1 Point, box2 Point, deltaY int) ([]Point, bool) {
	retBoxes := []Point{}
	visited := make(map[Point]bool)
	foundWall := false

	var check func(Point) bool
	check = func(box Point) bool {
		if visited[box] {
			return true
		}
		visited[box] = true

		nextPos := Point{box.X, box.Y + deltaY}
		nextPairPos := Point{boxes[box].X, boxes[box].Y + deltaY}

		if walls[nextPos] || walls[nextPairPos] {
			retBoxes = []Point{}
			foundWall = true
			return false
		}

		if boxes[nextPos] != (Point{}) {
			retBoxes = append(retBoxes, nextPos)
			if !check(nextPos) {
				return false
			}
		}

		if boxes[nextPairPos] != (Point{}) {
			retBoxes = append(retBoxes, nextPairPos)
			if !check(nextPairPos) {
				return false
			}
		}

		return true
	}

	if !check(box1) || !check(box2) {
		return []Point{}, true
	}

	return retBoxes, foundWall
}

func move(robot Point, moves []string, walls map[Point]bool, boxes map[Point]bool, xLen int, yLen int) map[Point]bool {
	for i := 0; i < len(moves); i++ {
		delta := getDelta(moves[i])
		nextPos := Point{robot.X + delta.X, robot.Y + delta.Y}

		if walls[nextPos] {
			continue
		}

		if !walls[nextPos] && !boxes[nextPos] {
			robot = nextPos
			continue
		}

		if boxes[nextPos] {
			foundEmpty := false
			foundWall := false
			positions := []Point{}
			positions = append(positions, nextPos)

			for !foundEmpty && !foundWall {
				nextPos = Point{nextPos.X + delta.X, nextPos.Y + delta.Y}

				// If wall move nothing
				if walls[nextPos] {
					foundWall = true
					positions = []Point{}
					break
				}

				if boxes[nextPos] {
					positions = append(positions, nextPos)
					continue
				}

				foundEmpty = true
				break
			}

			if len(positions) > 0 {
				robot.X = positions[0].X
				robot.Y = positions[0].Y
				delete(boxes, positions[0])
				newBox := positions[len(positions)-1]
				newBox.X += delta.X
				newBox.Y += delta.Y
				boxes[newBox] = true
			}
		}
	}

	return boxes
}

func expandGrid(input []string) (map[Point]bool, map[Point]Point, Point, int, int) {
	walls := make(map[Point]bool)
	boxes := make(map[Point]Point)
	var robot Point
	yLen := len(input)
	xLen := len(input[0]) * 2

	for y, row := range input {
		for x, char := range row {
			point := Point{x * 2, y}
			switch char {
			case '#':
				walls[point] = true
				walls[Point{point.X + 1, point.Y}] = true
			case '@':
				robot = point
			case 'O':
				left := point
				right := Point{point.X + 1, point.Y}
				boxes[left] = right
				boxes[right] = left
			}
		}
	}

	return walls, boxes, robot, xLen, yLen
}

func getDelta(direction string) Point {
	switch direction {
	case "^":
		return Point{0, -1}
	case "v":
		return Point{0, 1}
	case "<":
		return Point{-1, 0}
	case ">":
		return Point{1, 0}
	}
	return Point{0, 0}
}

func readInput(filename string) (Point, map[Point]bool, map[Point]bool, []string, int, int, []string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	robot := Point{}
	walls := make(map[Point]bool)
	boxes := make(map[Point]bool)
	isGrid := true
	moves := []string{}
	grid := []string{}

	y := 0
	xLen := 0
	for scanner.Scan() {
		line := scanner.Text()

		if xLen == 0 {
			xLen = len(line)
		}

		if len(line) == 0 {
			isGrid = false
			continue
		}

		if isGrid {
			grid = append(grid, line)
		}

		for i := 0; i < len(line); i++ {
			if isGrid {
				if string(line[i]) == "@" {
					robot = Point{i, y}
				} else if string(line[i]) == "O" {
					boxes[Point{i, y}] = true

				} else if string(line[i]) == "#" {
					walls[Point{i, y}] = true
				}
			} else {
				moves = append(moves, string(line[i]))
			}
		}

		if isGrid {
			y++
		}
	}

	return robot, boxes, walls, moves, y, xLen, grid
}
