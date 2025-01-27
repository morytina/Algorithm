package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 소수 판별 배열 초기화
	isPrime := make([]bool, 1001)
	for i := 2; i < 1001; i++ {
		isPrime[i] = true
	}

	// 에라토스테네스의 체
	for i := 2; i < 1001; i++ {
		if isPrime[i] {
			for cur := i * 2; cur < 1001; cur += i {
				isPrime[cur] = false
			}
		}
	}

	// 입력 처리
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text()) // 첫 번째 줄 입력: N

	scanner.Scan()
	numbers := strings.Split(scanner.Text(), " ") // 두 번째 줄 입력: 공백으로 구분된 숫자들

	// 소수 개수 카운트
	ans := 0
	for i := 0; i < N; i++ {
		num, _ := strconv.Atoi(numbers[i])
		if isPrime[num] {
			ans++
		}
	}

	// 결과 출력
	fmt.Println(ans)
}
