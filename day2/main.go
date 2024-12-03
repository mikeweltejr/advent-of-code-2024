package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readReportFile(filename string) [][]string {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	returnArr := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		list := strings.Split(line, " ")
		returnArr = append(returnArr, list)
	}

	return returnArr
}

func isSequenceSafe(line []int) bool {
	const (
		notSet = -999
		pos    = 1
		neg    = -1
	)
	diffState := notSet

	for i := 0; i < len(line)-1; i++ {
		diff := line[i] - line[i+1]

		// If equal return false
		if diff == 0 {
			return false
		}

		// Set initial state (pos or neg)
		if diffState == notSet {
			if diff > 0 {
				diffState = neg
			} else if diff < 0 {
				diffState = pos
			}
		}

		// Ensure there hasn't been a change in state if so return false
		if (diffState == neg && diff <= 0) || (diffState == pos && diff >= 0) {
			return false
		}

		// Ensure no differences are outside the range of 3
		if (diffState == neg && diff > 3) || (diffState == pos && diff < -3) {
			return false
		}
	}
	return true
}

func part1(report [][]string) int {
	safeCount := 0

	for _, line := range report {
		intLine := make([]int, len(line))
		for i, s := range line {
			intLine[i], _ = strconv.Atoi(s)
		}

		if isSequenceSafe(intLine) {
			safeCount++
		}
	}

	return safeCount
}

func part2(report [][]string) int {
	safeCount := 0
	const (
		notSet = -999
		pos    = 1
		neg    = -1
	)

	for _, line := range report {
		intLine := make([]int, len(line))
		for i, s := range line {
			intLine[i], _ = strconv.Atoi(s)
		}

		// Check if the sequence is already safe
		if isSequenceSafe(intLine) {
			safeCount++
			continue
		}

		// Check by removing each level
		isSafe := false
		for i := 0; i < len(intLine); i++ {
			modifiedLine := append([]int{}, intLine[:i]...)
			if i+1 < len(intLine) {
				modifiedLine = append(modifiedLine, intLine[i+1:]...)
			}
			if isSequenceSafe(modifiedLine) {
				isSafe = true
				break
			}
		}

		if isSafe {
			fmt.Println(intLine)
			safeCount++
		}
	}

	return safeCount
}

func main() {
	report := readReportFile("input.txt")
	safeCount := part1(report)
	part2SafeCount := part2(report)

	fmt.Printf("Part 1: %d\n", safeCount)
	fmt.Printf("Part 2: %d\n", part2SafeCount)
}
