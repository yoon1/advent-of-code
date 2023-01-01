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

func compose(a string) (alphas [53]int) {
	for _, c := range a {
		alphas[point(c)]++
	}
	return
}

func common(a, b, c string) int {
	alphasA := compose(a)
	alphasB := compose(b)
	for _, ch := range c {
		p := point(ch)
		if alphasA[p] > 0 && alphasB[p] > 0 {
			return p
		}
	}

	return 0
}

func pry(bag [3]string) int {
	return common(bag[0], bag[1], bag[2])
}

func solve(strs []string) int {
	var res int
	for i := 0; i < len(strs); i = i + 3 {
		res += pry([3]string{strs[i], strs[i+1], strs[i+2]})
	}
	return res
}
