package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Robot struct {
	Pos  Point
	Move Point
}

const (
	XLen = 101
	YLen = 103
)

func main() {
	robots := readInput("input.txt")

	points := make(map[Point]int)
	midX := XLen / 2
	midY := YLen / 2

	for _, robot := range robots {
		point := move(robot, 100)

		if point.X != midX && point.Y != midY {
			points[point]++
		}
	}

	quadrant1 := 0
	quadrant2 := 0
	quadrant3 := 0
	quadrant4 := 0

	for point, count := range points {
		if point.X < midX && point.Y < midY {
			quadrant1 += count
		} else if point.X > midX && point.Y < midY {
			quadrant2 += count
		} else if point.X < midX && point.Y > midY {
			quadrant3 += count
		} else if point.X > midX && point.Y > midY {
			quadrant4 += count
		}
	}

	sum := quadrant1 * quadrant2 * quadrant3 * quadrant4

	fmt.Printf("Part 1 Sum: %d\n", sum)

	minSecond, _ := simulate(robots, 10000)
	fmt.Println(minSecond)
}

func simulate(robots []Robot, maxSeconds int) (int, [][]int) {
	minSecond := 0
	var bestGrid [][]int

	grid := make([][]int, YLen)
	yPoints := make(map[int]int)

	for t := 0; t < maxSeconds; t++ {
		for i := range grid {
			grid[i] = make([]int, XLen)
		}

		for i, _ := range robots {
			robots[i].Pos.X = (robots[i].Pos.X + robots[i].Move.X + XLen) % XLen
			robots[i].Pos.Y = (robots[i].Pos.Y + robots[i].Move.Y + YLen) % YLen

			grid[robots[i].Pos.Y][robots[i].Pos.X]++

			yPoints[robots[i].Pos.Y]++
		}

		// Find an example where the coordinates are jumbled and start drawing every 101
		if t >= 1086 && (t-1086)%101 == 0 {
			file, err := os.OpenFile("picture.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			defer file.Close()

			writer := bufio.NewWriter(file)
			fmt.Fprintf(writer, "t = %d\n", t)

			for _, row := range grid {
				for _, cell := range row {
					if cell > 0 {
						fmt.Fprint(writer, "#")
					} else {
						fmt.Fprint(writer, ".")
					}
				}
				fmt.Fprintln(writer)
			}
			fmt.Fprintln(writer)
			writer.Flush()
		}

	}

	return minSecond, bestGrid
}

func move(robot Robot, moves int) Point {
	curPosition := robot.Pos
	for i := 0; i < moves; i++ {
		curPosition.X = (curPosition.X + robot.Move.X + XLen) % XLen
		curPosition.Y = (curPosition.Y + robot.Move.Y + YLen) % YLen
	}

	return curPosition
}

func readInput(filename string) []Robot {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	robots := []Robot{}
	for scanner.Scan() {
		line := scanner.Text()

		strArr := strings.Split(line, "=")
		start := strings.Split(strArr[1], ",")
		xStart, _ := strconv.Atoi(start[0])
		yStart, _ := strconv.Atoi(strings.Split(start[1], " ")[0])
		point := Point{xStart, yStart}

		move := strings.Split(strArr[2], ",")
		xMove, _ := strconv.Atoi(move[0])
		yMove, _ := strconv.Atoi(move[1])
		movePoint := Point{xMove, yMove}

		robot := Robot{point, movePoint}
		robots = append(robots, robot)
	}

	return robots
}
