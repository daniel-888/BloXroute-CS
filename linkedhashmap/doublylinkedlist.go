package linkedhashmap

import (
	"fmt"
	"strings"
)

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	len int
	tail *Node
	head *Node
}

func initDoublyList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (list *DoublyLinkedList) Add(values ...interface{}) {
	for _, value := range values {
		newNode := &Node{data: value, prev: list.tail}
		if list.len == 0 {
			list.head = newNode
			list.tail = newNode
		} else {
			list.tail.next = newNode
			list.tail = newNode
		}
		list.len++
	}
}

func (list *DoublyLinkedList) Append(values ...interface{}) {
	list.Add(values...)
}

func (list *DoublyLinkedList) Prepend(values ...interface{}) {
	for v := len(values) - 1; v >=0; v-- {
		newNode := &Node{data: values[v], next: list.head}
		if list.len == 0{
			list.head = newNode
			list.tail = newNode
		} else {
			list.head.prev = newNode
			list.head = newNode
		}
		list.len++
	}
}

func (list *DoublyLinkedList) Get(index int) (interface{}, bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	// determine search direction
	if list.len - index < index {
		node := list.tail
		for n := list.len - 1; n != index; n, node = n-1, node.prev {}
		return node.data, true
	}
	node := list.head
	for n := 0; n != index; n, node = n+1, node.next {}
	return node.data, true
}

func (list *DoublyLinkedList) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	if list.len == 1 {
		list.Clear()
		return
	}

	var node *Node
	// determine search direction
	if list.len-index < index {
		node = list.tail
		for e := list.len - 1; e != index; e, node = e-1, node.prev {
		}
	} else {
		node = list.head
		for e := 0; e != index; e, node = e+1, node.next {
		}
	}

	if node == list.head {
		list.head = node.next
	}
	if node == list.tail {
		list.tail = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	node = nil

	list.len--
}

func (list *DoublyLinkedList) Values() []interface{} {
	values := make([]interface{}, list.len)
	for n, node := 0, list.head; node != nil; n, node = n+1, node.next {
		values[n] = node.data
	}
	return values
}

func (list *DoublyLinkedList) IndexOf(value interface{}) int {
	if list.len == 0 {
		return -1
	}
	for index, node := range list.Values() {
		if node == value {
			return index
		}
	}
	return -1
}

func (list *DoublyLinkedList) String() string {
	str := "DoublyLinkedList\n"
	values := []string{}
	for node := list.head; node != nil; node = node.next {
		values = append(values, fmt.Sprintf("%v", node.data))
	}
	str += strings.Join(values, ", ")
	return str
}

func (list *DoublyLinkedList) Clear() {
	list.len = 0
	list.head = nil
	list.tail = nil
}

func (list *DoublyLinkedList) Size() int {
	return list.len
}

func (list *DoublyLinkedList) withinRange(index int) bool {
	return index >= 0 && index < list.len
}