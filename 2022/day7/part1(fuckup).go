package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	value     string
	directory bool
	weight    int
	sub       *LinkedList
	next      *Node
	prev      *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(val string) {
	n := Node{}
	n.value = val
	n.sub = nil
	n.prev = nil
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			n.prev = ptr
			l.len++
			return
		}
		ptr = ptr.next
	}
}

// InsertAt inserts new node at given position
func (l *LinkedList) InsertAt(pos int, value string) {
	// create a new node
	newNode := Node{}
	newNode.value = value
	// validate the position
	if pos < 0 {
		return
	}
	if pos == 0 {
		l.head = &newNode
		l.len++
		return
	}
	if pos > l.len {
		return
	}
	n := l.GetAt(pos)
	newNode.next = n
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++
}

// Print displays all the nodes from linked list
func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println(ptr.value)
		ptr = ptr.next
	}
}

// Search returns node position with given value from linked list
func (l *LinkedList) Search(val string) int {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			return i
		}
		ptr = ptr.next
	}
	return -1
}

// GetAt returns node at given position from linked list
func (l *LinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

func main() {
	input, err := os.ReadFile("./inputTest")
	if err != nil {
		panic(err)
	}
	inputString := string(input)
	commandesSlice := strings.Split(inputString, "\n")

	root := LinkedList{}
	var currentDir *Node
	for _, commande := range commandesSlice {
		if commande == "$ cd /" {
			root.Insert(commande)
			currentDir = root.head
			continue
		}
		if commande == "$ ls" {
			continue
		}
		if strings.HasPrefix(commande, "dir") {
			currentDir.sub = &LinkedList{}
			currentDir.directory = true
			currentDir.sub.Insert(commande)
			continue
		}
	}

	root.Print()
	root.head.sub.Print()
}
