package main

import (
	"advent-code-2021/util"
	"log"
)

func main() {
	const (
		path     = "inputs/"
		fileName = "input12-ex"
	)

	_, err := util.ReadNumsInFile(path + fileName)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
