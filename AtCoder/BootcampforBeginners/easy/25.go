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

func init() {
	scanner.Split(bufio.ScanWords)
}

func main() {
	n := scanInt()
	p_arr := make([]int, n)
	for i := range p_arr {

		p_arr[i] = scanInt()
	}
	q_arr := make([]int, n)
	for i := range p_arr {
		q_arr[i] = scanInt()
	}

	factorial := make([]int, n+1)
	factorial[0] = 1
	factorial[1] = 1
	for i := 2; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	used_p := make([]bool, n+1)

	p_index := 1
	for i, p := range p_arr {
		count := 0
		for j := 1; j <= p; j++ {
			if used_p[j] {
				continue
			}
			count++
		}
		p_index += count * factorial[n-1-i]
		used_p[p] = true
	}

	used_q := make([]bool, n+1)

	q_index := 1
	for i, q := range q_arr {
		count := 0
		for j := 1; j <= q; j++ {
			if used_q[j] {
				continue
			}
			count++
		}
		q_index += count * factorial[n-1-i]
		used_q[q] = true

	}
	if p_index > q_index {
		fmt.Println(p_index - q_index)
	} else {
		fmt.Println(q_index - p_index)
	}
}
