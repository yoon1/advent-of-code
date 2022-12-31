package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ErrOpenFile    = "Error file open!!"
	ErrReadFile    = "Error file read!!"
	ErrInvalidData = "Error invalid data!!"
)

func readFile(fileName string) ([]string, error) {
	//open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("%s: %s", ErrOpenFile, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// read line by line
	var lines []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	if err := fileScanner.Err(); err != nil {
		return nil, errors.New(ErrReadFile)
	}

	return lines, nil
}

func dive(commands []string) (int, error) {
	const (
		direction = 0
		distance  = 1
	)
	aim := 0
	depth := 0
	horize := 0
	for _, command := range commands {
		dirInfo := strings.Split(command, " ")
		var err error
		var dist int
		dir := dirInfo[direction]
		if dist, err = strconv.Atoi(dirInfo[distance]); err != nil {
			return 0, err
		}
		switch dir {
		case "forward":
			horize += dist
			depth += (dist * aim)
		case "up":
			aim -= dist
		case "down":
			aim += dist
		default:
			return 0, err
		}
	}

	log.Print("RES", depth, horize)
	return depth * horize, nil
}

func main() {
	const fileName = "input"

	mesurements, err := readFile(fileName)
	if err != nil {
		log.Fatalf("%s, %s", ErrInvalidData, err)
	}

	result, err := dive(mesurements)
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Print(result)
}
