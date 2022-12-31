package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"path/filepath"
	"testing"
	"yoon1/adventcode/util"
)

func Test_round(t *testing.T) {
	tcs := []struct {
		Input string
		Ans   Point
	}{
		{"A Y", 4},
		{"B X", 1},
		{"C Z", 7},
	}
	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			res := round(tc.Input)
			assert.Equal(t, res, tc.Ans)
		})
	}
}
func Test_main(t *testing.T) {
	t.Run("EXAMPLE PASS", func(t *testing.T) {
		path, _ := filepath.Abs("ex1")
		lines, err := util.ReadLinesInFile(path)
		if err != nil {
			log.Fatal(err)
		}

		res := solve(lines)
		assert.Equal(t, res, 12)
	})

	t.Run("SOLVE", func(t *testing.T) {
		path, _ := filepath.Abs("ex2")
		lines, err := util.ReadLinesInFile(path)
		if err != nil {
			log.Fatal(err)
		}

		res := solve(lines)
		log.Println(res)
	})
}
