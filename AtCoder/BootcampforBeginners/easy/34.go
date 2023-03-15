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
	limit_500 := scanInt()
	limit_100 := scanInt()
	limit_50 := scanInt()
	x := scanInt()

	ans := 0
	for i := 0; i <= limit_500; i++ {
		for j := 0; j <= limit_100; j++ {
			r := x - (500 * i) - (100 * j)
			if r >= 0 && r <= 50*limit_50 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
