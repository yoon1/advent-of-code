package main

import (
	"bytes"
	"errors"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
)

const (
	ErrOpenFile    = "Error file open!!"
	ErrReadFile    = "Error file read!!"
	ErrInvalidData = "Error invalid data!!"
)

func readNumsInFile(fileName string) ([][]int, error) {
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
 		} else if(bytes.Compare(bufData, []byte{13}) == 0) {
			 nums = append(nums, line)
			 line = []int{}
		}
	}

	return nums, nil
}

func epsilonRate(gammaRates [][]int) (result []int) {
	rowSize := len(gammaRates)
	if rowSize == 0 {
		return result
	}

	colSize := len(gammaRates[0])
	result = make([]int, colSize)
	rateArr := make([]int, colSize)
	wg := sync.WaitGroup{}
	for col := 0; col < colSize; col++ {
		wg.Add(1)
		go func(col int) {
			defer wg.Done()
			sum := 0
			for row := 0; row < rowSize; row++ {
				sum += gammaRates[row][col]
			}
			rateArr[col] = sum
		}(col)
	}

	middle := rowSize/2
	for idx, rate := range rateArr {
		if rate <= middle {
			result[idx] = 1
		}
	}

	return result
}

func reverseBits(bits []int) (result []int)  {
	result = make([]int, len(bits))
	for idx, bit := range bits {
		result[idx] = 1 - bit
	}

	return result
}

func binaryToDecimal(binary []int) int {
	len := len(binary)
	decimal := 0
	for idx, b := range binary {
		decimal += b * int(math.Pow(float64(2), float64(len - idx - 1)))
	}

	return decimal
}

func main() {
	const fileName = "input"

	gammaRates, err := readNumsInFile(fileName)
	if err != nil {
		log.Fatalf("%s, %s", ErrInvalidData, err)
	}

	epsilon := epsilonRate(gammaRates)
	gamma := reverseBits(epsilon)

	log.Printf("%d", binaryToDecimal(epsilon) * binaryToDecimal(gamma))
}