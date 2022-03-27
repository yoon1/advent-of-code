package main

import (
	"advent-code-2021/util"
	st "github.com/Workiva/go-datastructures/queue"
	"log"
	"math"
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

type Node struct {
	Point
	score int
}

func (n Node) Compare(other st.Item) int {
	return n.score - other.(Node).score
}

type Square struct {
	val     int
	checked bool
}

var board [][]Square

var yLen = 0
var xLen = 0

func bfs(start Node) int {
	result := math.MaxInt
	var queue = st.NewPriorityQueue(1, true)

	queue.Put(start)

	for {
		if queue.Empty() {
			return result
		}
		icur, _ := queue.Get(1)
		cur := icur[0].(Node)
		if cur.y == yLen-1 && cur.x == xLen-1 {
			if result > cur.score {
				result = cur.score
			}
		}

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

			board[ny][nx].checked = true
			queue.Put(
				Node{
					Point: Point{y: ny, x: nx},
					score: cur.score + board[ny][nx].val,
				},
			)
		}
	}

	return result
}

func set(nums [][]int) {
	for row := range nums {
		var rowSquare []Square
		for col := range nums[row] {
			s := Square{
				val:     nums[row][col],
				checked: false,
			}
			rowSquare = append(rowSquare, s)
		}
		board = append(board, rowSquare)
	}
}

func expand(nums [][]int, scale int) [][]int {
	newNums := make([][]int, scale*yLen)
	next := func(a, b int) int {
		return (a+b-1)%9 + 1
	}

	for y := 0; y < yLen; y++ {
		newNums[y] = make([]int, scale*xLen)
		for x := 0; x < xLen; x++ {
			for i := 0; i < scale; i++ {
				newNums[y][x+i*xLen] = next(nums[y][x], i)
			}
		}
	}

	for y := 0; y < yLen; y++ {
		for i := 1; i < scale; i++ {
			newNums[y+i*yLen] = make([]int, scale*xLen)
			for x := 0; x < scale*xLen; x++ {
				newNums[y+i*yLen][x] = next(newNums[y][x], i)
			}
		}
	}

	return newNums
}

func solve() {
	start := Node{
		Point: Point{y: 0, x: 0},
		score: 0,
	}

	board[0][0].checked = true
	log.Println(bfs(start))
}

func main() {
	const (
		fileName   = "input15-2"
		expandSize = 5
	)

	nums, err := util.ReadNumsInFile(fileName)
	if err != nil {
		log.Fatalf("%s", err)
	}
	yLen = len(nums)
	if yLen > 0 {
		xLen = len(nums[0])
	}

	// 1
	//set(nums)
	//solve()

	// 2
	nums2 := expand(nums, expandSize)
	xLen *= expandSize
	yLen *= expandSize
	set(nums2)
	solve()
}
