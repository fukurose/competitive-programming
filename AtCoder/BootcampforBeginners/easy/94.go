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
	word := scanString()
	a, _ := strconv.Atoi(string(word[0]))
	b, _ := strconv.Atoi(string(word[1]))
	c, _ := strconv.Atoi(string(word[2]))
	d, _ := strconv.Atoi(string(word[3]))

	abcd := [4]int{a, b, c, d}

	size := 3
	var ans int
	for bits := 0; bits < (1 << uint64(size)); bits++ {
		sum := abcd[0]
		for i := 0; i < size; i++ {
			if (bits>>uint64(i))&1 == 1 {
				sum += abcd[i+1]
			} else {
				sum -= abcd[i+1]
			}
		}
		if sum == 7 {
			ans = bits
		}
	}

	fmt.Print(abcd[0])
	for i := 0; i < size; i++ {
		if (ans>>uint64(i))&1 == 1 {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
		fmt.Print(abcd[i+1])
	}
	fmt.Println("=7")
}
