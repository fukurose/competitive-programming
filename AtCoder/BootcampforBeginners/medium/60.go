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
	n := scanInt()
	counter := make(map[byte]int, 5)
	total := 0

	for i := 0; i < n; i++ {
		s := scanString()
		if s[0] == 'M' || s[0] == 'A' || s[0] == 'R' || s[0] == 'C' || s[0] == 'H' {
			counter[s[0]]++
			total++
		}

	}
	ans := 0
	if len(counter) >= 3 {
		// 重複を排除せずに、3人選ぶ組み合わせ
		ans = (total) * (total - 1) * (total - 2) / 6
		// 重複がある組み合わせ
		dep := 0
		for _, c := range counter {
			if c >= 2 {
				// 2つ重複している場合
				dep += (c) * (c - 1) / 2 * (total - c) * 1
				if c >= 3 {
					// すべて重複している場合
					dep += (c) * (c - 1) * (c - 2) / 6
				}
			}

		}
		ans -= dep
	}
	fmt.Println(ans)
}
