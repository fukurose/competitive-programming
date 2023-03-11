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
	sockets := scanInt()
	goal := scanInt()

	ans := 0
	// すでに差し込み口は1つある
	if goal > 1 {
		ans = 1
		total_sockets := sockets
		for total_sockets < goal {
			total_sockets += (sockets - 1)
			ans++
		}
	}

	fmt.Println(ans)
}
