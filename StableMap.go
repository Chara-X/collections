package collections

import (
	"container/list"
)

type StableMap[K comparable, V any] struct {
	Values   map[K]V
	elements map[K]*list.Element
	Keys     *list.List
}

func NewStableMap[K comparable, V any]() StableMap[K, V] {
	return StableMap[K, V]{Values: map[K]V{}, elements: make(map[K]*list.Element), Keys: list.New()}
}
func (m StableMap[K, V]) Store(key K, value V) {
	if node, ok := m.elements[key]; !ok {
		m.elements[key] = m.Keys.PushBack(key)
	} else {
		node.Value = key
	}
	m.Values[key] = value
}
func (m StableMap[K, V]) Load(key K) (V, bool) {
	var value, ok = m.Values[key]
	return value, ok
}
func (m StableMap[K, V]) Delete(key K) {
	if node, ok := m.elements[key]; ok {
		delete(m.Values, key)
		delete(m.elements, key)
		m.Keys.Remove(node)
	}
}
