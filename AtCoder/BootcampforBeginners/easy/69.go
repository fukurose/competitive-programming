package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	s := scanString()
	arr := strings.Split(s, "")

	pre := ""
	cur := ""
	ans := 0
	for _, r := range arr {
		cur += string(r)
		if pre != cur {
			pre = cur
			cur = ""
			ans++
		}
	}

	fmt.Println(ans)
}
