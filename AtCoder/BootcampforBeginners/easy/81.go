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

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func filterEven(arr []int) []int {
	var evenArr []int
	for _, n := range arr {
		if n%2 == 0 {
			evenArr = append(evenArr, n)
		}
	}

	return evenArr
}

func IndexOfMaxValue(arr []int) int {
	index := 0
	maximum := arr[0]
	for i := range arr {
		if arr[i] > maximum {
			index = i
		}
	}

	return index
}

func main() {
	n := scanInt()
	arr := make([]int, n)

	for i := range arr {
		arr[i] = scanInt()
	}

	ans := 0
	for i := range arr {
		for {
			if arr[i]%2 != 0 || arr[i] < 2 {
				break
			}

			arr[i] /= 2
			ans++

		}
	}

	fmt.Println(ans)
}
