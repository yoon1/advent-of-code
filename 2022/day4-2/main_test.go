package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"path/filepath"
	"testing"
	"yoon1/adventcode/util"
)

func Test_include(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		res := include("2-8,3-7")
		assert.Equal(t, true, res)
	})
}

func Test_main(t *testing.T) {
	t.Run("EXAMPLE PASS", func(t *testing.T) {
		path, _ := filepath.Abs("ex1")
		lines, err := util.ReadLinesInFile(path)
		if err != nil {
			log.Fatal(err)
		}

		res := solve(lines)
		assert.Equal(t, res, 4)
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
