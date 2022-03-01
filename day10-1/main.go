package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	ErrOpenFile    = "Error file open!!"
	ErrReadFile    = "Error file read!!"
	ErrInvalidData = "Error invalid data!!"
)

var brackets = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var scores = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

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

func calc(str string) int {
	stack := ""
	for _, c := range str {
		cur := string(c)
		if _, exists := brackets[cur]; exists {
			stack += brackets[cur]
		} else {
			if cur == string(stack[len(stack)-1]) {
				stack = stack[:len(stack)-1]
			} else {
				return scores[cur]
			}
		}
	}

	return 0
}

func solve(strings []string) int {
	result := 0
	for _, str := range strings {
		result += calc(str)
	}

	return result
}

func main() {
	const fileName = "input2"

	strings, err := readFile(fileName)
	if err != nil {
		log.Fatalf("%s, %s", ErrInvalidData, err)
	}

	result := solve(strings)
	fmt.Printf("%v", result)
}
