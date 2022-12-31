package main

import (
	"sort"
	"yoon1/adventcode/util"
)

type Calories int

type Food struct {
	calories Calories
}

func solve(strs []string) int {
	foods := []Food{}
	food := Food{}

	for _, str := range strs {
		if len(str) == 0 {
			foods = append(foods, food)
			food = Food{}
			continue
		}
		food.calories += Calories(util.StringToInt(str))
	}

	foods = append(foods, food)

	// sort
	sort.Slice(foods, func(i, j int) bool {
		return foods[i].calories > foods[j].calories
	})

	var sum int
	var topLen = 3
	if len(foods) < 3 {
		topLen = len(foods)
	}
	for i := 0; i < topLen; i++ {
		sum += int(foods[i].calories)
	}

	return sum
}
