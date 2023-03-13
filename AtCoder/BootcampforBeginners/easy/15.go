package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for i := range arr {
		arr[i] = scanInt()
	}

	sort.Ints(arr)

	ans := arr[n/2] - arr[n/2-1]
	if ans <= 0 {
		ans = 0
	}
	fmt.Println(ans)
}
