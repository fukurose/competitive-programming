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

func main() {
	n := scanInt()

	ans := make([]int, n+1)

	if n == 1 {
		fmt.Println(1)
	} else if n == 2 {
		fmt.Println(2)
	} else {
		ans[0] = 1
		ans[1] = 1
		ans[2] = 2

		for i := 3; i <= n; i++ {
			ans[i] = ans[i-1] + ans[i-2] + ans[i-3]
		}

		fmt.Println(ans[n])

	}

}
