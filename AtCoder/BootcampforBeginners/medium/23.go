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

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

func init() {
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
}

var inValid int = -1
var blank int = -1

func main() {
	n := scanInt()
	m := scanInt()

	arr := make([]int, n)
	for i := range arr {
		arr[i] = blank // 初期化
	}

	ans := 0
	for i := 0; i < m; i++ {
		s := scanInt() - 1
		c := scanInt()
		if arr[s] != blank && arr[s] != c {
			ans = inValid
		} else {
			arr[s] = c
		}
	}

	// 2桁以上の場合は先頭に0を許可していない
	if n > 1 && arr[0] == 0 {
		ans = inValid
	}

	if ans != inValid {

		// 値の指定がない箇所は最小値で保管
		if arr[0] == blank {
			if n > 1 {
				arr[0] = 1 // 先頭に0を許可していないので1が最小
			} else {
				arr[0] = 0
			}
		}
		for i := 1; i < n; i++ {
			if arr[i] == blank {
				arr[i] = 0
			}
		}

		for i := range arr {
			ans += int(math.Pow10(n-1-i)) * arr[i]
		}

	}
	fmt.Println(ans)
}
