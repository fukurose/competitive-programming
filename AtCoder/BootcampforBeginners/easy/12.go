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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	k := scanInt()
	n := scanInt()

	arr := make([]int, n)
	for i := range arr {
		arr[i] = scanInt()
	}

	max_d := 0
	for i := 0; i < n-1; i++ {
		max_d = max(max_d, arr[i+1]-arr[i])
	}

	max_d = max(max_d, (k-arr[n-1])+arr[0])

	fmt.Println(k - max_d)
}
