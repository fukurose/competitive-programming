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
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	n := scanInt()

	up := make([]int, n)
	down := make([]int, n)
	upSum := make([]int, n)
	downSum := make([]int, n)

	upTotal := 0
	for i := range up {
		up[i] = scanInt()
		upTotal += up[i]
		upSum[i] = upTotal
	}

	downTotal := 0
	for i := range down {
		down[i] = scanInt()
		downTotal += down[i]
		downSum[i] = downTotal
	}

	ans := upSum[0] + downSum[n-1]
	for i := 1; i < n; i++ {
		upCandyCount := upSum[i]
		downCandyCount := downSum[n-1] - downSum[i-1]
		ans = max(ans, upCandyCount+downCandyCount)
	}

	fmt.Println(ans)
}
