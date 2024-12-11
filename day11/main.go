package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stones := readInput("input.txt")
	count := blink(25, stones)
	fmt.Printf("Part 1 Stone Count: %d\n", count)

	count = blink(75, stones)
	fmt.Printf("Part 2 Stone Count: %d\n", count)
}

func blink(count int, stones []int64) int {
	stoneCounts := map[int64]int{}

	for _, stone := range stones {
		stoneCounts[stone]++
	}

	for i := 0; i < count; i++ {
		newStoneCounts := map[int64]int{}

		for stone, freq := range stoneCounts {
			if stone == 0 {
				newStoneCounts[1] += freq
			} else {
				numStr := strconv.FormatInt(stone, 10)
				if len(numStr)%2 == 0 {
					mid := len(numStr) / 2
					leftNum, _ := strconv.ParseInt(numStr[:mid], 10, 64)
					rightNum, _ := strconv.ParseInt(numStr[mid:], 10, 64)
					newStoneCounts[leftNum] += freq
					newStoneCounts[rightNum] += freq
				} else {
					newStoneCounts[stone*2024] += freq
				}
			}
		}

		stoneCounts = newStoneCounts
	}

	totalStones := 0
	for _, freq := range stoneCounts {
		totalStones += freq
	}

	return totalStones
}

func readInput(filename string) []int64 {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	stones := []int64{}

	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Split(line, " ")

		for i := 0; i < len(nums); i++ {
			num, _ := strconv.Atoi(nums[i])
			num64 := int64(num)

			stones = append(stones, num64)
		}
	}

	return stones
}
