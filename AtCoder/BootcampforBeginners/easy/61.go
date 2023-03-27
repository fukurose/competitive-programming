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
	s := scanString()
	t := scanString()
	length := len(s)

	ans := "No"
	for i := range s {
		if s[i] == t[0] {
			match_counter := 1
			for j := 1; j < length; j++ {
				if s[(i+j)%length] == t[j] {
					match_counter++
				} else {
					break
				}
			}

			if match_counter == length {
				ans = "Yes"
				break
			}
		}
	}

	fmt.Println(ans)
}
