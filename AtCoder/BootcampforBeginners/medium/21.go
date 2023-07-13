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
	counter := make(map[int]int, n)

	for i := range arr {
		arr[i] = scanInt()
		counter[arr[i]]++
	}

	combinCounter := make([]int, 300000)
	for i := 0; i < 300000-1; i++ {
		combinCounter[i+1] = combinCounter[i] + i
	}

	totalCombin := 0
	for _, v := range counter {
		totalCombin += combinCounter[v]
	}

	for _, v := range arr {
		ans := totalCombin - combinCounter[counter[v]]
		ans += combinCounter[counter[v]-1]
		fmt.Println(ans)
	}

}
