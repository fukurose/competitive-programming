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
	nums := make([]int, n)

	sum := 0
	for i := range nums {
		nums[i] = scanInt()
		sum += nums[i]
	}

	sort.Ints(nums)

	alice_score := 0
	for i := n - 1; i >= 0; {
		alice_score += nums[i]
		i -= 2
	}
	bob_score := sum - alice_score

	fmt.Println(alice_score - bob_score)
}
