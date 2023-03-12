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

func isBingo(a, b, c int) bool {
	return (a + b + c) == 0
}

func main() {
	bingo := make([]int, 9)
	for i := range bingo {
		bingo[i] = scanInt()
	}

	limit := scanInt()

	for i := 0; i < limit; i++ {
		n := scanInt()
		for i, v := range bingo {
			if v == n {
				bingo[i] = 0 // ビンゴしたら0にしておく
			}
		}
	}

	ans := "No"
	// 横判定
	if isBingo(bingo[0], bingo[1], bingo[2]) || isBingo(bingo[3], bingo[4], bingo[5]) || isBingo(bingo[6], bingo[7], bingo[8]) {
		ans = "Yes"
	}
	// 縦判定
	if isBingo(bingo[0], bingo[3], bingo[6]) || isBingo(bingo[1], bingo[4], bingo[7]) || isBingo(bingo[2], bingo[5], bingo[8]) {
		ans = "Yes"
	}

	// ななめ判定
	if isBingo(bingo[0], bingo[4], bingo[8]) || isBingo(bingo[2], bingo[4], bingo[6]) {
		ans = "Yes"
	}

	fmt.Println(ans)
}
