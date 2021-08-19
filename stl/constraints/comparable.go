package constraints

type EqualityComparable interface {
	Equal(other EqualityComparable) bool
}

type LessThanComparable interface {
	Less(other LessThanComparable) bool
}
