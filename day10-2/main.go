package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
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
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
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

func incomp(str string) string {
	stack := ""
	for _, c := range str {
		cur := string(c)
		if _, exists := brackets[cur]; exists {
			stack += brackets[cur]
		} else {
			if cur == string(stack[len(stack)-1]) {
				stack = stack[:len(stack)-1]
			} else {
				return ""
			}
		}
	}

	return stack
}

func calc(str string) int {
	result := 0
	len := len(str)
	for i := len - 1; i >= 0; i-- {
		cur := string(str[i])
		result *= 5
		result += scores[cur]
	}

	return result
}

func solve(strings []string) int {
	var result []int
	for _, str := range strings {
		s := incomp(str)
		if len(s) > 0 {
			cur := calc(s)
			result = append(result, cur)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	return result[len(result)/2]
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
