package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFileToOrderedList(filename string) ([]string, []string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var firstColumn []string
	var secondColumn []string

	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)

		if len(columns) != 2 {
			continue
		}

		firstColumn = append(firstColumn, columns[0])
		secondColumn = append(secondColumn, columns[1])
	}

	return firstColumn, secondColumn, nil
}

func calculateSimilarityScore(firstColumn []string, secondColumn []string) int {
	list2 := make(map[string]int)
	sum := 0

	for i := 0; i < len(secondColumn); i++ {
		list2[secondColumn[i]]++
	}

	for i := 0; i < len(firstColumn); i++ {
		if list2[firstColumn[i]] != 0 {
			a, aErr := strconv.Atoi(firstColumn[i])
			if aErr != nil {
				fmt.Println("Error parsing integers")
				return 0
			}

			sum += a * list2[firstColumn[i]]
		}
	}

	return sum
}

func main() {
	firstColumn, secondColumn, err := readFileToOrderedList("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	sort.Strings(firstColumn)
	sort.Strings(secondColumn)
	var sum float64

	for i := 0; i < len(firstColumn); i++ {
		a, aErr := strconv.ParseFloat(firstColumn[i], 10)
		b, bErr := strconv.ParseFloat(secondColumn[i], 10)

		if aErr != nil || bErr != nil {
			fmt.Println("Error parsing integers")
			return
		}

		sum += math.Abs(a - b)
	}

	sum2 := calculateSimilarityScore(firstColumn, secondColumn)

	fmt.Printf("Part 1: %.0f\n", sum)
	fmt.Printf("Part 2: %d\n", sum2)
}
