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
	if containsAB(a, b) || containsAB(b, a) {
		return true
	}
	return false
}

type Line struct {
	Left  int
	Right int
}

func containsAB(a, b Line) bool {
	if a.Left <= b.Left && a.Right >= b.Right {
		return true
	}
	return false
}
