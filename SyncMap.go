package collections

import (
	"sync"
	"sync/atomic"
)

type SyncMap[K comparable, V any] struct {
	m      map[K]V
	locker sync.RWMutex
	cache  atomic.Pointer[map[K]*atomic.Pointer[V]]
}

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	var m = &SyncMap[K, V]{m: make(map[K]V)}
	m.cache = atomic.Pointer[map[K]*atomic.Pointer[V]]{}
	m.cache.Store(&map[K]*atomic.Pointer[V]{})
	return m
}
func (m *SyncMap[K, V]) Store(key K, value V) {
	if v, ok := (*m.cache.Load())[key]; ok {
		v.Store(&value)
	} else {
		m.locker.Lock()
		defer m.locker.Unlock()
		m.m[key] = value
		if len(m.m) > 10 {
			m.sync()
		}
	}
}
func (m *SyncMap[K, V]) Load(key K) (V, bool) {
	if v, ok := (*m.cache.Load())[key]; ok {
		if v := v.Load(); v != nil {
			return *v, true
		}
	}
	m.locker.RLock()
	defer m.locker.RUnlock()
	var v, ok = m.m[key]
	return v, ok
}
func (m *SyncMap[K, V]) Delete(key K) {
	if v, ok := (*m.cache.Load())[key]; ok {
		v.Store(nil)
	} else {
		m.locker.Lock()
		defer m.locker.Unlock()
		delete(m.m, key)
	}
}
func (m *SyncMap[K, V]) sync() {
	var cache = map[K]*atomic.Pointer[V]{}
	for k, v := range *m.cache.Load() {
		cache[k] = &atomic.Pointer[V]{}
		if v := v.Load(); v != nil {
			cache[k].Store(v)
		}
	}
	for k, v := range m.m {
		cache[k] = &atomic.Pointer[V]{}
		cache[k].Store(&v)
	}
	m.cache.Store(&cache)
	m.m = make(map[K]V)
}
