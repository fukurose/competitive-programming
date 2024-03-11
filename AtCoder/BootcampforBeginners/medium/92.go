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
	m := scanInt()

	a := make([][]int, m)

	for i := range a {
		k := scanInt()
		a[i] = make([]int, k)
		for j := range a[i] {
			a[i][j] = scanInt()
		}
	}

	p := make([]int, m)
	for i := range p {
		p[i] = scanInt()
	}

	ans := 0
	for bits := 0; bits < 1<<uint64(n); bits++ {
		isOK := true
		for i := range a {
			onCounter := 0
			for j := range a[i] {
				if (bits>>uint64(a[i][j]-1))&1 == 1 {
					onCounter++
				}
			}

			if onCounter%2 != p[i] {
				isOK = false
				break
			}
		}
		if isOK {
			ans++
		}
	}

	fmt.Println(ans)
}
