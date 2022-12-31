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
	"A X": Lose + Scissors,
	"A Y": Draw + Rock,
	"A Z": Win + Paper,
	"B X": Lose + Rock,
	"B Y": Draw + Paper,
	"B Z": Win + Scissors,
	"C X": Lose + Paper,
	"C Y": Draw + Scissors,
	"C Z": Win + Rock,
}

// 1 for Rock, 2 for Paper, and 3 for Scissors
// (A, B, C) (X, Y, Z)
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
