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

var even int = 0
var odd int = 1

func main() {
	n := scanInt()
	m := scanInt()
	w := make([]int, n)
	for i := range w {
		w[i] = scanInt()
	}

	c := make(map[int][]int, 256)
	for _, v := range w {
		color := scanInt()
		c[color] = append(c[color], v)

	}

	dp := make([][]bool, 256)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	for i := 0; i < 255; i++ {
		for j := range dp[i] {
			if !dp[i][j] {
				continue
			}

			//ボールを取らない場合
			dp[i+1][j] = true

			//ボールを取る場合
			for _, v := range c[i] {
				if v+j <= m {
					dp[i+1][j+v] = true
				}
			}
		}
	}

	if dp[255][m] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
