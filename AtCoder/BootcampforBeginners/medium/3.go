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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	s := strings.Split(scanString(), "")

	arr := make([]int, len(s)+1)
	size := len(arr)

	arr[0] = 0
	for i := 1; i < size; i++ {
		if s[i-1] == "<" {
			arr[i] = arr[i-1] + 1
		} else {
			arr[i] = 0
		}
	}

	arr[size-1] = max(0, arr[size-1])

	for i := size - 2; i >= 0; i-- {
		if s[i] == ">" {
			arr[i] = max(arr[i], arr[i+1]+1)
		} else {
			arr[i] = max(arr[i], 0)
		}
	}

	ans := 0
	for i := range arr {
		ans += arr[i]
	}
	fmt.Println(ans)
}
