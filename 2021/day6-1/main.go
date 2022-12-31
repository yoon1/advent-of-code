package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ErrOpenFile    = "Error file open!!"
	ErrReadFile    = "Error file read!!"
	ErrInvalidData = "Error invalid data!!"
)

const (
	NewTime      = 8
	DefaultTime  = 6
	DefaultRound = 18
)

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
		strNums := strings.Split(fileScanner.Text(), ",")
		for _, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("%s, %s", ErrReadFile, err)
	}

	return nums, nil
}

func passTime(nums []int) int {
	for i := 0; i < DefaultRound; i++ {
		for index, num := range nums {
			if num == 0 {
				nums[index] = DefaultTime
				nums = append(nums, NewTime)
				continue
			}
			nums[index]--
		}
		fmt.Println(i+1, ":", nums)
	}

	//fmt.Printf("%v", nums)

	return len(nums)
}

func main() {
	const (
		fileName = "input3"
	)

	lanternfish, err := readNumsInFile(fileName)
	if err != nil {
		log.Fatalf("%s, %s", ErrInvalidData, err)
	}

	fmt.Print(passTime(lanternfish))
}
