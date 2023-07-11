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

func main() {
	key := "keyence"
	str := scanString()

	isOK := false
	// 完全一致の場合
	if key == str {
		isOK = true
	} else if str[0] == key[0] { // 最初の文字が一致している場合
		matchedCount := 0
		for i, r := range key {
			if string(r) == string(str[i]) {
				matchedCount++
			} else {
				break
			}
		}

		// 末尾が残った文字と一致すればOK
		remainSize := len(key) - matchedCount
		isOK = str[len(str)-remainSize:] == key[matchedCount:]
	} else { // 先頭が一致していなければ
		// 末尾が key になっていればOK
		isOK = string(str[len(str)-len(key):]) == key
	}

	if isOK {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
