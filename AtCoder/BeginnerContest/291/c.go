package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
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
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
}

type Point struct {
	x int
	y int
}

func (p *Point) right() {
	p.x++
}

func (p *Point) left() {
	p.x--
}

func (p *Point) up() {
	p.y++
}

func (p *Point) down() {
	p.y--
}

func main() {
	n := scanInt()
	s := scanString()

	visited_points := make(map[Point]bool, n+1)

	point := Point{x: 0, y: 0}
	visited_points[point] = true

	ans := "No"
	for _, r := range s {
		switch string(r) {
		case "R":
			point.right()
		case "L":
			point.left()
		case "U":
			point.up()
		case "D":
			point.down()
		}

		if visited_points[point] {
			ans = "Yes"
			break
		} else {
			visited_points[point] = true
		}
	}

	fmt.Println(ans)
}
