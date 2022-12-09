package collections

type Collection[T interface{}] interface {
	CreateIterator() Iterator[T]
}
