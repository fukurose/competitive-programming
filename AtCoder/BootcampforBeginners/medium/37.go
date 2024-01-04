package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

type Queue [][]int

func (q *Queue) enQueue(x []int) {
	*q = append(*q, x)
}

func (q *Queue) deQueue() []int {
	x := (*q)[0]
	*q = (*q)[1:]

	return x
}

func min(a []int) int {
	m := 100

	for i := range a {
		if m > a[i] {
			m = a[i]
		}
	}
	return m
}

func splitZero(a []int) [][]int {
	s := 0
	e := len(a)
	x := [][]int{}
	count := 0
	for i := range a {
		if a[i] == 0 {
			if len(a[s:i]) > 0 {
				x = append(x, a[s:i])
				count++
			}
			s = i + 1
		}
	}

	if len(a[s:e]) > 0 {
		x = append(x, a[s:e])
		count++
	}

	return x
}

func main() {
	n := scanInt()

	arr := make([]int, n)

	for i := range arr {
		arr[i] = scanInt()
	}

	q := Queue{}
	q.enQueue(arr)

	ans := 0
	for {
		if len(q) == 0 {
			break
		}
		a := q.deQueue()
		m := min(a)
		ans += m

		for i := range a {
			a[i] -= m
		}

		for _, x := range splitZero(a) {
			q.enQueue(x)
		}
	}

	fmt.Println(ans)
}
