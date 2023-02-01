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

	ans := make([]int, 40)
	ans[0] = 1
	ans[1] = 2

	for i := 2; i < len(ans); i++ {
		ans[i] = ans[i-2] + ans[i-1]
	}

	fmt.Println(ans[n-1])
}
