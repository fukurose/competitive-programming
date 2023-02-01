package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type Process struct {
	name string
	time int
}

type Queue []Process

func (q *Queue) enQueue(p Process) {
	*q = append(*q, p)
}

func (q *Queue) deQueue() Process {
	p := (*q)[0]
	*q = (*q)[1:]

	return p
}

func f(q Queue, qTime int) []Process {
	var ans = []Process{}
	var currentTime = 0

	for {
		if len(q) == 0 {
			break
		}

		p := q.deQueue()
		if p.time <= qTime {
			currentTime += p.time
			ans = append(ans, Process{name: p.name, time: currentTime})
		} else {
			currentTime += qTime
			p.time = p.time - qTime
			q.enQueue(p)
		}
	}

	return ans
}

func main() {
	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(inputs[0])
	qTime, _ := strconv.Atoi(inputs[1])

	q := Queue{}

	for i := 0; i < n; i++ {
		scanner.Scan()
		inputs = strings.Split(scanner.Text(), " ")
		name := inputs[0]
		time, _ := strconv.Atoi(inputs[1])

		q.enQueue(Process{name: name, time: time})
	}

	ans := f(q, qTime)
	for _, p := range ans {
		fmt.Print(p.name)
		fmt.Print(" ")
		fmt.Print(p.time)
		fmt.Println()
	}
}
