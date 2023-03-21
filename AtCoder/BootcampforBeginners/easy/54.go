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

func main() {
	n := scanInt()
	words := make([]string, n)

	for i := range words {
		words[i] = scanString()
	}

	appeared := make(map[string]bool, n)

	ans := "Yes"
	first := words[0]
	appeared[first] = true
	word_end := first[len(first)-1]

	for i := 1; i < n; i++ {
		word := words[i]
		if appeared[word] || word_end != word[0] {
			ans = "No"
			break
		}
		appeared[words[i]] = true
		word_end = word[len(word)-1]
	}

	fmt.Println(ans)
}
