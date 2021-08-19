package types

type UnaryPredicate func(Data) bool

type BinaryPredicate func(Data, Data) bool
