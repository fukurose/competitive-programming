package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func f(arr []int) (int, []int) {
	count := 0
	for i := range arr {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				tmp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = tmp
				count++
			}
		}
	}

	return count, arr
}

func printArray(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println(arr[len(arr)-1])
}

func main() {

	scanner.Scan()
	len, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	arr := make([]int, len)
	for i, s := range strings.Split(scanner.Text(), " ") {
		n, _ := strconv.Atoi(s)
		arr[i] = n
	}

	n, sorted := f(arr)
	printArray(sorted)
	fmt.Println(n)
}
