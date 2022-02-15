package util

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	ErrOpenFile    = "Error file open!!"
	ErrReadFile    = "Error file read!!"
	ErrInvalidData = "Error invalid data!!"
)

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

func PrintMatrix(nums [][]int, rowLen, colLen int) {
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			fmt.Printf("%5d ", nums[i][j])
		}
		fmt.Println()
	}
}
