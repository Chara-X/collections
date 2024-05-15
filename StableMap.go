package collections

import (
	"container/list"
)

type StableMap[K comparable, V any] struct {
	m      map[K]*list.Element
	Stable *list.List
}

func NewStableMap[K comparable, V any]() StableMap[K, V] {
	return StableMap[K, V]{map[K]*list.Element{}, list.New()}
}
func (m StableMap[K, V]) Store(key K, value V) {
	if _, ok := m.m[key]; ok {
		m.Delete(key)
	}
	m.m[key] = m.Stable.PushBack(KVPair[K, V]{key, value})
}
func (m StableMap[K, V]) Load(key K) (V, bool) {
	var elem, ok = m.m[key]
	return elem.Value.(KVPair[K, V]).Value, ok
}
func (m StableMap[K, V]) Delete(key K) {
	if node, ok := m.m[key]; ok {
		m.Stable.Remove(node)
		delete(m.m, key)
	}
}
