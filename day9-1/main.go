package main

import (
	"bytes"
	"errors"
	"fmt"
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

type Point struct {
	y int
	x int
}

var dirs = []Point{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

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
		} else if bytes.Compare(bufData, []byte{13}) == 0 {
			nums = append(nums, line)
			line = []int{}
		}
	}
	nums = append(nums, line)

	return nums, nil
}

func check(board [][]int, yLen, xLen, cur, row, col int) bool {
	for _, dir := range dirs {
		if !(row+dir.y >= 0 && row+dir.y < yLen && col+dir.x >= 0 && col+dir.x < xLen) {
			continue
		}
		if board[row+dir.y][col+dir.x] <= cur {
			return false
		}
	}
	return true
}

func find(board [][]int) []Point {
	yLen := len(board)
	xLen := len(board[0])

	points := []Point{}
	for row := 0; row < yLen; row++ {
		for col := 0; col < xLen; col++ {
			if check(board, yLen, xLen, board[row][col], row, col) {
				p := Point{y: row, x: col}
				points = append(points, p)
			}
		}
	}

	return points
}

func aroundMin(board [][]int, yLen, xLen, cur, row, col int) float64 {
	result := math.Inf(0)
	for _, dir := range dirs {
		if !(row+dir.y >= 0 && row+dir.y < yLen && col+dir.x >= 0 && col+dir.x < xLen) {
			continue
		}
		result = math.Min(float64(board[row+dir.y][col+dir.x]), result)
	}

	return result
}

//func calc(board [][]int, points []Point) float64 {
//	result := float64(0)
//	yLen := len(board)
//	xLen := len(board[0])
//
//	for _, p := range points {
//		result += aroundMin(board, yLen, xLen, board[p.y][p.x], p.y, p.x)
//	}
//
//	return result
//}

func calc(board [][]int, points []Point) (result int) {
	for _, point := range points {
		result += board[point.y][point.x] + 1
	}
	return
}

func main() {
	const fileName = "input2"

	board, err := readNumsInFile(fileName)
	if err != nil {
		log.Fatalf("%s, %s", ErrInvalidData, err)
	}

	points := find(board)
	fmt.Printf("%v", points)

	result := calc(board, points)
	fmt.Println(result)
}
