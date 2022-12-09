package collections

type Iterator[T interface{}] interface {
	HasNext() bool
	GetNext() T
	GetAllElements() []T
}
