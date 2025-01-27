package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 입력 처리
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text()) // 첫 번째 줄: N 값 입력

	// 슬라이스 초기화 및 점수 입력
	score := make([]float64, N)
	max := 0.0
	scanner.Scan()
	inputScores := strings.Split(scanner.Text(), " ")

	// 점수 입력 및 최대값 찾기
	for i := 0; i < N; i++ {
		score[i], _ = strconv.ParseFloat(inputScores[i], 64)
		if score[i] > max {
			max = score[i]
		}
	}

	// 점수 변환 및 합계 계산
	sum := 0.0
	for i := 0; i < N; i++ {
		score[i] = score[i] / max * 100
		sum += score[i]
	}

	// 평균 출력
	fmt.Printf("%.6f\n", sum/float64(N))
}
