package main

func include(str string) bool {
	for i := 0; i < 4; i++ {
		for j := 1; j < 4; j++ {
			if i == j {
				continue
			}
			if str[i] == str[j] {
				return true
			}
		}
	}
	return false
}

func findFirst(str string) int {
	length := len(str)
	for i := 0; i < length-4; i++ {
		end := i + 4
		if length < end {
			end = length
		}
		t := str[i:end]
		if length < end+1 {
			continue
		}
		if !include(t) {
			return end
		}
	}

	return 0
}

func solve(str string) int {
	return findFirst(str)
}
