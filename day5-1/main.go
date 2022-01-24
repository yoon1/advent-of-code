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

var BOARD_MAX = 0

type Line struct {
	left  *Point
	right *Point
}

type Point struct {
	x int
	y int
}

func NewPoint(location string) (*Point, error) {
	location = strings.Trim(location, " ")
	strs := strings.Split(location, ",")
	xx, err := strconv.Atoi(strs[0])
	if err != nil {
		return nil, err
	}
	yy, err := strconv.Atoi(strs[1])
	if err != nil {
		return nil, err
	}

	newPoint := &Point{
		x: xx,
		y: yy,
	}

	return newPoint, nil
}

func count(board [][]int) int {
	result := 0
	for i := 0; i <= BOARD_MAX; i++ {
		for j := 0; j <= BOARD_MAX; j++ {
			if board[i][j] > 1 {
				result++
			}
		}
	}

	return result
}

func draw(lines []*Line) [][]int {
	board := make([][]int, BOARD_MAX+1)
	for i := 0; i < BOARD_MAX+1; i++ {
		board[i] = make([]int, BOARD_MAX+1)
	}
	for _, line := range lines {
		//fmt.Printf("<%d, %d> => <%d, %d>", line.left.x, line.left.y, line.right.x, line.right.y)
		if line.left.x == line.right.x {
			if line.left.y < line.right.y {
				//fmt.Println(":2")
				for i := line.left.y; i <= line.right.y; i++ {
					//fmt.Println("[%d, %d]", line.left.x, i)
					board[line.left.x][i]++
				}
			} else {
				//fmt.Println(":3")
				for i := line.right.y; i <= line.left.y; i++ {
					//fmt.Println("[%d, %d]", line.left.x, i)
					board[line.left.x][i]++
				}
			}
		}
		if line.left.y == line.right.y {
			if line.left.x < line.right.x {
				//fmt.Println(":5")
				for i := line.left.x; i <= line.right.x; i++ {
					//fmt.Println("[%d, %d]", i, line.left.y)
					board[i][line.left.y]++
				}
			} else {
				//fmt.Println(":6")
				for i := line.right.x; i <= line.left.x; i++ {
					//fmt.Println("[%d, %d]", i, line.left.y)
					board[i][line.left.y]++
				}
			}
		}
		//fmt.Println()
	}

	return board
}

func greater(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func readLinesInFile(fileName string) ([]*Line, error) {
	//open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("%s: %s", ErrOpenFile, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// read line by line
	var lines []*Line
	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		strs := strings.Split(lineText, "->")
		left, err := NewPoint(strs[0])
		if err != nil {
			fmt.Println(err)
		}
		right, err := NewPoint(strs[1])
		if err != nil {
			fmt.Println(err)
		}

		BOARD_MAX = greater(BOARD_MAX, left.x)
		BOARD_MAX = greater(BOARD_MAX, left.y)
		BOARD_MAX = greater(BOARD_MAX, right.x)
		BOARD_MAX = greater(BOARD_MAX, right.y)

		newLine := &Line{
			left:  left,
			right: right,
		}
		lines = append(lines, newLine)
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("%s, %s", ErrReadFile, err)
	}

	return lines, nil
}
func main() {
	const (
		fileName = "input2"
	)

	lines, err := readLinesInFile(fileName)
	if err != nil {
		log.Fatalf("%s, %s", ErrInvalidData, err)
	}

	board := draw(lines)

	log.Printf("%d", count(board))
}
