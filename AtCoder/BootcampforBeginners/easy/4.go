package main

import (
	"bufio"
	"fmt"
	"math"
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

	ans := -1
	for i := 1; i <= 50000; i++ {
		if math.Floor(float64(i)*1.08) == float64(n) {
			ans = i
		}
	}

	if ans < 0 {
		fmt.Println(":(")
	} else {
		fmt.Println(ans)
	}

}
