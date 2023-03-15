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
	isOdd := true
	arr := make([]int, n)
	for i := range arr {
		arr[i] = scanInt()
		if arr[i]%2 != 0 {
			isOdd = false
		}
	}

	ans := 0
	for isOdd {
		ans++
		for i := range arr {
			arr[i] /= 2
			if arr[i]%2 != 0 {
				isOdd = false
			}
		}
	}

	fmt.Println(ans)
}
