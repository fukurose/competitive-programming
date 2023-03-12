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
	a := scanString()
	b := scanString()

	n, _ := strconv.Atoi(a + b)

	ans := "No"
	limit := int(math.Sqrt(float64(n)))
	for i := 1; i <= limit; i++ {
		if i*i == n {
			ans = "Yes"
		}
	}
	fmt.Println(ans)
}
