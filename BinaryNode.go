package collections

type BinaryNode[K comparable, V any] struct {
	Key   K
	Value V
	Left  *BinaryNode[K, V]
	Right *BinaryNode[K, V]
}
