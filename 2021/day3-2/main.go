package main

import (
	"bytes"
	"errors"
	"io"
	"log"
	"math"
	"os"
	"strconv"
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
		log.Println(bufData)
		if data, err := strconv.Atoi(string(bufData)); err == nil {
			line = append(line, data)
		} else if bytes.Compare(bufData, []byte{13}) == 0 {
			nums = append(nums, line)
			line = []int{}
		}
	}
	nums = append(nums, line)

	return nums, nil
}

func find(rates [][]int, target, base int) [][]int {
	rowSize := len(rates)
	if rowSize == 1 {
		return rates
	}

	collection := [2][][]int{}

	for row := 0; row < rowSize; row++ {
		x := rates[row][target]
		collection[x] = append(collection[x], rates[row])
	}

	if len(collection[0]) == len(collection[1]) {
		return collection[base]
	}

	if base == 1 {
		if len(collection[0]) > len(collection[1]) {
			return collection[0]
		} else {
			return collection[1]
		}
	} else {
		if len(collection[0]) < len(collection[1]) {
			return collection[0]
		} else {
			return collection[1]
		}
	}
}

func lifeRate(rates [][]int, base int) (result []int) {
	colSize := len(rates[0])
	result = make([]int, colSize)
	for col := 0; col < colSize; col++ {
		rates = find(rates, col, base)
		if base == 0 {
			log.Printf("COL SIZE:: %v", rates)
		}
	}

	return rates[0]
}

func binaryToDecimal(binary []int) int {
	len := len(binary)
	decimal := 0
	for idx, b := range binary {
		decimal += b * int(math.Pow(float64(2), float64(len-idx-1)))
	}

	return decimal
}

func main() {
	const path = "../day11-1/"
	const fileName = "input"

	gammaRates, err := readNumsInFile(path + fileName)
	if err != nil {
		log.Fatalf("%s, %s", ErrInvalidData, err)
	}

	oxygen := lifeRate(gammaRates, 1)
	co2 := lifeRate(gammaRates, 0)
	log.Printf("%d", binaryToDecimal(oxygen)*binaryToDecimal(co2))
}
