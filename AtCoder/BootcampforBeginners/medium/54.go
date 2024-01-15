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
	c := scanInt()
	k := scanInt()

	arr := make([]int, n)
	for i := range arr {
		arr[i] = scanInt()
	}
	sort.Ints(arr)

	// 一番最初に到着した人
	ans := 1
	capa := c - 1
	limit := arr[0] + k

	// それ以降の人を順次
	for i := 1; i < n; i++ {
		arrivedTime := arr[i]
		// 枠があっても、バスは出発している
		if arrivedTime > limit {
			capa = 0
		}

		// 乗れない場合
		if capa == 0 {
			ans++
			capa = c - 1
			limit = arrivedTime + k
		} else {
			capa--
		}
	}
	fmt.Println(ans)
}
