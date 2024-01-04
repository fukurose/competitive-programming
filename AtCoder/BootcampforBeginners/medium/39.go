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
	k := scanInt()
	a := scanInt()
	b := scanInt()

	biscuit := 1
	if b-a > 2 {
		count := 0
		for count < k {
			if biscuit >= a && k-count >= 2 {
				// ビスケット A 枚を 1 円に交換する
				biscuit -= a
				count++
				// 1 円をビスケット B 枚に交換する
				biscuit += b
				count++
			} else {
				// 持っているビスケットを叩き、1 枚増やす
				biscuit++
				count++
			}
		}
	} else {
		biscuit += k
	}

	fmt.Println(biscuit)
}
