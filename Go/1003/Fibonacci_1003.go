package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	F0 := make([]int, 41) // F0[4] = Fibo(4)가 호출하는 Fibo(0) 횟수
	F1 := make([]int, 41) // F1[4] = Fibo(4)가 호출하는 Fibo(1) 횟수

	F0[1], F1[0] = 0, 0
	F0[0], F0[2], F1[1], F1[2] = 1, 1, 1, 1
	for i := 3; i < 41; i++ {
		F0[i] = F0[i-1] + F0[i-2]
		F1[i] = F1[i-1] + F1[i-2]
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < T; i++ {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())
		fmt.Println(F0[N], F1[N])
	}
}
