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

func init() {
	scanner.Split(bufio.ScanWords)
}

func main() {
	n := scanInt()
	p := ""
	for i := 0; i < n; i++ {
		p += scanString()
	}
	q := ""
	for i := 0; i < n; i++ {
		q += scanString()
	}
	fmt.Println(p)
	fmt.Println(q)

}
