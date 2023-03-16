package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	arr := strings.Split(scanner.Text(), "")

	for i := 1; i < n; i++ {
		l := 0
		for j := 0; i+j < n; j++ {
			if arr[j] == arr[i+j] {
				break
			} else {
				l = j + 1
			}
		}
		fmt.Println(l)
	}
}
