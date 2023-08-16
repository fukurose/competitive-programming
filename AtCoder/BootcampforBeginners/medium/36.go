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

func gcd(a, b int) int {
	if a > b {
		a, b = b, a
	}
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	n := scanInt()
	x := scanInt()

	arr := make([]int, n+1)

	arr[0] = x
	for i := 1; i <= n; i++ {
		arr[i] = scanInt()

	}
	sort.Ints(arr)

	dis := []int{}
	for i := 1; i <= n; i++ {
		d := arr[i] - arr[i-1]
		if d > 0 {
			dis = append(dis, d)
		}
	}

	for len(dis) > 1 {
		f := dis[0]
		s := dis[1]
		dis = dis[2:]

		g := gcd(f, s)
		dis = append(dis, g)
	}

	fmt.Println(dis[0])
}
