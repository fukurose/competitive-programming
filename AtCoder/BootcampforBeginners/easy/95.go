package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func isAllSame(arr []int) bool {
	return arr[0] == arr[1] && arr[1] == arr[2]
}

func main() {
	arr := make([]int, 3)

	for i := range arr {
		arr[i] = scanInt()
	}

	ans := 0
	for !isAllSame(arr) {
		sort.Ints(arr)
		diff := arr[2] - arr[0]
		if diff >= 2 {
			arr[0] += 2
		} else {
			arr[0]++
			arr[1]++
		}
		ans++
	}
	fmt.Println(ans)

}
