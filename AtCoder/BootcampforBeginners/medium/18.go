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

func main() {
	m := make([][]int, 3)
	for i := range m {
		m[i] = make([]int, 3)
		for j := range m[i] {
			m[i][j] = scanInt()
		}
	}

	firstTotal := 0
	for _, n := range m[0] {
		firstTotal += n
	}

	secondTotal := 0
	for _, n := range m[1] {
		secondTotal += n
	}
	thirdTotal := 0
	for _, n := range m[2] {
		thirdTotal += n
	}

	abDiff := firstTotal - secondTotal
	bcDiff := secondTotal - thirdTotal
	caDiff := thirdTotal - firstTotal

	isOK := true
	if abDiff%3 != 0 || bcDiff%3 != 0 || caDiff%3 != 0 {
		isOK = false
	}

	if isOK {
		if (m[0][0] - m[1][0]) != abDiff/3 {
			isOK = false
		}
		if (m[0][1] - m[1][1]) != abDiff/3 {
			isOK = false
		}
		if (m[0][2] - m[1][2]) != abDiff/3 {
			isOK = false
		}

		if (m[1][0] - m[2][0]) != bcDiff/3 {
			isOK = false
		}
		if (m[1][1] - m[2][1]) != bcDiff/3 {
			isOK = false
		}
		if (m[1][2] - m[2][2]) != bcDiff/3 {
			isOK = false
		}

		if (m[2][0] - m[0][0]) != caDiff/3 {
			isOK = false
		}
		if (m[2][1] - m[0][1]) != caDiff/3 {
			isOK = false
		}
		if (m[2][2] - m[0][2]) != caDiff/3 {
			isOK = false
		}
	}

	if isOK {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
