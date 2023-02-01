package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

var operators = []string{"+", "-", "*"}

type Stack struct {
	container [10000]int
	top       int
}

func (s *Stack) isFull() bool {
	return s.top == (len(s.container) - 1)
}

func (s *Stack) isEmpty() bool {
	return s.top == 0
}

func (s *Stack) push(x int) {
	s.top++
	s.container[s.top] = x
}

func (s *Stack) pop() int {
	x := s.container[s.top]
	s.top--

	return x
}

func isOperator(x string) bool {
	for _, ope := range operators {
		if ope == x {
			return true
		}
	}
	return false
}

func cal(ope string, x, y int) int {
	var ans int
	switch ope {
	case "+":
		ans = x + y
	case "-":
		ans = x - y
	case "*":
		ans = x * y
	}

	return ans
}

func F(arr []string) int {
	s := Stack{}
	for _, e := range arr {
		if isOperator(e) {
			a := s.pop()
			b := s.pop()
			x := cal(e, b, a)
			s.push(x)
		} else {
			n, _ := strconv.Atoi(e)
			s.push(n)
		}
	}
	return s.pop()
}

func main() {
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " ")
	ans := F(arr)
	fmt.Println(ans)
}
