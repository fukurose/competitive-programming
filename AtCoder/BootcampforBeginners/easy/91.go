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
	x float64
	y float64
}

func calDistance(a, b Point) float64 {
	x := math.Pow(a.x-b.x, 2)
	y := math.Pow(a.y-b.y, 2)
	return math.Sqrt(x + y)
}

func Perm(points []Point, f func([]Point)) {
	perm(points, f, 0)
}

func perm(points []Point, f func([]Point), i int) {
	if i > len(points) {
		f(points)
		return
	}
	perm(points, f, i+1)
	for j := i + 1; j < len(points); j++ {
		points[i], points[j] = points[j], points[i]
		perm(points, f, i+1)
		points[i], points[j] = points[j], points[i]
	}
}

func main() {
	n := scanInt()
	allPoints := make([]Point, n)

	for i := range allPoints {
		x := scanInt()
		y := scanInt()
		allPoints[i] = Point{x: float64(x), y: float64(y)}
	}

	var totalDis float64 = 0
	var count float64 = 0
	Perm(allPoints, func(points []Point) {
		count++
		for i := 1; i < n; i++ {
			totalDis += calDistance(points[i-1], points[i])
		}
	})

	fmt.Println(totalDis / count)
}
