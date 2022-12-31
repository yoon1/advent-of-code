package main

import (
	"math"
	"yoon1/adventcode/util"
)

type Calories int

func solve(strs []string) int {
	var max, cur Calories
	for _, str := range strs {
		if len(str) == 0 {
			max = Calories(math.Max(float64(max), float64(cur)))
			cur = 0
			continue
		}
		cur += Calories(util.StringToInt(str))
	}

	max = Calories(math.Max(float64(max), float64(cur)))

	return int(max)
}
