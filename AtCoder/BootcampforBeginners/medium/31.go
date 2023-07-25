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

	excludedArr := []int{}
	// 偶数が複数ある場合は、4にできるので、削除する
	for _, v := range arr {
		if v%4 == 0 || v%2 != 0 {
			excludedArr = append(excludedArr, v)
		}
	}

	if len(excludedArr) != n {
		// 偶数を削除した場合、一つは残す必要がある
		excludedArr = append(excludedArr, 2)
	}

	need := len(excludedArr) / 2
	counter := 0
	for _, v := range excludedArr {
		if v%4 == 0 {
			counter++
		}
	}

	if need <= counter {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
