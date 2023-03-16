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

	m := make(map[string]string, n)
	ans := "Yes"
	for i := 0; i < n; i++ {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")
		k := inputs[0]
		v := inputs[1]
		for {
			if newVal, ok := m[v]; ok {
				v = newVal
			} else {
				break
			}
		}

		m[k] = v
		if k == v {
			ans = "No"
		}
	}

	fmt.Println(ans)
}
