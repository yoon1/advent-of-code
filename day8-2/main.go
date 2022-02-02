package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"sort"
	"strconv"
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

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func solve(strs []string) int {
	sum := 0
	for _, str := range strs {
		s := strings.Split(str, "|")
		digits := strings.Split(strings.TrimRight(s[0], " "), " ")

		m := map[string]int{}
		arr := [10]string{}
		flag := [10]bool{}
		for idx, cur := range digits {
			digit := sortString(cur)

			switch len(digit) {
			case 2:
				m[digit] = 1
				arr[1] = digit
				flag[idx] = true
			case 3:
				m[digit] = 7
				arr[7] = digit
				flag[idx] = true
			case 4:
				m[digit] = 4
				arr[4] = digit
				flag[idx] = true
			case 7:
				m[digit] = 8
				arr[8] = digit
				flag[idx] = true
			}
		}

		for idx, cur := range digits {
			digit := sortString(cur)
			if flag[idx] {
				continue
			}

			// 1. 7을 얼마나 포함하지 않는지?
			sevenCount := 0
			for _, c := range arr[7] {
				if !strings.Contains(digit, string(c)) {
					sevenCount++
				}
			}

			// 2. 4를 얼마나 포함하는지?
			fourCount := 0
			for _, c := range arr[4] {
				if !strings.Contains(digit, string(c)) {
					fourCount++
				}
			}

			switch len(digit) {
			case 5:
				if sevenCount == 0 {
					m[digit] = 3
					arr[3] = digit
					flag[idx] = true
				} else if fourCount == 1 {
					m[digit] = 5
					arr[5] = digit
					flag[idx] = true
				} else {
					m[digit] = 2
					arr[2] = digit
					flag[idx] = true
				}
			case 6:
				if sevenCount == 0 && fourCount == 0 {
					m[digit] = 9
					arr[9] = digit
					flag[idx] = true
				} else if sevenCount == 0 {
					m[digit] = 0
					arr[0] = digit
					flag[idx] = true
				} else {
					m[digit] = 6
					arr[6] = digit
					flag[idx] = true
				}
			}
		}

		// decode
		encodes := strings.Split(s[1], " ")
		result := ""
		for _, cur := range encodes {
			code := sortString(cur)
			result += strconv.Itoa(m[code])
		}
		r, _ := strconv.Atoi(result)
		sum += r
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

	log.Print(solve(strs))
}
