package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"path/filepath"
	"testing"
	"yoon1/adventcode/util"
)

func Test_include(t *testing.T) {
	tcs := []struct {
		Str string
		Ans bool
	}{
		{"jpqmgbljsphdzt", true},
	}

	for _, tc := range tcs {
		t.Run("Success", func(t *testing.T) {
			res := include(tc.Str)
			assert.Equal(t, res, tc.Ans)
		})
	}
}
func Test_findFirst(t *testing.T) {
	tcs := []struct {
		Str string
		Ans int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		//{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		//{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		//{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		//{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, tc := range tcs {
		t.Run("SUCCESS", func(t *testing.T) {
			splitNum = 14
			res := findFirst(tc.Str)
			assert.Equal(t, res, tc.Ans)
		})
	}
}

func Test_main(t *testing.T) {
	t.Run("SOLVE", func(t *testing.T) {
		path, _ := filepath.Abs("ex2")
		lines, err := util.ReadLinesInFile(path)
		if err != nil {
			log.Fatal(err)
		}
		splitNum = 14
		res := solve(lines[0])
		log.Println(res)
	})
}
