package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	D := make([]int, 1000001)
	D[0], D[1] = 0, 0
	D[2], D[3] = 1, 1
	for i := 4; i <= 1000000; i++ {
		D[i] = D[i-1] + 1
		if i%2 == 0 && D[i/2]+1 < D[i] {
			D[i] = D[i/2] + 1
		}
		if i%3 == 0 && D[i/3]+1 < D[i] {
			D[i] = D[i/3] + 1
		}
	}

	N, _ := strconv.Atoi(scanner.Text())
	fmt.Println(D[N])
}
