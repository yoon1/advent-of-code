package util

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	ErrOpenFile    = "Error file open!!"
	ErrReadFile    = "Error file read!!"
	ErrInvalidData = "Error invalid data!!"
)

func ReadLinesInFile(fileName string) ([]string, error) {
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
		log.Fatalf("%s, %s", ErrReadFile, err)
	}

	return lines, nil
}

func ReadNumsInFile(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New(ErrOpenFile)
	}
	defer file.Close()

	buf := make([]byte, 1)

	var nums [][]int
	var line []int
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New(ErrReadFile)
		}

		bufData := buf[:n]
		if data, err := strconv.Atoi(string(bufData)); err == nil {
			line = append(line, data)
		} else if bytes.Compare(bufData, []byte{13}) == 0 ||
			bytes.Compare(bufData, []byte{10}) == 0 {
			nums = append(nums, line)
			line = []int{}
		}
	}
	nums = append(nums, line)

	return nums, nil
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("[ERROR] %s", err)
	}

	return i
}

func GreaterInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func PrintMatrix(nums [][]int, rowLen, colLen int) {
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			fmt.Printf("%5d ", nums[i][j])
		}
		fmt.Println()
	}
}

func PrintStringMatrix(array [][]string, rowLen, colLen int) {
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			fmt.Printf("%s ", array[i][j])
		}
		fmt.Println()
	}
}
