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
	None = 0
)

func solution(nums []int) int {
	largerCount := 0
	pre := None
	for _, num := range nums {
		if pre == None {
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

func main() {
	//open file
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("%s: %s", ErrOpenFile, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// read line by line
	var measurements []int
	for fileScanner.Scan() {
		num, err := strconv.Atoi(fileScanner.Text())
		measurements = append(measurements, num)
		if err != nil {
			log.Fatalf("%s, %s", ErrInvalidData, err)
		}
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("%s, %s", ErrReadFile, err)
	}

	// result
	fmt.Println(solution(measurements))
}
