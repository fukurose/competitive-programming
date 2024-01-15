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

func main() {
	n := scanInt()
	m := scanInt()

	arr := make([][]int, 200000+1)

	for i := 0; i < m; i++ {
		a := scanInt()
		b := scanInt()
		arr[a] = append(arr[a], b)
	}

	ans := false
	for _, s1 := range arr[1] {
		for _, s2 := range arr[s1] {
			if s2 == n {
				ans = true
				break
			}
		}
	}

	if ans {
		fmt.Println("POSSIBLE")
	} else {
		fmt.Println("IMPOSSIBLE")
	}
}
