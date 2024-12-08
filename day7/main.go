package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	testValues, nums := readInputFile("input.txt")

	validTestValues := []int{}
	validTestValuesPart2 := []int{}

	for i := 0; i < len(testValues); i++ {
		if isMatchingTestValue(testValues[i], nums[i], false) {
			validTestValues = append(validTestValues, testValues[i])
		}

		if isMatchingTestValue(testValues[i], nums[i], true) {
			validTestValuesPart2 = append(validTestValuesPart2, testValues[i])
		}
	}

	sum := 0
	for _, testValue := range validTestValues {
		sum += testValue
	}

	fmt.Printf("Part 1 Sum: %d\n", sum)

	sum = 0
	for _, testValue := range validTestValuesPart2 {
		sum += testValue
	}

	fmt.Printf("Part 2 Sum: %d\n", sum)
}

func isMatchingTestValue(testValue int, nums []int, isJoinOperator bool) bool {
	if len(nums) == 0 {
		return false
	}

	return evaluate(testValue, nums[1:], nums[0], isJoinOperator)
}

func evaluate(testValue int, nums []int, current int, isJoinOperator bool) bool {
	if len(nums) == 0 {
		return current == testValue
	}

	next := nums[0]
	remaining := nums[1:]

	//fmt.Printf("Evaluating %d %d %d\n", testValue, current, next)

	if evaluate(testValue, remaining, current+next, isJoinOperator) {
		return true
	}

	if evaluate(testValue, remaining, current*next, isJoinOperator) {
		return true
	}

	if isJoinOperator {
		strVal := strconv.Itoa(current) + strconv.Itoa(next)
		intVal, _ := strconv.Atoi(strVal)
		if evaluate(testValue, remaining, intVal, isJoinOperator) {
			return true
		}
	}

	return false
}

func readInputFile(filename string) ([]int, [][]int) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	testValues := []int{}
	returnNums := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Split(line, ":")
		testValue, _ := strconv.Atoi(values[0])
		testValues = append(testValues, testValue)
		nums := strings.Split(values[1], " ")
		intNums := []int{}

		for _, num := range nums {
			if string(num) == "" {
				continue
			}

			iNum, _ := strconv.Atoi(num)
			intNums = append(intNums, iNum)
		}

		returnNums = append(returnNums, intNums)
	}

	return testValues, returnNums
}
