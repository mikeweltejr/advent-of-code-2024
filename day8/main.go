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
	antennas, _, xLen, yLen := readFile("input.txt")

	antinodes := getAntinodes(antennas, xLen, yLen, false)

	fmt.Printf("Part 1 Antinodes: %d\n", len(antinodes))

	antinodes = getAntinodes(antennas, xLen, yLen, true)

	fmt.Printf("Part 2 Antinodes: %d\n", len(antinodes))
}

func getAntinodes(antennas map[string][]Point, xLen int, yLen int, isPart2 bool) []Point {
	antinodes := []Point{}
	addedNodes := make(map[Point]bool)

	for _, antennas := range antennas {
		for i := 0; i < len(antennas); i++ {
			curAntenna := antennas[i]

			if !addedNodes[curAntenna] && isPart2 {
				antinodes = append(antinodes, curAntenna)
				addedNodes[curAntenna] = true
			}

			for j := 0; j < len(antennas); j++ {
				antenna := antennas[j]

				if antenna.X == curAntenna.X && antenna.Y == curAntenna.Y {
					continue
				}

				xDiff := abs(curAntenna.X - antenna.X)
				yDiff := abs(curAntenna.Y - antenna.Y)
				yPoint := 0
				xPoint := 0

				if curAntenna.Y < antenna.Y && curAntenna.X < antenna.X {
					yPoint = curAntenna.Y - yDiff
					xPoint = curAntenna.X - xDiff

					if yPoint < 0 || xPoint < 0 {
						continue
					}

					if !addedNodes[Point{xPoint, yPoint}] {
						antinodes = append(antinodes, Point{xPoint, yPoint})
						addedNodes[Point{xPoint, yPoint}] = true
					}

					if isPart2 {
						newNodes, existingNodes := getAllAntinodes(xPoint, yPoint, xDiff*-1, yDiff*-1, xLen, yLen, addedNodes)
						addedNodes = existingNodes
						antinodes = append(antinodes, newNodes...)
					}
				}

				if curAntenna.Y < antenna.Y && curAntenna.X > antenna.X {
					yPoint = curAntenna.Y - yDiff
					xPoint = curAntenna.X + xDiff

					if yPoint < 0 || xPoint > xLen {
						continue
					}

					if !addedNodes[Point{xPoint, yPoint}] {
						antinodes = append(antinodes, Point{xPoint, yPoint})
						addedNodes[Point{xPoint, yPoint}] = true
					}

					if isPart2 {
						newNodes, existingNodes := getAllAntinodes(xPoint, yPoint, xDiff, yDiff*-1, xLen, yLen, addedNodes)
						addedNodes = existingNodes
						antinodes = append(antinodes, newNodes...)
					}
				}

				if curAntenna.Y > antenna.Y && curAntenna.X > antenna.X {
					yPoint = curAntenna.Y + yDiff
					xPoint = curAntenna.X + xDiff

					if yPoint > yLen || xPoint > xLen {
						continue
					}

					if !addedNodes[Point{xPoint, yPoint}] {
						antinodes = append(antinodes, Point{xPoint, yPoint})
						addedNodes[Point{xPoint, yPoint}] = true
					}

					if isPart2 {
						newNodes, existingNodes := getAllAntinodes(xPoint, yPoint, xDiff, yDiff, xLen, yLen, addedNodes)
						addedNodes = existingNodes
						antinodes = append(antinodes, newNodes...)
					}
				}

				if curAntenna.Y > antenna.Y && curAntenna.X < antenna.X {
					yPoint = curAntenna.Y + yDiff
					xPoint = curAntenna.X - xDiff

					if yPoint > yLen || xPoint < 0 {
						continue
					}

					if !addedNodes[Point{xPoint, yPoint}] {
						antinodes = append(antinodes, Point{xPoint, yPoint})
						addedNodes[Point{xPoint, yPoint}] = true
					}

					if isPart2 {
						newNodes, existingNodes := getAllAntinodes(xPoint, yPoint, xDiff*-1, yDiff, xLen, yLen, addedNodes)
						addedNodes = existingNodes
						antinodes = append(antinodes, newNodes...)
					}
				}
			}
		}
	}

	return antinodes
}

func getAllAntinodes(x int, y int, xDiff int, yDiff int, xLen int, yLen int, existingNodes map[Point]bool) ([]Point, map[Point]bool) {
	points := []Point{}

	for true {
		x += xDiff
		y += yDiff

		if x < 0 || x > xLen || y < 0 || y > yLen {
			break
		}

		if (x <= xLen || x >= 0) && (y <= yLen || y >= 0) {
			if !existingNodes[Point{x, y}] {
				points = append(points, Point{x, y})
				existingNodes[Point{x, y}] = true
			}
		}
	}

	return points, existingNodes
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readFile(filename string) (map[string][]Point, map[Point]string, int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Array of all points
	points := make(map[Point]string)
	antennas := make(map[string][]Point)

	y := 0
	xLen := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if xLen == 0 {
			xLen = len(line) - 1
		}

		for i := 0; i < len(line); i++ {
			str := string(line[i])
			point := Point{i, y}

			if str != "." {
				if antennas[str] == nil {
					antennas[str] = []Point{}
				}

				antennas[str] = append(antennas[str], point)
			}

			points[point] = str
		}
		y++
	}

	return antennas, points, xLen, y - 1
}
