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

var white string = "."
var black string = "#"

func isAroundAllWhite(i, j int, arr [][]string) bool {
	if arr[i-1][j] != white {
		return false
	} else if arr[i][j-1] != white {
		return false
	} else if arr[i][j+1] != white {
		return false
	} else if arr[i+1][j] != white {
		return false
	}

	return true
}

func main() {
	h := scanInt()
	w := scanInt()

	arr := make([][]string, h+2)
	arr[0] = make([]string, w+2)
	arr[h+1] = make([]string, w+2)

	for i := 0; i < w+2; i++ {
		arr[0][i] = white
		arr[h+1][i] = white
	}

	for i := 1; i < h+1; i++ {
		str := white + scanString() + white
		arr[i] = strings.Split(str, "")
	}

	isOK := true
	for i := 1; i < h+1; i++ {
		for j := 1; j < w+1; j++ {
			if arr[i][j] == black && isAroundAllWhite(i, j, arr) {
				isOK = false
			}
		}
	}

	if isOK {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
