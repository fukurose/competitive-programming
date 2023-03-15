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
	arr := make([]int, 5)
	for i := range arr {
		arr[i] = scanInt()
	}
	sort.Slice(arr, func(i, j int) bool {
		a := arr[i] % 10
		if a == 0 {
			a = 10
		}
		b := arr[j] % 10
		if b == 0 {
			b = 10
		}
		return a > b
	})

	count := 0
	for i, v := range arr {
		count += v
		if v%10 == 0 {
			// 10 の倍数の時刻 の場合は次の注文
			continue
		}
		if i < 4 {
			// 最後の注文は待つ必要なし
			count += 10 - v%10
		}
	}
	fmt.Println(count)
}
