package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func replace(arr [][]string) [][]string {
	replace_arr := make([][]string, len(arr[0]))
	for i := range replace_arr {
		replace_arr[i] = make([]string, len(arr))
	}

	for i := range arr[0] {
		for j := range arr {
			replace_arr[i][j] = arr[j][i]
		}
	}

	return replace_arr
}

func dedupeBlank(arr [][]string) [][]string {
	var okLine []int
	for i := range arr {
		isAllBlank := true
		for j := range arr[i] {
			if arr[i][j] != "." {
				isAllBlank = false
				break
			}
		}

		if !isAllBlank {
			okLine = append(okLine, i)
		}
	}

	dedupedArr := make([][]string, len(okLine))
	for i, line := range okLine {
		dedupedArr[i] = arr[line]
	}

	return dedupedArr
}

func main() {
	h := scanInt()
	w := scanInt()

	arr := make([][]string, h)

	for i := range arr {
		arr[i] = make([]string, w)
		arr[i] = strings.Split(scanString(), "")
	}

	for {
		currentH := len(arr)
		currentW := len(arr[0])

		arr = dedupeBlank(arr)
		arr = replace(arr)
		arr = dedupeBlank(arr)
		arr = replace(arr)

		if currentH == len(arr) && currentW == len(arr[0]) {
			break
		}
	}

	for i := range arr {
		fmt.Println(strings.Join(arr[i], ""))
	}
}
