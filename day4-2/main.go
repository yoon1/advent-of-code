package main

import (
	"strings"
	"yoon1/adventcode/util"
)

func solve(strs []string) int {
	var res int
	for _, str := range strs {
		if include(str) {
			res++

		}
	}
	return res
}

func extractLine(str string) Line {
	splited := strings.Split(str, "-")
	return Line{util.StringToInt(splited[0]), util.StringToInt(splited[1])}
}

func extract(str string) (Line, Line) {
	splited := strings.Split(str, ",")
	return extractLine(splited[0]), extractLine(splited[1])
}

func include(str string) bool {
	a, b := extract(str)
	if overlapAB(a, b) || overlapAB(b, a) {
		//log.Println(str)
		return true
	}
	return false
}

type Line struct {
	Left  int
	Right int
}

func overlapAB(a, b Line) bool {
	if a.Left <= b.Left && a.Right >= b.Left {
		return true
	}
	if a.Left <= b.Right && a.Right >= b.Right {
		return true
	}

	return false
}
