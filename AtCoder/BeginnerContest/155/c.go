package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

var m = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

func f(s string) int {
	arr := strings.Split(s, "")
	ans := 0
	m_len := len(m)
	s_len := len(s)
	for i, v := range arr {
		n := m[v]
		keta := s_len - i - 1
		if keta == 0 {
			ans = ans + n
		} else {
			ans = ans + int(math.Pow(float64(m_len), float64(keta)))*n
		}
	}
	return ans
}

func main() {
	scanner.Scan()
	fmt.Println(f(scanner.Text()))
}
