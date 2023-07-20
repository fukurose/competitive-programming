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

func judge(a, b int) string {
	if a > b {
		return ">"
	} else if a < b {
		return "<"
	}
	return "="

}

func main() {
	n := scanInt()
	arr := make([]int, n)

	for i := range arr {
		arr[i] = scanInt()
	}

	ans := 1
	if n > 2 {
		direction := judge(arr[0], arr[1])

		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				if direction == "<" {
					ans++
					if i < n-1 {
						direction = judge(arr[i], arr[i+1])
					}
				} else if direction == "=" {
					direction = ">"
				}

			} else if arr[i-1] < arr[i] {
				if direction == ">" {
					ans++
					if i < n-1 {
						direction = judge(arr[i], arr[i+1])
					}
				} else if direction == "=" {
					direction = "<"
				}
			}
		}
	}
	fmt.Println(ans)
}
