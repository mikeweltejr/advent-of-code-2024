package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid, fullGrid, startPos, xLen, yLen := readInputFile("input.txt")

	numVisited, _ := patrolPerimeter(grid, startPos, xLen, yLen)

	fmt.Printf("Part 1 Visited: %d\n", numVisited)

	loops := findLoops(fullGrid, startPos, xLen, yLen)

	fmt.Printf("Part 2 Loops: %d\n", loops)
}

func findLoops(grid map[string]string, startPos string, xLen int, yLen int) int {
	loopCount := 0
	for pos := range grid {
		if pos == startPos || grid[pos] != "." {
			continue
		}

		grid[pos] = "#"

		if patrolPerimeterLookForLoop(grid, startPos, xLen, yLen) {
			loopCount++
		}

		grid[pos] = "."
	}

	return loopCount
}

func patrolPerimeterLookForLoop(grid map[string]string, startPos string, xLen int, yLen int) bool {
	const (
		UP    = 0
		RIGHT = 1
		DOWN  = 2
		LEFT  = 3
	)

	visited := make(map[string]bool)
	direction := UP
	outOfBounds := false

	curPos := startPos

	for !outOfBounds {
		posArr := strings.Split(curPos, ",")
		x, _ := strconv.Atoi(posArr[0])
		y, _ := strconv.Atoi(posArr[1])

		switch direction {
		case UP:
			y--

			if y < 0 {
				outOfBounds = true
				break
			}

			if grid[strconv.Itoa(x)+","+strconv.Itoa(y)] == "#" {

				if visited[curPos+"|"+strconv.Itoa(direction)] {
					return true
				}

				visited[curPos+"|"+strconv.Itoa(direction)] = true

				direction = RIGHT
				y++
			}
			curPos = strconv.Itoa(x) + "," + strconv.Itoa(y)
		case RIGHT:
			x++

			if x > xLen {
				outOfBounds = true
				break
			}

			if grid[strconv.Itoa(x)+","+strconv.Itoa(y)] == "#" {
				if visited[curPos+"|"+strconv.Itoa(direction)] {
					return true
				}

				visited[curPos+"|"+strconv.Itoa(direction)] = true

				direction = DOWN
				x--
			}
			curPos = strconv.Itoa(x) + "," + strconv.Itoa(y)
		case DOWN:
			y++

			if y > yLen {
				outOfBounds = true
				break
			}

			if grid[strconv.Itoa(x)+","+strconv.Itoa(y)] == "#" {
				if visited[curPos+"|"+strconv.Itoa(direction)] {
					return true
				}

				visited[curPos+"|"+strconv.Itoa(direction)] = true

				direction = LEFT
				y--
			}
			curPos = strconv.Itoa(x) + "," + strconv.Itoa(y)
		case LEFT:
			x--

			if x < 0 {
				outOfBounds = true
				break
			}

			if grid[strconv.Itoa(x)+","+strconv.Itoa(y)] == "#" {
				if visited[curPos+"|"+strconv.Itoa(direction)] {
					return true
				}

				visited[curPos+"|"+strconv.Itoa(direction)] = true

				direction = UP
				x++
			}
			curPos = strconv.Itoa(x) + "," + strconv.Itoa(y)
		}

	}

	return !outOfBounds
}

func patrolPerimeter(grid map[string]string, startPos string, xLen int, yLen int) (int, map[string][]int) {
	const (
		UP    = 0
		RIGHT = 1
		DOWN  = 2
		LEFT  = 3
	)

	visited := make(map[string][]int)
	direction := UP
	outOfBounds := false

	curPos := startPos
	for !outOfBounds {
		posArr := strings.Split(curPos, ",")
		x, _ := strconv.Atoi(posArr[0])
		y, _ := strconv.Atoi(posArr[1])

		// Record the current position and direction
		if !contains(visited[curPos], direction) {
			if len(visited[curPos]) == 0 {
				visited[curPos] = []int{}
			}
			visited[curPos] = append(visited[curPos], direction)
		}

		switch direction {
		case UP:
			y--

			if y < 0 {
				outOfBounds = true
				break
			}

			if grid[strconv.Itoa(x)+","+strconv.Itoa(y)] == "#" {
				direction = RIGHT
				y++
			}
			curPos = strconv.Itoa(x) + "," + strconv.Itoa(y)
		case RIGHT:
			x++

			if x > xLen {
				outOfBounds = true
				break
			}

			if grid[strconv.Itoa(x)+","+strconv.Itoa(y)] == "#" {
				direction = DOWN
				x--
			}
			curPos = strconv.Itoa(x) + "," + strconv.Itoa(y)
		case DOWN:
			y++

			if y > yLen {
				outOfBounds = true
				break
			}

			if grid[strconv.Itoa(x)+","+strconv.Itoa(y)] == "#" {
				direction = LEFT
				y--
			}
			curPos = strconv.Itoa(x) + "," + strconv.Itoa(y)
		case LEFT:
			x--

			if x < 0 {
				outOfBounds = true
				break
			}

			if grid[strconv.Itoa(x)+","+strconv.Itoa(y)] == "#" {
				direction = UP
				x++
			}
			curPos = strconv.Itoa(x) + "," + strconv.Itoa(y)
		}
	}

	return len(visited), visited
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func readInputFile(filename string) (map[string]string, map[string]string, string, int, int) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	grid := make(map[string]string)
	fullGrid := make(map[string]string)

	y := 0
	xLen := 0
	startIndex := ""
	for scanner.Scan() {
		line := scanner.Text()
		xLen = len(line) - 1

		for i := 0; i < len(line); i++ {
			fullGrid[strconv.Itoa(i)+","+strconv.Itoa(y)] = string(line[i])
			if string(line[i]) == "#" {
				grid[strconv.Itoa(i)+","+strconv.Itoa(y)] = "#"
			} else if string(line[i]) == "^" {
				startIndex = strconv.Itoa(i) + "," + strconv.Itoa(y)
				grid[strconv.Itoa(i)+","+strconv.Itoa(y)] = "^"
			}
		}
		y++
	}
	yLen := y - 1

	return grid, fullGrid, startIndex, xLen, yLen
}
