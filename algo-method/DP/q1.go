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

	arr := make([]int, n)
	arr[0] = scanInt()
	arr[1] = scanInt()

	for i := 2; i < n; i++ {
		x := arr[i-2]
		y := arr[i-1]
		arr[i] = (x + y) % 100
	}

	fmt.Println(arr[n-1])
}
