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
	x := scanInt()
	arr := make([]int, n)
	for i := range arr {
		arr[i] = scanInt()
	}

	sort.Ints(arr)

	count := 0
	for i := 0; i < n-1; i++ {
		x -= arr[i]
		if x < 0 {
			break
		}
		count++
	}

	if x == arr[n-1] {
		count++
	}

	fmt.Println(count)
}
