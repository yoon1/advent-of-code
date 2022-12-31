package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"path/filepath"
	"testing"
	"yoon1/adventcode/util"
)

func Test_main(t *testing.T) {
	t.Run("EXAMPLE PASS", func(t *testing.T) {
		path, _ := filepath.Abs("ex1")
		lines, err := util.ReadLinesInFile(path)
		if err != nil {
			log.Fatal(err)
		}

		res := solve(lines)
		assert.Equal(t, res, 24000)
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
