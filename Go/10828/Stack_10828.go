package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stack 자료구조 정의
type Stack struct {
	data []int
}

// Push: 스택에 값을 추가
func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

// Pop: 스택의 최상단 값을 제거하고 반환
func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return -1, false
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, true
}

// Top: 스택의 최상단 값 반환
func (s *Stack) Top() (int, bool) {
	if len(s.data) == 0 {
		return -1, false
	}
	return s.data[len(s.data)-1], true
}

// Size: 스택 크기 반환
func (s *Stack) Size() int {
	return len(s.data)
}

// IsEmpty: 스택이 비어있는지 확인
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// 메인 함수
func main() {
	// 입력 처리
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N)

	stack := &Stack{}

	// 명령 처리
	for i := 0; i < N; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		cmd := strings.Fields(line)

		switch cmd[0] {
		case "push":
			// push 명령어 처리
			value, _ := strconv.Atoi(cmd[1])
			stack.Push(value)
		case "top":
			// top 명령어 처리
			if value, ok := stack.Top(); ok {
				fmt.Fprintln(writer, value)
			} else {
				fmt.Fprintln(writer, -1)
			}
		case "pop":
			// pop 명령어 처리
			if value, ok := stack.Pop(); ok {
				fmt.Fprintln(writer, value)
			} else {
				fmt.Fprintln(writer, -1)
			}
		case "empty":
			// empty 명령어 처리
			if stack.IsEmpty() {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case "size":
			// size 명령어 처리
			fmt.Fprintln(writer, stack.Size())
		}
	}
}
