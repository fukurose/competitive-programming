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
	n := scanInt()
	a := scanInt()
	b := scanInt()

	all := a + b

	ans := 0
	if n >= all {
		c := n / all
		ans += c * a
		n = n % all
	}

	if n > a {
		ans += a
	} else {
		ans += n
	}

	fmt.Println(ans)
}
