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
	s := scanString()
	done := make(map[string]bool, 26)

	ans := "yes"
	for _, r := range s {
		if done[string(r)] {
			ans = "no"
			break
		}

		done[string(r)] = true
	}

	fmt.Println(ans)
}
