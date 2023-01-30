package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func ScanInit() int {
	scanner.Scan()
	input := scanner.Text()
	i, _ := strconv.Atoi(input)
	return i
}

func printArray(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println(arr[len(arr)-1])
}

func main() {
	len := ScanInit()
	arr := make([]int, len)

	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	for i := 0; i < len; i++ {
		num, _ := strconv.Atoi(inputs[i])
		arr[i] = num
	}

	printArray(arr)

	for i := 1; i < len; i++ {
		e := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > e {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = e

		printArray(arr)
	}
}
