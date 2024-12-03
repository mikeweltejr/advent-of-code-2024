package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func readInputFile(filename string) string {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := ""

	for scanner.Scan() {
		line += scanner.Text()
	}

	return line
}

func getMulString(input string) string {
	// String is valid and should be returned if it is mul(num,num) anything else is invalid
	for i := 0; i < len(input); i++ {
		if string(input[i]) == ")" {
			return input[:i+1]
		}
	}

	return ""
}

func getMulNums(input string) (int, int, bool) {
	// Get the numbers from the string
	nums := []int{}
	num := ""
	openParenChar := "("
	closeParenChar := ")"

	openCharIndex := strings.Index(input, openParenChar)
	closeCharIndex := strings.Index(input, closeParenChar)

	if openCharIndex != 3 || closeCharIndex == -1 {
		return 0, 0, false
	}

	for i := 4; i < len(input); i++ {
		if unicode.IsDigit(rune(input[i])) {
			num += string(input[i])
		} else if string(input[i]) == "," && num != "" {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
			num = ""
		} else if string(input[i]) == ")" && num != "" {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
			break
		} else {
			return 0, 0, false
		}
	}

	if len(nums) != 2 {
		return 0, 0, false
	}

	return nums[0], nums[1], true
}

func getState(input string, isEnabled bool) bool {
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	// Find if do or dont regex is at the 0 position of the string
	doMatchesIndex := doRegex.FindStringIndex(input)
	dontMatchesIndex := dontRegex.FindStringIndex(input)

	// If do matches at index 0 return true if dont matches at index 0 return false
	if doMatchesIndex != nil && doMatchesIndex[0] == 0 {
		return true
	} else if dontMatchesIndex != nil && dontMatchesIndex[0] == 0 {
		return false
	} else {
		return isEnabled
	}
}

func main() {
	input := readInputFile("input.txt")
	sum := 0

	// Part 1
	for i := 0; i < len(input); i++ {
		// If next 3 characters are mul get mul string
		if string(input[i]) == "m" && string(input[i+1]) == "u" && string(input[i+2]) == "l" {
			// Pass string starting from i to getMulString
			mulString := getMulString(input[i:])
			num1, num2, valid := getMulNums(mulString)

			if valid {
				sum += num1 * num2
			}
		}
	}

	fmt.Printf("Part 1 Sum: %d\n", sum)

	// Part 2
	isEnabled := true
	sum = 0
	for i := 0; i < len(input); i++ {
		if string(input[i]) == "d" && string(input[i+1]) == "o" {
			isEnabled = getState(input[i:], isEnabled)
		}

		if string(input[i]) == "m" && string(input[i+1]) == "u" && string(input[i+2]) == "l" {
			// Pass string starting from i to getMulString
			mulString := getMulString(input[i:])
			num1, num2, valid := getMulNums(mulString)

			if valid && isEnabled {
				sum += num1 * num2
			}
		}
	}

	fmt.Printf("Part 2 Sum: %d\n", sum)
}
