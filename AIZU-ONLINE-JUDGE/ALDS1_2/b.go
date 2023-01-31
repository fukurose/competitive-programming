package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func printArray(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println(arr[len(arr)-1])
}

func f(arr []int) (int, []int) {
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		min := arr[i+1]
		min_i := i + 1
		for j := i + 2; j < len(arr); j++ {
			if min > arr[j] {
				min = arr[j]
				min_i = j
			}
		}

		if arr[i] > min {
			tmp := arr[i]
			arr[i] = min
			arr[min_i] = tmp
			count++
		}
	}

	return count, arr
}

func main() {
	scanner.Scan()
	len, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, len)
	scanner.Scan()
	for i, v := range strings.Split(scanner.Text(), " ") {
		n, _ := strconv.Atoi(v)
		arr[i] = n
	}

	n, sorted := f(arr)

	printArray(sorted)
	fmt.Println(n)
}
