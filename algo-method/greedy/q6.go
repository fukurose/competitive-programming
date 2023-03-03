package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var INF float64 = 9999999999

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

type Point struct {
	x float64
	y float64
}

func calDistance(a, b Point) float64 {
	x := math.Pow(a.x-b.x, 2)
	y := math.Pow(a.y-b.y, 2)
	return math.Sqrt(x + y)
}

func main() {
	n := scanInt()
	points := make([]Point, n)
	arrived := make([]bool, n)

	for i := range points {
		x := float64(scanInt())
		y := float64(scanInt())
		points[i] = Point{x: x, y: y}
	}

	now_p := points[0]
	arrived[0] = true

	var cost float64 = 0
	for i := 1; i < n; i++ {
		distance := INF
		next_p := 0
		for j, p := range points {
			if arrived[j] {
				continue
			}

			if distance > calDistance(now_p, p) {
				distance = calDistance(now_p, p)
				next_p = j
			}
		}

		cost += distance
		now_p = points[next_p]
		arrived[next_p] = true
	}

	cost += calDistance(now_p, points[0])
	fmt.Println(cost)

}
