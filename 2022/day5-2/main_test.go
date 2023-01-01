package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"path/filepath"
	"testing"
	"yoon1/adventcode/util"
)

func Test_extractOperand(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		operand := extractOperand("move 1 from 2 to 1")
		assert.Equal(t, Operand{1, 2, 1}, operand)
	})
}

//    [D]
//[N] [C]
//[Z] [M] [P]
// 1   2   3
func Test_main(t *testing.T) {
	t.Run("EXAMPLE PASS", func(t *testing.T) {
		path, _ := filepath.Abs("ex1")
		lines, err := util.ReadLinesInFile(path)
		if err != nil {
			log.Fatal(err)
		}
		boxes = []*Box{
			{"Z", "N"},
			{"M", "C", "D"},
			{"P"},
		}
		res := solve(lines)
		assert.Equal(t, res, "MCD")
	})

	t.Run("SOLVE", func(t *testing.T) {
		path, _ := filepath.Abs("ex2")
		lines, err := util.ReadLinesInFile(path)
		if err != nil {
			log.Fatal(err)
		}

		//         [J]         [B]     [T]
		//        [M] [L]     [Q] [L] [R]
		//        [G] [Q]     [W] [S] [B] [L]
		//[D]     [D] [T]     [M] [G] [V] [P]
		//[T]     [N] [N] [N] [D] [J] [G] [N]
		//[W] [H] [H] [S] [C] [N] [R] [W] [D]
		//[N] [P] [P] [W] [H] [H] [B] [N] [G]
		//[L] [C] [W] [C] [P] [T] [M] [Z] [W]
		boxes = []*Box{
			{"L", "N", "W", "T", "D"},
			{"C", "P", "H"},
			{"W", "P", "H", "N", "D", "G", "M", "J"},
			{"C", "W", "S", "N", "T", "Q", "L"},
			{"P", "H", "C", "N"},
			{"T", "H", "N", "D", "M", "W", "Q", "B"},
			{"M", "B", "R", "J", "G", "S", "L"},
			{"Z", "N", "W", "G", "V", "B", "R", "T"},
			{"W", "G", "D", "N", "P", "L"},
		}
		res := solve(lines)
		log.Println(res)
	})
}
