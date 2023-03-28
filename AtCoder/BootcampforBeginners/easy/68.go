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

type Restaurant struct {
	city      string
	point     int
	init_sort int
}

func main() {
	n := scanInt()
	restaurants := make([]Restaurant, n)

	for i := range restaurants {
		city := scanString()
		point := scanInt()
		restaurants[i] = Restaurant{city: city, point: point, init_sort: i + 1}
	}

	sort.Slice(restaurants, func(i, j int) bool {
		if restaurants[i].city == restaurants[j].city {
			return restaurants[i].point > restaurants[j].point
		} else {
			return restaurants[i].city < restaurants[j].city
		}
	})

	for _, restaurant := range restaurants {
		fmt.Println(restaurant.init_sort)
	}

}
