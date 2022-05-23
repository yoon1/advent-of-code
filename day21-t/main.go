package main

func quantum_dice(p1_pos, p2_pos, p1_score, p2_score int) (p1_wins, p2_wins int) {
	cases := [][]int{{1, 1, 1}, {1, 1, 2}, {1, 1, 3}, {1, 2, 1}, {1, 2, 2}, {1, 2, 3}, {1, 3, 1}, {1, 3, 2}, {1, 3, 3}, {2, 1, 1}, {2, 1, 2}, {2, 1, 3}, {2, 2, 1}, {2, 2, 2}, {2, 2, 3}, {2, 3, 1}, {2, 3, 2}, {2, 3, 3}, {3, 1, 1}, {3, 1, 2}, {3, 1, 3}, {3, 2, 1}, {3, 2, 2}, {3, 2, 3}, {3, 3, 1}, {3, 3, 2}, {3, 3, 3}}
	p1_wins, p2_wins = 0, 0
	for _, c := range cases {
		total := c[0] + c[1] + c[2]
		p1_pos = (p1_pos+total-1)%10 + 1 // 20 > 10
		p1_score = p1_score + p1_pos
		p2_pos = (p2_pos+total-1)%10 + 1
		p2_score = p2_score + p2_pos
		if p1_score >= 21 || p2_score >= 21 {
			if p1_score >= 21 {
				p1_wins += 1
			} else {
				p2_wins += 1
			}
			// 27!
		} else {
			p1, p2 := quantum_dice(p1_pos, p2_pos, p1_score, p2_score)
			p1_wins += p1
			p2_wins += p2
		}

	}
	return p1_wins, p2_wins
}

func main() {
	p1, p2 := quantum_dice(4, 8, 0, 0)
	print(p1, ",", p2)
}
