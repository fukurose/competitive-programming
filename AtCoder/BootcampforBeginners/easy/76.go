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

type Anser struct {
	n    int
	isAC bool
}

func main() {
	n := scanInt()
	m := scanInt()

	ansers := make([]Anser, m)

	for i := range ansers {
		p := scanInt()
		s := scanString()
		ansers[i] = Anser{n: p - 1, isAC: s == "AC"}

	}

	ac := make([]int, n)
	wa := make([]int, n)

	for _, anser := range ansers {
		if anser.isAC {
			ac[anser.n] = 1
		} else if ac[anser.n] < 1 {
			// 未正解の問題で間違った場合
			wa[anser.n]++
		}
	}

	ac_count := 0
	wa_count := 0

	for _, n := range ac {
		ac_count += n
	}

	for i, n := range wa {
		if ac[i] < 1 {
			// 未正解のものは集計しない
			continue
		}
		wa_count += n
	}

	fmt.Println(ac_count, wa_count)
}
