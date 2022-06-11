package constraints

type Cloneable[S any] interface {
	Clone() S
}

func IsCloneable[S Cloneable[S]]() bool { return true }

type EqualityComparable[S any] interface {
	Equal(other S) bool
}

func IsEqualityComparable[S EqualityComparable[S]]() bool { return true }

type LessThanComparable[S any] interface {
	Less(other S) bool
}

func IsLessThanComparable[S LessThanComparable[S]]() bool { return true }

type Incrementable[S any] interface {
	Next() S
}

func IsIncrementable[S Incrementable[S]]() bool { return true }

type Decrementable[S any] interface {
	Prev() S
}

func IsDecrementable[S Decrementable[S]]() bool { return true }

type Readable[T any] interface {
	Read() T
}

func IsReadable[S Readable[S]]() bool { return true }

type Writeable[T any] interface {
	Write(data T)
}

func IsWriteable[S Writeable[S]]() bool { return true }
