package main

import (
	"advent-code-2021/util"
	"fmt"
	"log"
)

type Point struct {
	y int
	x int
}

var dy = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dx = []int{-1, 0, 1, -1, 1, -1, 0, 1}

var rowLen = 0
var colLen = 0

func calc(nums [][]int) int {
	count := 0
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if nums[i][j] == 0 {
				count++
			}
		}
	}
	return count
}

func add(nums *[][]int) {
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			(*nums)[i][j]++
		}
	}
}

func explore(nums *[][]int, row, col int) {
	for i := 0; i < 8; i++ {
		ny := row + dy[i]
		nx := col + dx[i]
		if ny >= 0 && ny < rowLen && nx >= 0 && nx < colLen {
			(*nums)[ny][nx]++
		}
	}
}

func simulation(nums [][]int) [][]int {
	// 1. 각 문어의 에너지 레벨이 1씩 증가한다.
	add(&nums)

	// 2. 그리고 나서, 에너지 레벨이 9 이상인 문어는 모두 번쩍입니다.
	// 이것은 대각선으로 인접한 문어를 포함하여 인접한 모든 문어의 에너지 수준을 1만큼 높인다. (8방)
	// 만약 이것이 문어의 에너지 레벨이 9가 넘도록 만든다면, 문어도 번쩍입니다.
	// 이 과정은 새로운 문어의 에너지 레벨이 9 이상으로 계속 증가하는 한 계속됩니다.
	// (문어는 한 걸음 당 최대 한 번만 깜박일 수 있습니다.)
	// 3. 마지막으로, 이 단계에서 번쩍이는 문어는 모든 에너지를 사용하여 번쩍이기 때문에 에너지 수준이 0으로 설정됩니다.
	explores := []Point{}
	for {
		exploreY := false
		for i := 0; i < rowLen; i++ {
			for j := 0; j < colLen; j++ {
				if nums[i][j] > 9 {
					exploreY = true
					explore(&nums, i, j)
					ex := Point{i, j}
					explores = append(explores, ex)
					nums[i][j] -= 10
				}
			}
		}

		if !exploreY {
			break
		}
	}

	for _, ex := range explores {
		nums[ex.y][ex.x] = 0
	}

	return nums
}

func main() {
	const (
		path     = "/home/yoon/hj/SOURCE/practice/advent-code-2021/day11-1/"
		fileName = "input1"
	)

	nums, err := util.ReadNumsInFile(path + fileName)
	if err != nil {
		log.Fatalf("%s", err)
	}

	rowLen = len(nums)
	if rowLen > 0 {
		colLen = len(nums[0])
	}

	sum := 0
	for i := 0; i < 100; i++ {
		nums = simulation(nums)
		sum += calc(nums)
	}

	util.PrintMatrix(nums, rowLen, colLen)
	fmt.Println(sum)

}
