package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) [][]string {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	returnArr := [][]string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		list := []string{line}
		returnArr = append(returnArr, list)
	}

	return returnArr
}

func getAllCoordinatesOfLetter(input [][]string, letter string) []string {
	returnArr := []string{}

	for y, line := range input {
		for x, char := range line[0] {
			if strings.ToUpper(string(char)) == strings.ToUpper(letter) {
				returnArr = append(returnArr, strconv.Itoa(x)+","+strconv.Itoa(y))
			}
		}
	}

	return returnArr
}

func findXmasCount(coordinates string, input [][]string) int {
	xy := strings.Split(coordinates, ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])
	count := 0

	count += checkHorizontal(x, y, input)
	count += checkVertical(x, y, input)
	count += checkDiagonalDown(x, y, input)
	count += checkDiagonalUp(x, y, input)

	return count
}

func checkHorizontal(x int, y int, input [][]string) int {
	stringToCheck := strings.Join(input[y], "")
	count := 0

	if x <= len(stringToCheck)-4 {
		if stringToCheck[x:x+4] == "XMAS" {
			count++
		}
	}

	if x >= 3 {
		if stringToCheck[x-3:x+1] == "SAMX" {
			count++
		}
	}

	return count
}

func checkVertical(x int, y int, input [][]string) int {
	stringToCheck := ""
	count := 0

	if y <= len(input)-4 {
		for i := 0; i < 4; i++ {
			stringToCheck += input[y+i][0][x : x+1]
		}

		if stringToCheck == "XMAS" {
			count++
		}
	}

	if y >= 3 {
		stringToCheck = ""
		for i := 0; i < 4; i++ {
			stringToCheck += input[y-i][0][x : x+1]
		}

		if stringToCheck == "XMAS" {
			count++
		}
	}

	return count
}

func checkDiagonalDown(x int, y int, input [][]string) int {
	count := 0

	if x <= len(input[y][0])-4 && y <= len(input)-4 {
		stringToCheck := ""
		for i := 0; i < 4; i++ {
			stringToCheck += input[y+i][0][x+i : x+i+1]
		}

		if stringToCheck == "XMAS" {
			count++
		}
	}

	if x >= 3 && y <= len(input)-4 {
		stringToCheck := ""
		for i := 0; i < 4; i++ {
			stringToCheck += input[y+i][0][x-i : x-i+1]
		}

		if stringToCheck == "XMAS" {
			count++
		}
	}

	return count
}

func checkDiagonalUp(x int, y int, input [][]string) int {
	count := 0

	if x <= len(input[y][0])-4 && y >= 3 {
		stringToCheck := ""
		for i := 0; i < 4; i++ {
			stringToCheck += input[y-i][0][x+i : x+i+1]
		}

		if stringToCheck == "XMAS" {
			count++
		}
	}

	if x >= 3 && y >= 3 {
		stringToCheck := ""
		for i := 0; i < 4; i++ {
			stringToCheck += input[y-i][0][x-i : x-i+1]
		}

		if stringToCheck == "XMAS" {
			count++
		}
	}

	return count
}

func findMASCount(coordinates string, input [][]string) int {
	xy := strings.Split(coordinates, ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])
	count := 0

	count += findMASInXShape(x, y, input)

	return count
}

func findMASInXShape(x int, y int, input [][]string) int {
	count := 0

	if x == 0 || x == len(input[y][0])-1 || y == 0 || y == len(input)-1 {
		return 0
	}

	if string(input[y-1][0][x-1]) == "S" && string(input[y+1][0][x-1]) == "S" && string(input[y-1][0][x+1]) == "M" && string(input[y+1][0][x+1]) == "M" {
		count++
	}

	if string(input[y-1][0][x-1]) == "M" && string(input[y+1][0][x-1]) == "S" && string(input[y-1][0][x+1]) == "M" && string(input[y+1][0][x+1]) == "S" {
		count++
	}

	if string(input[y-1][0][x-1]) == "M" && string(input[y+1][0][x-1]) == "M" && string(input[y-1][0][x+1]) == "S" && string(input[y+1][0][x+1]) == "S" {
		count++
	}

	if string(input[y-1][0][x-1]) == "S" && string(input[y+1][0][x-1]) == "M" && string(input[y-1][0][x+1]) == "S" && string(input[y+1][0][x+1]) == "M" {
		count++
	}

	return count
}

func main() {
	input := readInput("input.txt")

	// Get X to ensure you can potentially spell XMAS
	coordinates := getAllCoordinatesOfLetter(input, "X")

	count := 0
	for _, coord := range coordinates {
		count += findXmasCount(coord, input)
	}

	fmt.Printf("XMAS count: %d\n", count)

	// Get A to make it easy to find X shapes with A in the middle
	coordinates = getAllCoordinatesOfLetter(input, "A")

	count = 0
	for _, coord := range coordinates {
		count += findMASCount(coord, input)
	}

	fmt.Printf("MAS count: %d\n", count)
}
