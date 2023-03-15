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

func init() {
	scanner.Split(bufio.ScanWords)
}

func main() {
	a := scanInt()
	scanInt()
	s := scanString()

	isPostCode := true
	_, e := strconv.Atoi(s[0:a])
	if e != nil {
		isPostCode = false
	}

	if s[a:a+1] != "-" {
		isPostCode = false
	}

	_, e = strconv.Atoi(s[a+1:])
	if e != nil {
		isPostCode = false
	}

	if isPostCode {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
