package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

	a_index := -1
	c_index := -1
	check := true
	for i, r := range s {
		if string(r) == "A" {
			if a_index != -1 {
				check = false
				break
			}
			a_index = i
		} else if string(r) == "C" {
			if c_index != -1 {
				check = false
				break
			}
			c_index = i
		} else if unicode.IsUpper(r) {
			check = false
			break

		}
	}

	// 大文字Aは先頭のみ
	if a_index != 0 {
		check = false
	}

	// 大文字Cは3文字目と末尾から2文字目の間（両端含む）
	if c_index < 2 || c_index > len(s)-2 {
		check = false
	}

	if check {
		fmt.Println("AC")
	} else {
		fmt.Println("WA")
	}
}
