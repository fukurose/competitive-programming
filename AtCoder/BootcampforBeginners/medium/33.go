package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

func init() {
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
}

type Stack struct {
	container [10]string
	top       int
}

func (s *Stack) isEmpty() bool {
	return s.top == 0
}

func (s *Stack) push(x string) {
	s.container[s.top] = x
	s.top++
}

func (s *Stack) pop() string {
	x := s.container[s.top]
	s.top--
	return x
}

func (s *Stack) result() string {
	r := ""
	for i := 0; i < s.top; i++ {
		r += s.container[i]
	}
	return r
}

func main() {
	inputs := strings.Split(scanString(), "")
	s := Stack{}
	for _, v := range inputs {
		if v == "B" {
			if !s.isEmpty() {
				s.pop()
			}
		} else {
			s.push(v)
		}
	}

	fmt.Println(s.result())
}
