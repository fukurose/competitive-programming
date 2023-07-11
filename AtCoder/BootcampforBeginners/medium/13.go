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

type Point struct {
	t int
	x int
	y int
}

func calDistance(a, b Point) int {
	d := 0
	if a.x > b.x {
		d += a.x - b.x
	} else {
		d += b.x - a.x
	}

	if a.y > b.y {
		d += a.y - b.y
	} else {
		d += b.y - a.y
	}

	return d
}

func main() {
	n := scanInt()
	points := make([]Point, n+1)

	points[0] = Point{t: 0, x: 0, y: 0}
	for i := 1; i <= n; i++ {
		points[i] = Point{t: scanInt(), x: scanInt(), y: scanInt()}
	}

	isOK := true
	for i := 1; i <= n; i++ {
		time := points[i].t - points[i-1].t
		dis := calDistance(points[i], points[i-1])

		if time < dis {
			// 時間に対して距離の方が多い場合は無条件で不可能
			isOK = false
		} else if dis%2 == 0 && time%2 != 0 {
			// 距離が偶数で、時間が奇数の場合は到達不可
			isOK = false
		} else if dis%2 != 0 && time%2 == 0 {
			// 距離が奇数で、時間が偶数の場合は到達不可
			isOK = false
		}
	}

	if isOK {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
