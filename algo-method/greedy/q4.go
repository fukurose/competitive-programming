package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	m := scanInt()

	p := make([]int, n)
	for i := range p {
		p[i] = scanInt()
	}

	b := make([]int, m)
	for i := range b {
		b[i] = scanInt()
	}

	sort.Ints(p)

	count := 0
	b_index := 0
	for i := range p {
		for j := b_index; j < m; j++ {
			if b[j] < p[i] {
				continue
			}

			count++
			b_index = j + 1
			break
		}
	}

	fmt.Println(count)
}
