package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	len  int
}

func (l *linkedList) prepend(node *node) {
	if l.head == nil {
		l.head = node
		l.len++
		return
	}

	second := l.head
	l.head = node
	l.head.next = second
	l.len++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.len != 0 {
		fmt.Printf("%d\n", toPrint.data)
		toPrint = toPrint.next
		l.len--
	}
}

func (l *linkedList) deleteWithValue(val int) {
	if l.len == 0 {
		return
	}

	if l.head.data == val {
		l.head = l.head.next
		l.len--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != val {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next
	l.len--
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 20}
	node3 := &node{data: 30}
	node4 := &node{data: 30}
	node5 := &node{data: 76}
	node6 := &node{data: 44}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)

	myList.printListData()
	myList.deleteWithValue(1000)
	myList.deleteWithValue(44)
	myList.printListData()

}
