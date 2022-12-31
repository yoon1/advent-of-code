package main

type Point int

const (
	Win      Point = 6
	Draw     Point = 3
	Lose     Point = 0
	Rock     Point = 1
	Paper    Point = 2
	Scissors Point = 3
)

var winPointMap = map[string]Point{
	"A X": Draw + Rock,
	"A Y": Win + Paper,
	"A Z": Lose + Scissors,
	"B X": Lose + Rock,
	"B Y": Draw + Paper,
	"B Z": Win + Scissors,
	"C X": Win + Rock,
	"C Y": Lose + Paper,
	"C Z": Draw + Scissors,
}

// 1 for Rock, 2 for Paper, and 3 for Scissors
// (A, B, C) (X, Y, Z)
// X: lose, Y: draw, Z: win
func game(round string) Point {
	return winPointMap[round]
}

func round(str string) Point {
	return game(str)
}

func solve(strs []string) int {
	var point int
	for _, str := range strs {
		point += int(round(str))
	}
	return point
}
