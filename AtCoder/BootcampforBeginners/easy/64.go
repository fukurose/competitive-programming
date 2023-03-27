package main

import (
	"bufio"
	"bytes"
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
	o := scanString()
	e := scanString()

	e_len := len(e)

	buf := bytes.NewBufferString("")

	for i := range o {
		buf.WriteByte(o[i])
		if i < e_len {
			buf.WriteByte(e[i])
		}
	}

	fmt.Println(buf.String())
}
