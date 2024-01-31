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
	arr := make([]int, n)
	for i := range arr {
		arr[i] = scanInt()
	}
	ans := make([]int, n)

	isOK := true
	for i := 0; i < n; i++ {
		target := -1
		for j := 0; j < len(arr); j++ {
			if arr[j] == j+1 {
				target = j
			}
		}
		if target == -1 {
			isOK = false
			break
		}
		ans[n-i-1] = arr[target]
		arr = append(arr[:target], arr[target+1:]...)
	}

	if isOK {
		for i := range ans {
			fmt.Println(ans[i])
		}
	} else {
		fmt.Println(-1)
	}
}
