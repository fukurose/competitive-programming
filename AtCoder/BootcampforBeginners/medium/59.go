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

	ans := 0

	for i := 0; i < n-1; i++ {
		if i+1 == arr[i] {
			t := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = t
			ans++
		}
	}

	if arr[n-1] == n {
		ans++
	}

	fmt.Println(ans)
}
