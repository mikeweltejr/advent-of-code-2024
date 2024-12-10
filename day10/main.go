package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func main() {
	grid, zeroSlice, xLength, yLength := readInput("input.txt")

	sum, sum2 := traverseTrailHeads(grid, zeroSlice, xLength, yLength)

	fmt.Printf("Part 1 Sum: %d\n", sum)
	fmt.Printf("Part 2 Sum: %d\n", sum2)
}

func traverseTrailHeads(grid map[Point]int, trailHeads []Point, xLength int, yLength int) (int, int) {
	count := 0
	countDistinct := 0
	for _, trailHead := range trailHeads {
		visited := make(map[Point]bool)
		traverse(trailHead, -1, grid, xLength, yLength, &count, visited, true)
		traverse(trailHead, -1, grid, xLength, yLength, &countDistinct, visited, false)
	}

	return count, countDistinct
}

func traverse(curPoint Point, prevNum int, grid map[Point]int, xLength int, yLength int, count *int, visited map[Point]bool, isDistinct bool) {
	if curPoint.X < 0 || curPoint.Y < 0 || curPoint.X >= xLength || curPoint.Y >= yLength {
		return
	}

	if visited[curPoint] && isDistinct {
		return
	}

	if grid[curPoint]-prevNum != 1 && prevNum != -1 {
		return
	}

	if grid[curPoint] == 9 && grid[curPoint]-prevNum == 1 {
		*count++
		if isDistinct {
			visited[curPoint] = true
		}
		return
	}

	visited[curPoint] = true
	num := grid[curPoint]

	// Left
	traverse(Point{curPoint.X - 1, curPoint.Y}, num, grid, xLength, yLength, count, visited, isDistinct)

	// Right
	traverse(Point{curPoint.X + 1, curPoint.Y}, num, grid, xLength, yLength, count, visited, isDistinct)

	// Down
	traverse(Point{curPoint.X, curPoint.Y + 1}, num, grid, xLength, yLength, count, visited, isDistinct)

	// Up
	traverse(Point{curPoint.X, curPoint.Y - 1}, num, grid, xLength, yLength, count, visited, isDistinct)
}

func readInput(filename string) (map[Point]int, []Point, int, int) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	y := 0
	xLength := 0
	grid := make(map[Point]int)
	zeroSlice := []Point{}

	for scanner.Scan() {
		line := scanner.Text()

		if xLength == 0 {
			xLength = len(line)
		}

		for i := 0; i < len(line); i++ {
			curStr := string(line[i])
			num, _ := strconv.Atoi(curStr)

			if num == 0 {
				zeroSlice = append(zeroSlice, Point{i, y})
			} else {
				grid[Point{i, y}] = num
			}
		}
		y++
	}

	return grid, zeroSlice, xLength, y
}
