package main

import (
	"advent-code-2021/util"
	"fmt"
	"log"
	"strings"
)

const (
	LEFT  = 0
	RIGHT = 1
)

func alphaCount(countMap *map[string]uint64) map[string]uint64 {
	alphaCountMap := map[string]uint64{}
	for key, count := range *countMap {
		alphaCountMap[string(key[LEFT])] += count
	}

	return alphaCountMap
}

func calc(stringCountMap *map[string]uint64) uint64 {
	var (
		max = uint64(0)
		min = uint64(^uint(0) >> 1)
	)

	m := alphaCount(stringCountMap)
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

func first(s string) *map[string]uint64 {
	counts := &map[string]uint64{}
	length := len(s)

	for i := 1; i < length; i++ {
		p := string(s[i-1]) + string(s[i])
		(*counts)[p]++
	}

	return counts
}

func step(num int, q string, m *map[string]string) uint64 {
	countMap := first(q)

	for i := 0; i < num; i++ {
		cur := map[string]uint64{}
		for key, cnt := range *countMap {
			cur[string(key[LEFT])+(*m)[key]] += cnt
			cur[(*m)[key]+string(key[RIGHT])] += cnt
		}
		countMap = &cur
	}

	return calc(countMap)
}

func parse(lines *[]string) (string, *map[string]string) {
	q := ""
	m := map[string]string{}
	for _, line := range *lines {
		if line == "" {
			continue
		} else if strings.Contains(line, " -> ") {
			arrow := strings.Split(line, " -> ")
			m[arrow[LEFT]] = arrow[RIGHT]
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
	log.Println(step(10, q, m))
	log.Println(step(40, q, m))
}
