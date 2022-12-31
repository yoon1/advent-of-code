package main

import (
	"bufio"
	"errors"
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

const (
	boardSize = 5
)

type point struct {
	row int
	col int
}

type Square struct {
	num     int
	checked bool
}
type Board struct {
	square [][]*Square
}
type Game struct {
	order  []int
	boards []*Board
}

func (b *Board) setSquare(nums [][]int) {
	var square [][]*Square
	for row := range nums {
		var rowSquare []*Square
		for col := range nums[row] {
			s := &Square{
				num:     nums[row][col],
				checked: false,
			}
			rowSquare = append(rowSquare, s)
		}
		square = append(square, rowSquare)
	}
	b.square = square
}

func convertStringToTwoBytesArray(str string) (result []string) {
	strLength := len(str)
	t := ""
	tCount := 0
	for idx := 0; idx < strLength; idx++ {
		ch := str[idx]
		t = t + string(ch)
		tCount++
		if tCount == 2 {
			result = append(result, strings.Trim(t, " "))
			t = ""
			tCount = 0
			idx++
		}
	}

	return
}

func convertStringArrayToIntArray(strings []string) ([]int, error) {
	nums := make([]int, len(strings))
	var err error
	for idx, str := range strings {
		nums[idx], err = strconv.Atoi(str)
		if err != nil {
			continue
		}
	}
	return nums, nil
}

func readGameInFile(fileName string) (*Game, error) {
	game := &Game{}
	//open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("%s: %s", ErrOpenFile, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// read game order
	if fileScanner.Scan() {
		strNums := strings.Split(fileScanner.Text(), ",")
		game.order, err = convertStringArrayToIntArray(strNums)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New(ErrInvalidData)
	}

	//var lines []string
	var lines [][]int
	board := new(Board)
	for fileScanner.Scan() {
		str := fileScanner.Text()
		if len(str) > 0 {
			text := fileScanner.Text()
			strNums := convertStringToTwoBytesArray(text)
			line, err := convertStringArrayToIntArray(strNums)
			if err != nil {
				return nil, errors.New(ErrInvalidData)
			}
			lines = append(lines, line)
		} else if len(lines) > 0 {
			board.setSquare(lines)
			game.boards = append(game.boards, board)
			board = &Board{}
			lines = [][]int{}
		}
	}
	if err := fileScanner.Err(); err != nil {
		return nil, errors.New(ErrReadFile)
	}

	if len(lines) > 0 {
		board.setSquare(lines)
		game.boards = append(game.boards, board)
		board = &Board{}
		lines = [][]int{}
	}

	return game, nil
}

func (game *Game) check(board *Board, p point) bool {
	// row check
	count := 0
	for col := 0; col < boardSize; col++ {
		if board.square[p.row][col].checked {
			count++
		}
	}
	if count == boardSize {
		return true
	}

	// col check
	count = 0
	for row := 0; row < boardSize; row++ {
		if board.square[row][p.col].checked {
			count++
		}
	}
	if count == boardSize {
		return true
	}

	return false
}

func (game *Game) mark(board *Board, num int) point {
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			if board.square[row][col].num == num {
				board.square[row][col].checked = true
				return point{row, col}
			}
		}
	}

	return point{-1, -1}
}

func (game *Game) calc(board *Board) int {
	score := 0
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			if !board.square[row][col].checked {
				score += board.square[row][col].num
			}
		}
	}

	return score
}

func (game *Game) turn(num int) int {
	score := 0
	none := point{-1, -1}
	for _, board := range game.boards {
		markPoint := game.mark(board, num)
		if markPoint != none {
			if game.check(board, markPoint) {
				boardScore := game.calc(board)
				if boardScore > score {
					score = boardScore
				}
			}
		}
	}

	return score
}

func (game *Game) play() int {
	score := 0
	for _, turnNumber := range game.order {
		score = game.turn(turnNumber)
		if score > 0 {
			return score * turnNumber
		}
	}
	return 0
}

func main() {
	const (
		fileName = "input"
	)
	game, err := readGameInFile(fileName)
	if err != nil {
		log.Fatalf("%s, %s", ErrInvalidData, err)
	}

	result := game.play()
	log.Printf("%d", result)
}
