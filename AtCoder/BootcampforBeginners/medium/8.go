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

func rateColor(a int) string {
	if a < 400 {
		return "灰"
	} else if a < 800 {
		return "茶"
	} else if a < 1200 {
		return "緑"
	} else if a < 1600 {
		return "水"
	} else if a < 2000 {
		return "青"
	} else if a < 2400 {
		return "黄"
	} else if a < 2800 {
		return "橙"
	} else if a < 3200 {
		return "赤"
	} else {
		return "虹"
	}
}

func main() {
	n := scanInt()
	arr := make([]int, n)

	for i := range arr {
		arr[i] = scanInt()
	}

	colorCount := make(map[string]int, 9)

	for _, rate := range arr {
		color := rateColor(rate)
		colorCount[color]++
	}

	min := 0

	for color, count := range colorCount {
		if color != "虹" && count > 0 {
			min++
		}
	}

	max := min + colorCount["虹"]

	if min < 1 {
		min = 1
	}

	fmt.Println(min, max)

}
