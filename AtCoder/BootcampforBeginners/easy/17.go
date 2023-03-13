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
	v_arr := make([]int, n)

	for i := range v_arr {
		v_arr[i] = scanInt()
	}

	sort.Ints(v_arr)

	ans := float64(v_arr[0]+v_arr[1]) / 2
	for i := 2; i < n; i++ {
		ans = (ans + float64(v_arr[i])) / 2
	}

	fmt.Println(ans)
}
