package main

import (
	"advent-code-2021/util"
	"fmt"
	"log"
	"strconv"
)

// 1번 라인 - 이미지 향상 알고리즘:
// 2번: 입력 이미지, 밝은 픽셀('#'), 어두운 픽셀('.')

// 입력 영상의 모든 픽셀을 출력 영상으로 동시에 변환해 영상을 향상시키는 방법을 설명한다.
// 출력 영상의 각 픽셀은 해당 입력 영상 픽셀을 중심으로 한, 3x3 정사각형의 픽셀을 보고 결정됩니다.
// 따라서 출력 이미지에서 (5,10)에서 픽셀의 값을 결정하려면 입력 이미지에서 (4,9), (4,10), (4,11), (5,10), (5,11), (5,11), (6,9), (6,10) 및 (6,11)의 9개의 픽셀을 고려해야 합니다.
// -- 1개의 픽셀 값을 결정하려면 테두리를 보라고
// -- NOTE: ** 이 9개의 입력 픽셀은 이미지 향상 알고리즘 문자열에서 인덱스로 사용되는 `단일 이진수로 결합`됩니다.

// TODO: exapmle
// # . . # .
// #[. . .].
// #[# . .]#  <-- target: 정 가운데
// .[. # .].
// . . # # #
// 왼쪽 위에서 시작하여 각 행에 걸쳐 읽으면 이 픽셀은 ..., 그 다음 #..., 그리고 이 폼들을 조합한 .#...#이 됩니다.
// 어두운 픽셀(.)을 0으로, 밝은 픽셀(#)을 1로 바꿈으로써 이진수 000100010(34)을 만들 수 있다.
// 9bits -> 2^9 = 512가지

// 이미지 향상 알고리즘 문자열은 가능한 모든 `9비트 이진수`와 일치하기에 충분한 길이로 정확히 `512`자입니다.
// 문자열의 처음 몇 문자(0부터 번호가 매겨짐)는 다음과 같습니다.
// 0         10        20        30  [34]    40        50        60        70
// |         |         |         |   |     |         |         |         |
// ..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..##
// NOTE: 34번째가 '#' 임을 확인.
// NOTE: 따라서 출력 이미지 중앙의 출력 픽셀은 `#`, 즉 밝은 픽셀이어야 합니다. 결정!!
// 그런 다음 이 프로세스를 반복하여 출력 이미지의 모든 픽셀을 계산할 수 있습니다. -> 이거는 전부 순환한 다음에 하는 행위인가?
// TODO: 사용자가 가진 작은 입력 이미지는 실제 무한 입력 이미지의 작은 영역일 뿐이며, 나머지 입력 이미지는 어두운 픽셀(.)로 구성됩니다.
// 예를 들어 공간을 절약하기 위해 무한 크기의 입력 및 출력 영상의 일부만 표시됩니다.
// -- -> 아마 테두리를 만들어서 쓰라는 거겠지?? ^-^ ;
// 이미지 향상 알고리즘을 모든 픽셀에 `동시에` 적용하면 다음과 같은 출력 이미지를 얻을 수 있습니다. => 즉, temp 배열에 담고 있어야 할 것 같다!!
func main() {
	const (
		//fileName1 = "input20-1"
		//fileName2 = "input20-2"
		fileName1 = "input20-3"
		fileName2 = "input20-4"
	)

	lines, err := util.ReadLinesInFile(fileName1)
	if err != nil {
		log.Fatalf("%s", err)
	}

	algorithm := lines[0]

	image, err := util.ReadCharsInFile(fileName2)
	if err != nil {
		log.Fatalf("%s", err)
	}

	ySize := len(image)
	xSize := len(image[0])

	const simulationNum = 2 // 반복 횟수
	for i := 0; i <= simulationNum*2; i++ {
		image = makeFrame(image, ySize, xSize)
		ySize += 2
		xSize += 2
	}

	for i := 0; i < simulationNum; i++ {
		image = solve(image, ySize, xSize, algorithm, simulationNum)
	}

	log.Printf("%d", calc(image, ySize, xSize, simulationNum))
}

func makeFrame(image [][]string, ySize, xSize int) [][]string {
	newImage := [][]string{}
	line := []string{}
	for j := 0; j < xSize+2; j++ {
		line = append(line, ".")
	}
	newImage = append(newImage, line)

	for i := 0; i < ySize; i++ {
		line = []string{}
		line = append(line, ".")
		for j := 0; j < xSize; j++ {
			line = append(line, image[i][j])
		}
		line = append(line, ".")
		newImage = append(newImage, line)
	}

	line = []string{}
	for j := 0; j < xSize+2; j++ {
		line = append(line, ".")
	}
	newImage = append(newImage, line)

	return newImage
}

func around(image [][]string, y, x, ySize, xSize int) string {
	dy := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}
	dx := []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}

	str := ""
	const length = 9
	for i := 0; i < length; i++ {
		ny := dy[i] + y
		nx := dx[i] + x
		if !(ny >= 0 && ny < ySize && nx >= 0 && nx < xSize) {
			str += "."
		} else {
			str += image[ny][nx]
		}
	}

	return str
}

func dex(bin string) int {
	num := 0
	mul := 1
	length := len(bin)
	for i := length - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(bin[i]))
		num += n * mul
		mul *= 2
	}

	return num
}

func decode(code string) int {
	length := len(code)
	var bin = ""
	for i := 0; i < length; i++ {
		if string(code[i]) == "." {
			bin += "0"
		} else {
			bin += "1"
		}
	}

	return dex(bin)
}

func solve(image [][]string, ySize, xSize int, algorithm string, n int) [][]string {
	newImage := [][]string{}
	for i := 0; i < ySize; i++ {
		line := []string{}
		for j := 0; j < xSize; j++ {
			line = append(line, ".")
		}
		newImage = append(newImage, line)
	}

	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			code := around(image, y, x, ySize, xSize)
			newImage[y][x] = enhance(decode(code), algorithm)
		}
	}

	return newImage
}

func enhance(idx int, algorithm string) string {
	return string(algorithm[idx])
}

func calc(image [][]string, ySize, xSize int, n int) int {
	count := 0
	for y := n; y < ySize-n; y++ {
		for x := n; x < xSize-n; x++ {
			if image[y][x] == "#" {
				count++
			}
		}
	}

	return count
}

func printImage(image [][]string) {
	for _, img := range image {
		fmt.Println(img)
	}
	log.Printf("\n\n")
}
