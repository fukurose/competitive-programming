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

func main() {
	s := scanString()
	k := scanInt()
	size := len(s)
	arr1 := strings.Split(s, "")
	arr2 := make([]string, size+1)
	table := make(map[string]int, 26)

	for i, c := range arr1 {
		arr2[i+1] = c
		table[c]++
	}
	if size == 1 || len(table) == 1 {
		fmt.Println(size * k / 2)
	} else {
		counter1 := 0
		for i := 0; i < size-1; i++ {
			if arr1[i] == arr1[i+1] {
				arr1[i+1] = "A"
				counter1++
			}
		}
		arr2[0] = arr1[size-1]

		counter2 := 0
		for i := 0; i < size; i++ {
			if arr2[i] == arr2[i+1] {
				arr2[i+1] = "A"
				counter2++
			}
		}
		ans := counter1 + (counter2 * (k - 1))
		fmt.Println(ans)
	}

}
