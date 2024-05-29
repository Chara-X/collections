package collections

import "cmp"

type BinarySearchTree[K cmp.Ordered, V any] struct{ root *BinaryNode[K, V] }

func (t *BinarySearchTree[K, V]) Store(key K, value V) { t.root = t.store(t.root, key, value) }
func (t *BinarySearchTree[K, V]) store(node *BinaryNode[K, V], key K, value V) *BinaryNode[K, V] {
	if node == nil {
		return &BinaryNode[K, V]{Key: key, Value: value}
	}
	if key < node.Key {
		node.Left = t.store(node.Left, key, value)
	} else if key > node.Key {
		node.Right = t.store(node.Right, key, value)
	} else {
		node.Value = value
	}
	return node
}
func (t *BinarySearchTree[K, V]) Query(predicate func(node *BinaryNode[K, V]) bool, toLeft func(node *BinaryNode[K, V]) bool, toRight func(node *BinaryNode[K, V]) bool) []V {
	return t.query(t.root, predicate, toLeft, toRight)
}
func (t *BinarySearchTree[K, V]) query(node *BinaryNode[K, V], predicate func(node *BinaryNode[K, V]) bool, toLeft func(node *BinaryNode[K, V]) bool, toRight func(node *BinaryNode[K, V]) bool) []V {
	if node == nil {
		return []V{}
	}
	var res []V
	if toLeft(node) {
		res = append(res, t.query(node.Left, predicate, toLeft, toRight)...)
	}
	if predicate(node) {
		res = append(res, node.Value)
	}
	if toRight(node) {
		res = append(res, t.query(node.Right, predicate, toLeft, toRight)...)
	}
	return res
}
func (t *BinarySearchTree[K, V]) ForEach(fn func(node *BinaryNode[K, V])) { t.forEach(t.root, fn) }
func (t *BinarySearchTree[K, V]) forEach(node *BinaryNode[K, V], fn func(node *BinaryNode[K, V])) {
	if node == nil {
		return
	}
	t.forEach(node.Left, fn)
	fn(node)
	t.forEach(node.Right, fn)
}
