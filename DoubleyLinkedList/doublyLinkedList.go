package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	next *Node
	prev *Node
	key int
}

type DoublyLinkedList struct {
	firstNode *Node
	lastNode *Node
}

func (list *DoublyLinkedList) Insert(v int) {
	node := Node{list.firstNode, nil, v}
	if list.Empty() {
		list.lastNode = &node
	} else {
		list.firstNode.prev = &node
	}
	list.firstNode = &node
}

func (list *DoublyLinkedList) Delete(v int) {
	for node := list.firstNode; node != nil; node = node.next {
		if node.key == v {
			if node.prev == nil {
				list.firstNode = node.next
			} else {
				node.prev.next = node.next
			}
			if node.next == nil {
				list.lastNode = node.prev
			} else {
				node.next.prev = node.prev
			}
			return
		}
	}
}

func (list *DoublyLinkedList) DeleteFirst() {
	if list.Size() <= 1 {
		list.firstNode = nil
		list.lastNode = nil
	} else {
		list.firstNode = list.firstNode.next
		list.firstNode.prev = nil
	}
}

func (list *DoublyLinkedList) DeleteLast() {
	if list.Size() <= 1 {
		list.firstNode = nil
		list.lastNode = nil
	} else {
		list.lastNode = list.lastNode.prev
		list.lastNode.next = nil
	}
}

func (list *DoublyLinkedList) Size() int {
	cnt := 0
	for node := list.firstNode; node != nil; node = node.next {
		cnt++
	}
	return cnt
}

func (list *DoublyLinkedList) Empty() bool {
	return list.Size() == 0
}

func (list * DoublyLinkedList) Dump() {
	for node := list.firstNode; node != nil; node = node.next {
		fmt.Printf("%d ", node.key)
	}
	fmt.Print("\n")
}

func NewDoublyLinkedList() *DoublyLinkedList {
	list := new(DoublyLinkedList)
	return list
}

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, err := strconv.Atoi(next())
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	list := NewDoublyLinkedList()
	fmt.Println("Empty: ", list.Empty())
	for i := 5; i > 0; i-- {
		list.Insert(i)
	}
	list.Dump()
	fmt.Println("Size: ", list.Size())
	fmt.Println("Empty: ", list.Empty())
	fmt.Println("first: ", *list.firstNode, "last: ", *list.lastNode)
	list.Delete(3)
	list.DeleteLast()
	list.DeleteFirst()
	list.Dump()

	//sc.Split(bufio.ScanWords)
	//n := nextInt()
	//
	//var s []int
	//for i := 0; i < n; i++ {
	//	tmp := nextInt()
	//	s = append(s, tmp)
	//}
	//
	//fmt.Println(n, s)
}
