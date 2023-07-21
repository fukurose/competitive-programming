package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	s := strings.Split(scanString(), "")
	t := strings.Split(scanString(), "")

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	sort.Slice(t, func(i, j int) bool {
		return t[i] > t[j]
	})

	if strings.Join(s, "") < strings.Join(t, "") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
