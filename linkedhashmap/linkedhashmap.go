package linkedhashmap

import (
	// "fmt"
)

type Map struct {
	table map[interface{}]interface{}
	ordering *DoublyLinkedList
}

func New() *Map {
	return &Map{
		table: make(map[interface{}]interface{}),
		ordering: initDoublyList(),
	}
}

func (m *Map) Put(key interface{}, value interface{}) {
	if _, contains := m.table[key]; !contains {
		m.ordering.Append(key)
	}
	m.table[key] = value
}

func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	value = m.table[key]
	found = value != nil
	return
}

func (m *Map) Remove(key interface{}) {
	if _, contains := m.table[key]; contains {
		delete(m.table, key)
		index := m.ordering.IndexOf(key)
		m.ordering.Remove(index)
	}
}

func (m *Map) Size() int {
	return m.ordering.Size()
}

func (m *Map) Keys() []interface{} {
	return m.ordering.Values()
}

func (m *Map) All(f func(key interface{}, value interface{}) bool) bool {
	iterator := m.Iterator()
	for iterator.Next() {
		if !f(iterator.Key(), iterator.Value()) {
			return false
		}
	}
	return true
}