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

func isExist(s string, b byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return true
		}
	}
	return false
}

func main() {
	s := scanString()

	table := make(map[byte]int, 26)

	for i := 0; i < len(s); i++ {
		table[s[i]]++
	}

	ans := "-1"
	if len(table) < 26 {
		for i := "a"[0]; i < "a"[0]+26; i++ {
			if table[i] < 1 {
				ans = s + string(i)
				break
			}
		}
	} else {
		for i := len(s) - 1; i >= 0; i-- {
			b := s[i] + 1
			for j := b; j < "a"[0]+26; j++ {
				if isExist(s[i:], j) {
					ans = s[0:i] + string(j)
					break
				}
			}
			if ans != "-1" {
				break
			}
		}
	}

	fmt.Println(ans)

}
