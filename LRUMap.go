package collections

import "container/list"

type LRUMap[K comparable, V any] struct {
	m   map[K]*list.Element
	LRU *list.List
}

func NewLRUMap[K comparable, V any]() *LRUMap[K, V] {
	return &LRUMap[K, V]{map[K]*list.Element{}, list.New()}
}
func (m *LRUMap[K, V]) Store(key K, value V) {
	if _, ok := m.m[key]; ok {
		m.Delete(key)
	}
	m.m[key] = m.LRU.PushBack(KVPair[K, V]{key, value})
}
func (m *LRUMap[K, V]) Load(key K) (V, bool) {
	if elem, ok := m.m[key]; ok {
		m.LRU.MoveToBack(elem)
		return elem.Value.(KVPair[K, V]).Value, ok
	}
	return *new(V), false
}
func (m *LRUMap[K, V]) Delete(key K) {
	if e := m.m[key]; e != nil {
		m.LRU.Remove(e)
		delete(m.m, key)
	}
}
