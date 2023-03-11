package main

import (
	"bufio"
	"fmt"
	"math"
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

func min(a, b float64) float64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	n := scanInt()
	points := make([]int, n)
	for i := range points {
		points[i] = scanInt()
	}

	ans := math.Inf(0)
	for i := 1; i <= 100; i++ {
		var score float64 = 0
		for _, p := range points {
			score += math.Pow(float64(i-p), 2)
		}
		ans = min(ans, score)
	}
	fmt.Println(ans)
}
