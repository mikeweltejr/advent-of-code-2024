package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	X int
	Y int
}

func main() {
	alphaNumMap, grid, xLen, yLen := readInput("input.txt")
	visited := make(map[Point]bool)
	sum := 0
	part2Sum := 0

	for _, val := range alphaNumMap {
		for _, point := range val {
			if !visited[point] {
				plots := getAllConnectedPoints(grid, point, xLen, yLen)
				for _, p := range plots {
					visited[p] = true
				}

				area := len(plots)
				perimeter := calculatePerimeter(plots, grid, xLen, yLen)
				sides := calculateSides(plots)

				sum += area * perimeter
				part2Sum += area * sides
			}
		}
	}

	fmt.Printf("Part 1 sum: %d\n", sum)
	fmt.Printf("Part 2 sum: %d\n", part2Sum)
}

func calculateSides(plots []Point) int {
	sort.Slice(plots, func(i, j int) bool {
		if plots[i].Y == plots[j].Y {
			return plots[i].X < plots[j].X
		}
		return plots[i].Y < plots[j].Y
	})

	sides := 0

	edges := make(map[string]bool)

	for _, point := range plots {
		edgesAtPoint := 0

		if !hasPoint(plots, point.X, point.Y-1) {
			if !edgeExists(edges, point.X-1, point.Y, "top") {
				addEdge(edges, point.X, point.Y, "top")
				edgesAtPoint++
			} else {
				delete(edges, fmt.Sprintf("%d,%d-%s", point.X-1, point.Y, "top"))
				addEdge(edges, point.X, point.Y, "top")
			}
		}

		if !hasPoint(plots, point.X-1, point.Y) {
			if !edgeExists(edges, point.X, point.Y-1, "left") {
				addEdge(edges, point.X, point.Y, "left")
				edgesAtPoint++
			} else {
				delete(edges, fmt.Sprintf("%d,%d-%s", point.X, point.Y-1, "left"))
				addEdge(edges, point.X, point.Y, "left")
			}
		}

		if !hasPoint(plots, point.X, point.Y+1) {
			if !edgeExists(edges, point.X-1, point.Y, "bottom") {
				addEdge(edges, point.X, point.Y, "bottom")
				edgesAtPoint++
			} else {
				delete(edges, fmt.Sprintf("%d,%d-%s", point.X-1, point.Y, "bottom"))
				addEdge(edges, point.X, point.Y, "bottom")
			}
		}

		if !hasPoint(plots, point.X+1, point.Y) {
			if !edgeExists(edges, point.X, point.Y-1, "right") {
				addEdge(edges, point.X, point.Y, "right")
				edgesAtPoint++
			} else {
				delete(edges, fmt.Sprintf("%d,%d-%s", point.X, point.Y-1, "right"))
				addEdge(edges, point.X, point.Y, "right")
			}
		}

		sides += edgesAtPoint
	}

	return sides
}

func hasPoint(plots []Point, x, y int) bool {
	for _, p := range plots {
		if p.X == x && p.Y == y {
			return true
		}
	}

	return false
}

func edgeExists(edges map[string]bool, x int, y int, direction string) bool {
	edge := fmt.Sprintf("%d,%d-%s", x, y, direction)

	return edges[edge]
}

func addEdge(edges map[string]bool, x int, y int, direction string) {
	edge := fmt.Sprintf("%d,%d-%s", x, y, direction)
	edges[edge] = true
}

func calculatePerimeter(plots []Point, grid map[Point]string, xLen int, yLen int) int {
	perimeter := 0
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, plot := range plots {
		for _, dir := range directions {
			neighbor := Point{plot.X + dir.X, plot.Y + dir.Y}
			if neighbor.X < 0 || neighbor.X > xLen || neighbor.Y < 0 || neighbor.Y > yLen || grid[neighbor] != grid[plot] {
				perimeter++
			}
		}
	}

	return perimeter
}

func getAllConnectedPoints(grid map[Point]string, start Point, xLen int, yLen int) []Point {
	visited := make(map[Point]bool)
	var result []Point
	var dfs func(Point)

	dfs = func(p Point) {
		if p.X < 0 || p.X > xLen || p.Y < 0 || p.Y > yLen {
			return
		}
		if visited[p] {
			return
		}
		if grid[p] != grid[start] {
			return
		}

		visited[p] = true
		result = append(result, p)

		dfs(Point{p.X - 1, p.Y})
		dfs(Point{p.X + 1, p.Y})
		dfs(Point{p.X, p.Y - 1})
		dfs(Point{p.X, p.Y + 1})
	}

	dfs(start)
	return result
}

func readInput(filename string) (map[string][]Point, map[Point]string, int, int) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	y := 0
	xLen := 0
	alphaNumMap := make(map[string][]Point)
	grid := make(map[Point]string)

	for scanner.Scan() {
		line := scanner.Text()

		if xLen == 0 {
			xLen = len(line) - 1
		}

		for i := 0; i < len(line); i++ {
			curStr := string(line[i])
			grid[Point{i, y}] = curStr

			val, exists := alphaNumMap[curStr]

			if !exists || len(val) == 0 {
				alphaNumMap[curStr] = []Point{}
			}

			alphaNumMap[curStr] = append(alphaNumMap[curStr], Point{i, y})
		}
		y++
	}

	return alphaNumMap, grid, xLen, y - 1
}
