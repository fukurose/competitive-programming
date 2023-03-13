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

func eatenDays(days, interval int) int {
	count := 0
	for i := 1; i <= days; i += interval {
		count++
	}
	return count
}

func main() {
	n := scanInt()
	d := scanInt()
	x := scanInt()
	arr := make([]int, n)

	for i := range arr {
		arr[i] = scanInt()
	}

	eaten := 0
	for i := range arr {
		eaten += eatenDays(d, arr[i])
	}

	fmt.Println(eaten + x)

}
