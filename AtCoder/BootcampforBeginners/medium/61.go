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

func reverse(s string) string {

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	s := scanString()
	n := scanInt()

	ans := s
	rFlg := 0

	for i := 0; i < n; i++ {
		q := scanInt()
		if q == 1 {
			rFlg = (rFlg + 1) % 2
		} else {
			f := scanInt()
			c := scanString()

			if (f+rFlg)%2 == 0 {
				ans = ans + c
			} else {
				ans = c + ans
			}
		}
	}
	if rFlg == 1 {
		ans = reverse(ans)
	}

	fmt.Println(ans)
}
