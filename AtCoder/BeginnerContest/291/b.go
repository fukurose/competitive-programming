package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n := scanInt()
	points := make([]int, 5*n)
	for i := range points {
		points[i] = scanInt()
	}

	sort.Ints(points)
	size := len(points)
	points = points[n : size-n]
	var total float64 = 0
	for _, point := range points {
		total += float64(point)
	}

	fmt.Println(total / float64(3*n))
}
