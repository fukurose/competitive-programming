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
	arr := make([]int, n+1)

	for i := 1; i <= n; i++ {
		arr[scanInt()] = i
	}

	for i := 1; i < n; i++ {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
	fmt.Println(arr[n])

}
