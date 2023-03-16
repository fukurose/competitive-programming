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
	inputs := strings.Split(scanner.Text(), " ")
	a, _ := strconv.Atoi(inputs[0])
	b, _ := strconv.Atoi(inputs[1])

	if b == a*2 || b == (a*2)+1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
非負整数