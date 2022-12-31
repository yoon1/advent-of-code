package main

import (
	"advent-code-2021/util"
	"log"
)

type Point struct {
	y int
	x int
}

func shoot(y, x, xStart, xEnd, yStart, yEnd int) (int, bool) {
	maxY := 0

	mv := Point{y, x}
	pos := Point{0, 0}

	for {
		pos.y += mv.y
		pos.x += mv.x

		mv.x = util.MaxInt(mv.x-1, 0)
		mv.y--

		maxY = util.MaxInt(pos.y, maxY)

		if pos.x >= xStart && pos.x <= xEnd && pos.y >= yStart && pos.y <= yEnd {
			return maxY, true
		}

		if pos.y < yStart && mv.y < 0 {
			break
		}

		if pos.x > xEnd {
			break
		}
	}

	return 0, false
}

func solve(xStart, xEnd, yStart, yEnd int) (result int, count int) {
	for y := yStart; y <= -yStart; y++ {
		for x := 0; x <= xEnd; x++ {
			r, ok := shoot(y, x, xStart, xEnd, yStart, yEnd)
			result = util.MaxInt(result, r)
			if ok {
				count++
			}
		}
	}
	return
}

func main() {
	// solve 1, solve 2
	log.Println(solve(20, 30, -10, -5))
	log.Println(solve(25, 67, -260, -200))
}
