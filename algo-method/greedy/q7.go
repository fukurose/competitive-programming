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
func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func init() {
	scanner.Split(bufio.ScanWords)
}

type Task struct {
	start int
	end   int
}

func filter(arr []Task, limit int) []Task {
	var tasks []Task
	for _, t := range arr {
		if t.start >= limit {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

func main() {
	n := scanInt()
	tasks := make([]Task, n)

	for i := range tasks {
		s := scanInt()
		e := scanInt()
		tasks[i] = Task{start: s, end: e}
	}

	count := 0
	for len(tasks) > 0 {
		sort.Slice(tasks, func(i, j int) bool { return tasks[i].end < tasks[j].end })

		task := tasks[0]
		count++
		tasks = filter(tasks, task.end)
	}
	fmt.Println(count)
}
