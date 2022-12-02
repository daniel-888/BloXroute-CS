package linkedhashmap

import ()

type Iterator struct {
	table map[interface{}]interface{}
	list *DoublyLinkedList
	index int
	node *Node
}

func (m *Map) Iterator() Iterator {
	return Iterator{
		table: m.table,
		list: m.ordering,
		index: -1,
		node: nil,
	}
}

func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.list.len {
		iterator.index++
	}
	if !iterator.list.withinRange(iterator.index) {
		iterator.node = nil
		return false
	}
	if iterator.index != 0 {
		iterator.node = iterator.node.next
	} else {
		iterator.node = iterator.list.head
	}
	return true
}

func (iterator *Iterator) Key() interface{} {
	return iterator.node.data
}

func (iterator *Iterator) Value() interface{} {
	key := iterator.node.data
	return iterator.table[key]
}