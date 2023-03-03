package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func init() {
	scanner.Split(bufio.ScanWords)
}

type Task struct {
	time  int
	limit int
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	n := scanInt()
	tasks := make([]Task, n)

	for i := range tasks {
		tasks[i] = Task{time: scanInt()}
	}

	for i := range tasks {
		tasks[i].limit = scanInt()
	}

	sort.Slice(tasks, func(i, j int) bool { return tasks[i].limit < tasks[j].limit })

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 10000)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	dp[0][0] = 0
	dp[0][tasks[0].time] = 1

	for i := 1; i < n; i++ {
		for j := range dp[i] {
			if dp[i-1][j] < 0 {
				continue
			}
			dp[i][j] = max(dp[i][j], dp[i-1][j])
			if j+tasks[i].time <= tasks[i].limit {
				dp[i][j+tasks[i].time] = max(dp[i][j+tasks[i].time], dp[i-1][j]+1)
			}
		}
	}

	ans := 0
	for _, v := range dp[n-1] {
		ans = max(ans, v)
	}

	fmt.Println(ans)
}
