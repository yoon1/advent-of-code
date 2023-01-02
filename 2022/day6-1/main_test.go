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
		{"bvwb", true},
		{"bvwa", false},
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
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tc := range tcs {
		t.Run("SUCCESS", func(t *testing.T) {
			res := findFirst(tc.Str)
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
		res := solve(lines[0])
		assert.Equal(t, 5, res)
	})

	t.Run("SOLVE", func(t *testing.T) {
		path, _ := filepath.Abs("ex2")
		lines, err := util.ReadLinesInFile(path)
		if err != nil {
			log.Fatal(err)
		}
		res := solve(lines[0])
		log.Println(res)
	})
}
