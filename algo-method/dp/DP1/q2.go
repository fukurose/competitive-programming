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

func init() {
	scanner.Split(bufio.ScanWords)
}

func main() {
	n := scanInt()

	arr := make([]int, n)
	ans := make([]int, n)

	for i := range arr {
		arr[i] = scanInt()
	}

	for i := 1; i < n; i++ {
		if i == 1 {
			ans[i] = arr[i]
		} else {
			a := arr[i] + ans[i-1]
			b := arr[i]*2 + ans[i-2]
			if a > b {
				ans[i] = b
			} else {
				ans[i] = a
			}
		}
	}
	fmt.Println(ans[n-1])
}
