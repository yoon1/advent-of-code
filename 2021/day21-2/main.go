package main

import "log"

// * * * * VARIABLE
// 게임보드: 1 - 10이 표시된 원형 트랙
// 주사위 1개, 말 2개
// 각 플레이어의 starting point는 무작위(퍼즐에서 입력받음.)

// * * * * GAME
// 1. Player 1이 먼저 이동한다. (순서대로)
// 2. 각 플레이어의 차례에 플레이어는 주사위를 `3번` 굴리고 결과를 더한다.
// 3. 플레이어는 트랙을 돌면서 폰을 여러 번 앞으로 이동시킨다(즉, `값이 증가하는 순서대로 공간에서` 시계 방향으로 이동, 10시 이후 1로 감는다).
// 4. 따라서 플레이어가 7번 공간에 있을 때 2, 2, 1번을 굴리면 8, 9, 10, 1번 공간으로 5번 전진하다가 2번에서 멈춥니다. -> TODO: modula

// 5. 각 플레이어가 이동한 후, 폰이 정지한 공간의 값만큼 점수를 올립니다.
// 6. 선수들의 점수는 0부터 시작한다.
// 7. 따라서 첫 번째 플레이어가 7번 공간에서 시작하여 총 5번을 굴리면 2번 공간에서 멈추고 2번을 더한다(총점 2).
// 8. 이 게임은 점수가 1000점 이상인 플레이어가 승리하는 것으로 즉시 종료됩니다.
// 9. 첫 번째 게임은 연습 게임이기 때문에 잠수함은 결정론적 주사위라는 라벨이 붙은 칸을 열고 100면 주사위가 빠진다. 이 다이는 항상 1을 먼저 굴리고, 2를 굴리고, 3을 굴리고, 100까지 굴리고, 그 후 다시 1에서 다시 시작합니다. 이 다이를 사용하여 재생합니다.

type Player struct {
	pawn  int
	point int
	count int
}

func main() {
	player1 := &Player{1, 0, 0}
	player2 := &Player{10, 0, 0}

	dice := 0
	for {
		if player1.turn(&dice, player2) {
			return
		}
		if player2.turn(&dice, player1) {
			return
		}
	}
}

func (player *Player) turn(dice *int, opponent *Player) bool {
	turn := 0
	for i := 0; i < 3; i++ {
		*dice = (*dice % 100) + 1
		turn += *dice
		player.count++
	}
	t := (player.pawn + turn) % 10
	if t == 0 {
		t = 10
	}
	if player.point+t >= 1000 {
		log.Println((player.count + opponent.count) * opponent.point)
		return true
	}
	player.point += t
	player.pawn = t

	return false
}
