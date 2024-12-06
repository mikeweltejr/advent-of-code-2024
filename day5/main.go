package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rulesForward, rulesBackward, updates := readInputFile("input.txt")

	sum := 0
	for _, update := range updates {
		sum += calculateUpdates(rulesBackward, update)
	}

	fmt.Println("Part 1 Sum: ", sum)

	sum = 0
	for _, update := range updates {
		sum += fixPageNumbers(rulesForward, update)
	}

	fmt.Println("Part 2 Sum: ", sum)
}

func fixPageNumbers(rulesForward map[int][]int, updates []int) int {
	middleNum := 0
	isOrdered := true
	updateMap := make(map[int]int)

	for i := 0; i < len(updates); i++ {
		updateMap[updates[i]] = i
	}

	for i := 0; i < len(updates); i++ {
		if rulesForward[updates[i]] != nil {
			indexesToSwap := []int{}
			for _, rule := range rulesForward[updates[i]] {
				// If the rule is not in the update list, skip
				if _, exists := updateMap[rule]; !exists {
					continue
				}
				ruleIndex := updateMap[rule]
				// If out of order add to swap list
				if ruleIndex < i {
					indexesToSwap = append(indexesToSwap, ruleIndex)
				}
			}

			// If there are swaps to make, sort and swap descending to put in proper order
			if len(indexesToSwap) > 0 {
				isOrdered = false
				slices.Sort(indexesToSwap)
				curNum := updates[i]
				curIndex := i
				for j := len(indexesToSwap) - 1; j >= 0; j-- {
					swapNum := updates[indexesToSwap[j]]
					updateMap[curNum] = indexesToSwap[j]
					updateMap[swapNum] = curIndex
					updates[curIndex] = updates[indexesToSwap[j]]
					updates[indexesToSwap[j]] = curNum
					curIndex--
				}
			}
		}
	}

	if !isOrdered {
		middleNum = updates[len(updates)/2]
	}

	return middleNum
}

// Part 1 - only return middle number if in order
func calculateUpdates(rulesBackward map[int][]int, updates []int) int {
	middleNum := updates[len(updates)/2]
	updateMap := make(map[int]int)

	for i := 0; i < len(updates); i++ {
		updateMap[updates[i]] = i
	}

	for i := len(updates) - 1; i >= 0; i-- {
		if rulesBackward[updates[i]] != nil {
			for _, rule := range rulesBackward[updates[i]] {
				if updateMap[rule] > i {
					return 0
				}
			}
		}
	}

	return middleNum
}

func readInputFile(filename string) (map[int][]int, map[int][]int, [][]int) {
	var rulesForward = make(map[int][]int)
	var rulesBackward = make(map[int][]int)
	var updates = [][]int{}

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || len(updates) > 0 {
			updates = append(updates, convertToIntArray(line))
		} else {
			rule1, rule2 := parseRules(line)

			if rulesForward[rule1] == nil {
				rulesForward[rule1] = []int{rule2}
			} else {
				rulesForward[rule1] = append(rulesForward[rule1], rule2)
			}

			if rulesBackward[rule2] == nil {
				rulesBackward[rule2] = []int{rule1}
			} else {
				rulesBackward[rule2] = append(rulesBackward[rule2], rule1)
			}
		}
	}

	updates = updates[1:]

	return rulesForward, rulesBackward, updates
}

func parseRules(rule string) (int, int) {
	ruleArr := strings.Split(rule, "|")

	if len(ruleArr) != 2 {
		fmt.Println("Error: Rule must have two values")
		os.Exit(1)
	}

	rule1, _ := strconv.Atoi(ruleArr[0])
	rule2, _ := strconv.Atoi(ruleArr[1])

	return rule1, rule2
}

func convertToIntArray(input string) []int {
	returnArr := []int{}
	strArray := strings.Split(input, ",")

	for _, str := range strArray {
		num, _ := strconv.Atoi(str)
		returnArr = append(returnArr, num)
	}

	return returnArr
}
