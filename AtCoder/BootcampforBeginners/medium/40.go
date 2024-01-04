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

func main() {
	n := scanInt()
	n--
	m := scanInt()
	arr := make([]int, m)
	for i := range arr {
		arr[i] = scanInt()
	}

	sort.Ints(arr)

	dis := make([]int, m-1)
	for i := 1; i < m; i++ {
		dis[i-1] = arr[i] - arr[i-1]
	}
	sort.Ints(dis)

	ans := 0
	for i := 0; i < len(dis)-n; i++ {
		ans += dis[i]
	}
	fmt.Println(ans)
}
