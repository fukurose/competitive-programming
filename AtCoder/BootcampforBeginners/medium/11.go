package main

import (
	"bufio"
	"fmt"
	"math"
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
	x int
	y int
}

func calDis(a, b Point) int {
	xDis := math.Abs(float64((a.x - b.x)))
	yDis := math.Abs(float64((a.y - b.y)))

	return int(xDis + yDis)
}

func main() {
	n := scanInt()
	m := scanInt()

	studentPoints := make([]Point, n)
	checkPoints := make([]Point, m)

	for i := range studentPoints {
		x := scanInt()
		y := scanInt()

		studentPoints[i] = Point{x: x, y: y}
	}

	for i := range checkPoints {
		x := scanInt()
		y := scanInt()

		checkPoints[i] = Point{x: x, y: y}
	}

	ans := make([]int, n)
	for i, sPoint := range studentPoints {
		minDis := calDis(sPoint, checkPoints[0])
		ans[i] = 1
		for j, cPoint := range checkPoints {
			dis := calDis(sPoint, cPoint)

			if minDis > dis {
				minDis = dis
				ans[i] = j + 1
			}
		}
	}

	for _, pointIndex := range ans {
		fmt.Println(pointIndex)
	}
}
