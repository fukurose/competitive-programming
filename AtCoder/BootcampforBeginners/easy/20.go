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

func f(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return (3 * n) + 1

	}
}

func main() {
	s := scanInt()
	arr := make([]int, 1000000)
	m := make(map[int]bool, 1000000)

	arr[0] = s
	m[s] = true

	for i := 1; i < len(arr); i++ {
		arr[i] = f(arr[i-1])
		if m[arr[i]] {
			fmt.Println(i + 1)
			break
		} else {
			m[arr[i]] = true
		}
	}

}
