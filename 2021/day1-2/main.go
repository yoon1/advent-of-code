package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	ErrOpenFile    = "Error file open!!"
	ErrReadFile    = "Error file read!!"
	ErrInvalidData = "Error invalid data!!"
)

const (
	none = 0
)

func countLargerThanPrev(nums []int) int {
	largerCount := 0
	pre := none
	for _, num := range nums {
		if pre == none {
			pre = num
			continue
		}
		cur := num

		if cur > pre {
			largerCount++
		}

		pre = cur
	}

	return largerCount
}

func sumOfList(nums []int) int {
	result := 0
	for _, val := range nums {
		result += val
	}
	return result
}

func convertListSumOfSlidingWindow(nums []int) []int {
	const slidingWindowSize = 3
	size := len(nums) - slidingWindowSize

	var list []int
	for i := 0; i <= size; i++ {
		var slidingWindow []int
		for slidingIdx := 0; slidingIdx < slidingWindowSize; slidingIdx++ {
			slidingWindow = append(slidingWindow, nums[i+slidingIdx])
		}
		list = append(list, sumOfList(slidingWindow))
	}

	return list
}

func solution(nums []int) int {
	listSumOfSlidingWindow := convertListSumOfSlidingWindow(nums)
	return countLargerThanPrev(listSumOfSlidingWindow)
}

func readNumsInFile(fileName string) ([]int, error) {
	//open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("%s: %s", ErrOpenFile, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// read line by line
	var nums []int
	for fileScanner.Scan() {
		num, err := strconv.Atoi(fileScanner.Text())
		nums = append(nums, num)
		if err != nil {
			return nil, err
		}
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("%s, %s", ErrReadFile, err)
	}

	return nums, nil
}

func main() {
	const fileName = "input"

	measurements, err := readNumsInFile(fileName)
	if err != nil {
		log.Fatalf("%s, %s", ErrInvalidData, err)
	}

	// result
	fmt.Println(solution(measurements))
}
