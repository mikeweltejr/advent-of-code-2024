package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack []int

func (s *Stack) Push(element int) {
	*s = append(*s, element)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false // Stack is empty
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

type Disk struct {
	Index int
	Count int
}

func main() {
	diskMap := readInput("input.txt")

	disk, numCount, numStack, freeMap, freeIdx := buildDisk(diskMap)

	newDisk := eliminateFreeSpace(disk)
	checksum := calculateChecksum(newDisk)

	fmt.Printf("Part 1 Checksum: %d\n", checksum)

	disk, numCount, numStack, freeMap, freeIdx = buildDisk(diskMap)
	newDisk = eliminateFreeSpaceExact(disk, numCount, numStack, freeMap, freeIdx)

	checksum = calculateChecksum(newDisk)
	fmt.Println(checksum)
}

func calculateChecksum(disk []string) int {
	sum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == "." {
			continue
		}

		num, _ := strconv.Atoi(disk[i])

		sum += (num * i)
	}

	return sum
}

func eliminateFreeSpaceExact(disk []string, numCount map[int]Disk, numStack Stack, freeMap map[int]int, freeIdx []int) []string {
	for i := 0; i < len(disk); i++ {
		if disk[i] == "." {
			num, _ := numStack.Pop()

			diskInfo := numCount[num]

			for j := 0; j < len(freeIdx); j++ {
				diff := freeMap[freeIdx[j]] - diskInfo.Count

				if freeIdx[j] >= diskInfo.Index {
					break
				}

				if diff >= 0 {
					// replace free space with nums
					for x := 0; x < diskInfo.Count; x++ {
						disk[freeIdx[j]+x] = strconv.Itoa(num)
					}
					// replace num with free space
					for x := 0; x < diskInfo.Count; x++ {
						disk[diskInfo.Index+x] = "."
					}

					// delete free map with this index
					delete(freeMap, freeIdx[j])
					freeIdx[j] = freeIdx[j] + diskInfo.Count

					// Add the new index
					if diff > 0 {
						freeMap[freeIdx[j]] = diff
					}

					break
				}
			}
		} else {
			continue
		}
	}

	return disk
}

func eliminateFreeSpace(disk []string) []string {
	maxIndex := 0
	for i := len(disk) - 1; i >= 0; i-- {
		_, err := strconv.Atoi(disk[i])

		if err == nil {
			maxIndex = i
			break
		}
	}

	for i := 0; i < len(disk); i++ {
		if disk[i] == "." {
			disk[i] = disk[maxIndex]
			disk[maxIndex] = "."

			startIndex := len(disk) - maxIndex
			for j := len(disk) - startIndex - 1; j >= 0; j-- {
				_, err := strconv.Atoi(disk[j])
				if err == nil {
					maxIndex = j
					break
				}
			}
		}

		if maxIndex <= i {
			break
		}

		//fmt.Println(disk)
	}

	return disk
}

func buildDisk(diskMap []int) ([]string, map[int]Disk, Stack, map[int]int, []int) {
	id := 0
	disk := []string{}
	numCount := make(map[int]Disk)
	freeMap := make(map[int]int)
	freeIdx := []int{}
	numStack := Stack{}

	for i := 0; i < len(diskMap); i++ {
		isFree := false
		if i%2 != 0 {
			isFree = true
		}

		count := 0
		for j := 0; j < diskMap[i]; j++ {
			count++
			if isFree {
				disk = append(disk, ".")
			} else {
				disk = append(disk, strconv.Itoa(id))
			}
		}

		if !isFree {
			numStack.Push(id)
			curDisk := Disk{len(disk) - diskMap[i], count}
			numCount[id] = curDisk
			id++
		} else {
			if count > 0 {
				freeMap[len(disk)-diskMap[i]] = count
				freeIdx = append(freeIdx, len(disk)-diskMap[i])
			}
		}
	}

	return disk, numCount, numStack, freeMap, freeIdx
}

func readInput(filename string) []int {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	returnArr := []int{}
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			x, _ := strconv.Atoi(string(line[i]))
			returnArr = append(returnArr, x)
		}
	}

	return returnArr
}
