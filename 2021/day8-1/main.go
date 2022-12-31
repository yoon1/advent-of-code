package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ErrOpenFile    = "Error file open!!"
	ErrReadFile    = "Error file read!!"
	ErrInvalidData = "Error invalid data!!"
)

func readFile(fileName string) ([]string, error) {
	//open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("%s: %s", ErrOpenFile, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// read line by line
	var lines []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	if err := fileScanner.Err(); err != nil {
		return nil, errors.New(ErrReadFile)
	}

	return lines, nil
}

func contains(arr []int, t int) bool {
	for _, v := range arr {
		if v == t {
			return true
		}
	}

	return false
}

func solve(strs []string) int {
	sum := 0
	targets := []int{2, 3, 4, 7}
	for _, str := range strs {
		s := strings.Split(str, "|")
		digits := strings.Split(s[1], " ")
		for _, digit := range digits {
			if contains(targets, len(digit)) {
				sum++
			}
		}
	}
	return sum
}

func main() {
	const (
		fileName = "input2"
	)

	strs, err := readFile(fileName)
	if err != nil {
		log.Fatalf("%s, %s", ErrInvalidData, err)
	}

	// fmt.Printf("%v", strs)
	fmt.Print(solve(strs))
}
