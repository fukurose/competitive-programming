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
	m := scanInt()
	c := scanInt()

	b := make([]int, m)
	for i := range b {
		b[i] = scanInt()
	}
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			a[i][j] = scanInt()
		}
	}

	count := 0
	for i := range a {
		sum := 0
		for j := range a[i] {
			sum += (a[i][j] * b[j])
		}
		if sum+c > 0 {
			count++
		}
	}

	fmt.Println(count)
}
