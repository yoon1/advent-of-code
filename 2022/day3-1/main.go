package main

func point(c int32) int {
	var p int
	if c >= 'a' && c <= 'z' {
		p = int(c-'a') + 1
	} else if c >= 'A' && c <= 'Z' {
		p = int(c-'A') + 26 + 1
	}

	return p
}

func common(a, b string) int {
	var alphas [53]int
	for _, c := range a {
		alphas[point(c)]++
	}

	for _, c := range b {
		p := point(c)
		if alphas[p] > 0 {
			return p
		}
	}

	return 0
}

func pry(bag string) int {
	mid := len(bag) / 2
	return common(bag[:mid], bag[mid:])
}

func solve(strs []string) int {
	var res int
	for _, str := range strs {
		res += pry(str)
	}
	return res
}
