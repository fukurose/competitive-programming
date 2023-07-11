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

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

func init() {
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
}

func main() {
	h := scanInt()
	w := scanInt()

	boms := make([][]bool, h+2)
	ans := make([][]int, h+2)
	for i := range boms {
		boms[i] = make([]bool, w+2)
		ans[i] = make([]int, w+2)
	}

	for i := 1; i <= h; i++ {
		s := scanString()
		for j, r := range s {
			if string(r) == "#" {
				boms[i][j+1] = true
			}
		}
	}

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			// ボムの場合は -1 を設定
			if boms[i][j] {
				ans[i][j] = -1
				continue
			}

			// ボムじゃない場合は周りをカウント
			if boms[i-1][j-1] {
				ans[i][j]++
			}
			if boms[i-1][j] {
				ans[i][j]++
			}
			if boms[i-1][j+1] {
				ans[i][j]++
			}

			if boms[i][j-1] {
				ans[i][j]++
			}
			if boms[i][j] {
				ans[i][j]++
			}
			if boms[i][j+1] {
				ans[i][j]++
			}

			if boms[i+1][j-1] {
				ans[i][j]++
			}
			if boms[i+1][j] {
				ans[i][j]++
			}
			if boms[i+1][j+1] {
				ans[i][j]++
			}
		}
	}

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if ans[i][j] < 0 {
				fmt.Print("#")
			} else {
				fmt.Print(ans[i][j])
			}
		}
		fmt.Println()
	}
}
