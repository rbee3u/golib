package constraints

type Iterator[S any] interface {
	Cloneable[S]
	Incrementable[S]
}

func IsIterator[S Iterator[S]]() bool { return true }

type InputIterator[S any, T any] interface {
	Iterator[S]
	EqualityComparable[S]
	Readable[T]
}

func IsInputIterator[S InputIterator[S, T], T any]() bool { return true }

type ForwardIterator[S any, T any] interface {
	InputIterator[S, T]
	// Multiple passes allowed.
}

func IsForwardIterator[S ForwardIterator[S, T], T any]() bool { return true }

type BidirectionalIterator[S any, T any] interface {
	ForwardIterator[S, T]
	Decrementable[S]
}

func IsBidirectionalIterator[S BidirectionalIterator[S, T], T any]() bool { return true }

type RandomAccessIterator[S any, T any] interface {
	BidirectionalIterator[S, T]
	LessThanComparable[S]
	Advance(offset int) S
	At(offset int) T
	Distance(other S) int
}

func IsRandomAccessIterator[S RandomAccessIterator[S, T], T any]() bool { return true }

type MutableIterator[T any] interface {
	Writeable[T]
}

func IsMutableIterator[S MutableIterator[S]]() bool { return true }

type MutableInputIterator[S any, T any] interface {
	MutableIterator[T]
	InputIterator[S, T]
}

func IsMutableInputIterator[S MutableInputIterator[S, T], T any]() bool { return true }

type MutableForwardIterator[S any, T any] interface {
	MutableIterator[T]
	ForwardIterator[S, T]
}

func IsMutableForwardIterator[S MutableForwardIterator[S, T], T any]() bool { return true }

type MutableBidirectionalIterator[S any, T any] interface {
	MutableIterator[T]
	BidirectionalIterator[S, T]
}

func IsMutableBidirectionalIterator[S MutableBidirectionalIterator[S, T], T any]() bool { return true }

type MutableRandomAccessIterator[S any, T any] interface {
	MutableIterator[T]
	RandomAccessIterator[S, T]
}

func IsMutableRandomAccessIterator[S MutableRandomAccessIterator[S, T], T any]() bool { return true }
