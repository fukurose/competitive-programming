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

func count(a, b map[string]int) int {
	ans := 0

	for k, v := range a {
		if v < 1 {
			continue
		}

		if b[k] > 0 {
			ans++
		}
	}
	return ans
}

func main() {
	scanInt()
	word := scanString()

	right := make(map[string]int, 26)
	left := make(map[string]int, 26)

	for _, r := range word {
		right[string(r)]++
	}

	ans := -1
	for _, r := range word {

		s := string(r)
		left[s]++
		right[s]--

		mathCount := count(left, right)

		ans = max(ans, mathCount)

	}
	fmt.Println(ans)

}
