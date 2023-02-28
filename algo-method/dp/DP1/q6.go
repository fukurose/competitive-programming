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
	m := scanInt()

	d := make([]int, m)
	ans := make([]bool, n+1)

	for i := range d {
		d[i] = scanInt()
		if d[i] < n+1 {
			ans[d[i]] = true
		}
	}

	for i, v := range ans {
		if v {
			for _, e := range d {
				if i+e < n+1 {
					ans[i+e] = true
				}
			}
		}
	}

	if ans[n] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
