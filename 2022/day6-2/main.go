package main

var splitNum int

func include(str string) bool {
	for i := 0; i < splitNum; i++ {
		for j := 1; j < splitNum; j++ {
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
	for i := 0; i < length-splitNum; i++ {
		end := i + splitNum
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
