package main

import (
	"advent-code-2021/util"
	"log"
)

//해삼이 하나는 동쪽(`>`)으로 이동하고 다른 하나는 남쪽(`v`)
// 각 위치에는 최대 한 개의 해삼이 포함될 수 있다.
// 나머지 위치는 비어있다.
var ySize, xSize int

const (
	EAST  = 0
	SOUTH = 1
)

type point struct {
	cy, cx int
	ny, nx int
	dir    string
}

func compare(a [][]string, b [][]string) bool {
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			if a[y][x] != b[y][x] {
				return false
			}
		}
	}

	return true
}

func dirMove(dir int, m *[][]string) {
	points := []point{}
	switch dir {
	case EAST:
		for y := 0; y < ySize; y++ {
			for x := 0; x < xSize; x++ {
				if (*m)[y][x] == ">" {
					nx := (x + 1) % xSize
					if (*m)[y][nx] == "." {
						points = append(points, point{y, x, y, nx, ">"})
					}
				}
			}
		}

	case SOUTH:
		for x := 0; x < xSize; x++ {
			for y := 0; y < ySize; y++ {
				if (*m)[y][x] == "v" {
					ny := (y + 1) % ySize
					if (*m)[ny][x] == "." {
						points = append(points, point{y, x, ny, x, "v"})
					}
				}
			}
		}
	}

	for _, p := range points {
		(*m)[p.cy][p.cx] = "."
		(*m)[p.ny][p.nx] = p.dir
	}
}

func move(m [][]string) {
	// 동쪽으로 이동
	dirMove(EAST, &m)

	//// 남쪽으로 이동
	dirMove(SOUTH, &m)
}

func copyVal(a [][]string, b [][]string) {
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			a[y][x] = b[y][x]
		}
	}
}

func main() {
	const (
		fileName0 = "input25-0"
		fileName1 = "input25-1"
		fileName2 = "input25-2"
		fileName3 = "input25-3"
	)

	pre, err := util.ReadCharsInFile(fileName3, []string{"v", ".", ">"})
	if err != nil {
		log.Fatalf("%s", err)
	}

	ySize = len(pre)
	xSize = len(pre[0])

	cur := make([][]string, ySize)
	for i := range cur {
		cur[i] = make([]string, xSize)
	}

	copyVal(cur, pre)

	step := 1
	for {
		move(cur)

		if compare(pre, cur) {
			break
		}

		copyVal(pre, cur)

		step++
	}

	log.Print(step)
}
