package main

import (
	"bytes"
	"errors"
	queue "github.com/Workiva/go-datastructures/queue"
	"io"
	"log"
	"os"
	"sort"
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

type Square struct {
	val     int
	checked bool
}

var board [][]Square

func readNumsInFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return errors.New(ErrOpenFile)
	}
	defer file.Close()

	buf := make([]byte, 1)

	var line []Square
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.New(ErrReadFile)
		}

		bufData := buf[:n]
		if data, err := strconv.Atoi(string(bufData)); err == nil {
			square := Square{
				val:     data,
				checked: false,
			}
			line = append(line, square)
		} else if bytes.Compare(bufData, []byte{13}) == 0 {
			board = append(board, line)
			line = []Square{}
		}
	}
	board = append(board, line)

	return nil
}

func bfs(start Point, yLen, xLen int) int {
	var queue = queue.New(100)
	queue.Put(start)
	maxSize := 0
	for {
		if queue.Empty() {
			return maxSize
		}
		icur, _ := queue.Get(1)
		cur := icur[0].(Point)
		maxSize++
		for _, dir := range dirs {
			ny := cur.y + dir.y
			nx := cur.x + dir.x
			// 좌표 확인
			if !(ny >= 0 && ny < yLen && nx >= 0 && nx < xLen) {
				continue
			}

			// 체크 확인
			if board[ny][nx].checked {
				continue
			}

			// 9확인
			if board[ny][nx].val < 9 {
				queue.Put(Point{y: ny, x: nx})
				board[ny][nx].checked = true
			}
		}
	}

	return 0
}

func solve() int {
	yLen := len(board)
	xLen := len(board[0])
	var result []int

	for row := 0; row < yLen; row++ {
		for col := 0; col < xLen; col++ {
			if board[row][col].checked {
				continue
			}
			if board[row][col].val == 9 {
				continue
			}
			board[row][col].checked = true
			curResult := bfs(
				Point{
					y: row,
					x: col,
				},
				yLen,
				xLen,
			)
			result = append(result, curResult)
		}
	}

	// sort result
	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	return result[0] * result[1] * result[2]
}

func main() {
	const fileName = "input2"

	err := readNumsInFile(fileName)
	if err != nil {
		log.Fatalf("%s, %s", ErrInvalidData, err)
	}

	result := solve()
	log.Println(result)
}
