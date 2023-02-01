package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	num, _ := strconv.Atoi(scanString())
	return num
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

var dummy = &Node{}

type Node struct {
	key        int
	next, prev *Node
}

func insert(key int) {
	node := &Node{key: key}
	first := dummy.next

	first.prev = node
	node.next = first

	node.prev = dummy
	dummy.next = node
}

func delete(key int) {
	node := search(key)
	if node == dummy {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

func deleteFirst() {
	second := dummy.next.next
	dummy.next = second
	second.prev = dummy
}

func deleteLast() {
	secondLast := dummy.prev.prev
	secondLast.next = dummy
	dummy.prev = secondLast
}

func search(key int) *Node {
	node := dummy.next
	for {
		if node.key == key {
			return node
		}
		node = node.next
		if node == dummy {
			break
		}
	}
	return dummy
}

func print() string {
	buf := bytes.Buffer{}
	current := dummy.next
	for {
		buf.WriteString(strconv.Itoa(current.key))
		current = current.next
		if current == dummy {
			break
		}
		buf.WriteString(" ")
	}
	return buf.String()
}

func init() {
	sc.Split(bufio.ScanWords)
	dummy.prev = dummy
	dummy.next = dummy
}

func main() {
	n := scanInt()

	for i := 0; i < n; i++ {
		command := scanString()
		switch command {
		case "insert":
			insert(scanInt())
		case "delete":
			delete(scanInt())
		case "deleteFirst":
			deleteFirst()
		case "deleteLast":
			deleteLast()
		}
	}

	fmt.Println(print())
}
