package main

import (
	"advent-code-2021/util"
	"fmt"
	"log"
	"strings"
)

func calc(s string) int {
	m := map[string]int{}
	for _, c := range s {
		m[string(c)]++
	}

	max := 0
	min := int(^uint(0) >> 1)
	for _, val := range m {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	return max - min
}

func step(q string, m *map[string]string) string {
	vals := ""

	length := len(q)
	for i := 1; i < length; i++ {
		key := string(q[i-1]) + string(q[i])
		vals += (*m)[key]
	}

	res := string(q[0])
	for i := 1; i < length; i++ {
		res += string(vals[i-1]) + string(q[i])
	}

	return res
}

func parse(lines *[]string) (string, *map[string]string) {
	q := ""
	m := map[string]string{}
	for _, line := range *lines {
		if line == "" {
			continue
		} else if strings.Contains(line, " -> ") {
			arrow := strings.Split(line, " -> ")
			m[arrow[0]] = arrow[1]
		} else {
			q = line
		}
	}

	return q, &m
}

func main() {
	lines, err := util.ReadLinesInFile("input14-2")
	if err != nil {
		fmt.Errorf("[READ FILE ERROR] %s", err)
	}

	q, m := parse(&lines)
	for i := 1; i <= 40; i++ {
		q = step(q, m)
	}

	log.Println(calc(q))
}
