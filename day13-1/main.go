package main

import (
	"advent-code-2021/util"
	"fmt"
	"log"
	"strings"
)

var (
	RowSize = 5000
	ColSize = 5000
)

func getArray() *[][]string {
	array := [][]string{}
	for i := 0; i < RowSize; i++ {
		row := []string{}
		for j := 0; j < ColSize; j++ {
			row = append(row, ".")
		}
		array = append(array, row)
	}

	return &array
}

func check(array *[][]string, row, col int) bool {
	return (*array)[row][col] == "#"
}

func mark(array *[][]string, row, col int) {
	(*array)[row][col] = "#"
}

func count(array *[][]string, row, col int) (cnt int) {
	for row := 0; row < RowSize; row++ {
		for col := 0; col < ColSize; col++ {
			if check(array, row, col) {
				cnt++
			}
		}
	}
	return
}

func fold(array *[][]string, variable string, value int) {
	switch variable {
	case "y":
		for row := value; row >= 0; row-- {
			for col := 0; col < ColSize; col++ {
				if check(array, value+value-row, col) {
					mark(array, row, col)
				}
			}
		}
		RowSize = value
	case "x":
		for row := 0; row < RowSize; row++ {
			for col := value; col >= 0; col-- {
				if check(array, row, value+value-col) {
					mark(array, row, col)
				}
			}
		}
		ColSize = value
	}
}

func solve(array *[][]string, lines *[]string) {
	maxRow := 0
	maxCol := 0

	const search = "fold along "
	for _, line := range *lines {
		if strings.Contains(line, ",") {
			// read points
			points := strings.Split(line, ",")
			col := util.StringToInt(points[0])
			row := util.StringToInt(points[1])
			mark(array, row, col)
			maxCol = util.GreaterInt(col, maxCol)
			maxRow = util.GreaterInt(row, maxRow)
		} else if strings.Contains(line, "=") {
			// read folding lines
			idx := strings.LastIndex(line, search) + len(search)
			t := fmt.Sprintf("%s", line[idx:])
			points := strings.Split(t, "=")
			variable := points[0]
			value := util.StringToInt(points[1])
			fold(array, variable, value)
			log.Println("[FOLD RESULT] ", count(array, RowSize, ColSize))
		} else {
			RowSize = maxRow + 1
			ColSize = maxCol + 1
		}
	}
}

func main() {
	lines, err := util.ReadLinesInFile("input13-2")
	if err != nil {
		fmt.Errorf("[READ FILE ERROR] %s", err)
	}

	array := getArray()
	solve(array, &lines)
	util.PrintStringMatrix(*array, RowSize, ColSize)
}
