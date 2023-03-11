package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

func init() {
	scanner.Split(bufio.ScanWords)
}

func main() {
	s := scanString()

	for i, r := range s {
		if unicode.IsUpper(r) {
			fmt.Println(i + 1)
			break
		}
	}
}
