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

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

func init() {
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
}

type Shop struct {
	price int
	count int
}

func main() {
	n := scanInt()
	m := scanInt()

	shopList := make([]Shop, n)

	for i := range shopList {
		price := scanInt()
		count := scanInt()
		shopList[i] = Shop{price: price, count: count}
	}

	sort.Slice(shopList, func(i, j int) bool { return shopList[i].price < shopList[j].price })

	ans := 0
	total := 0
	for _, shop := range shopList {
		r := m - total
		if r == 0 {
			break
		}

		if shop.count <= r {
			ans += shop.count * shop.price
			total += shop.count
		} else {
			ans += r * shop.price
			total += r
		}
	}
	fmt.Println(ans)

}
