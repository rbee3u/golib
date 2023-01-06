package memo

import (
	"errors"
	"time"
	"unsafe"
)

var (
	// ErrNotFound is an error returned when the key is not found.
	ErrNotFound = errors.New("memo: not found")
	// ErrInvalidExpiration represents an invalid expiration error.
	ErrInvalidExpiration = errors.New("memo: invalid expiration")
)

// A Loader returns the value of the key.
type Loader[K comparable, V any] func(K) (V, error)

// options holds all extra configs needed when creating a new memo.
type options[K comparable, V any] struct {
	// The clock provides the current time in nanoseconds.
	clock Clock
	// Default loader used in memo.Get method.
	loader Loader[K, V]
	// Default expiration used in memo.Get and memo.Set method.
	expiration time.Duration
}

// Option specifies the option when creating a new memo.
type Option[K comparable, V any] func(*options[K, V])

func newOptions[K comparable, V any](opts ...Option[K, V]) options[K, V] {
	o := options[K, V]{clock: NewRealClock()}
	for _, opt := range opts {
		opt((*options[K, V])(noescape(unsafe.Pointer(&o))))
	}

	if o.expiration < 0 {
		panic(ErrInvalidExpiration)
	}

	return o
}

// WithClock provides a clock option when creating a new memo.
func WithClock[K comparable, V any](clock Clock) Option[K, V] {
	return func(o *options[K, V]) {
		o.clock = clock
	}
}

// WithLoader provides a loader option when creating a new memo.
func WithLoader[K comparable, V any](loader Loader[K, V]) Option[K, V] {
	return func(o *options[K, V]) {
		o.loader = loader
	}
}

// WithExpiration provides an expiration option when creating a new memo.
func WithExpiration[K comparable, V any](expiration time.Duration) Option[K, V] {
	return func(o *options[K, V]) {
		o.expiration = expiration
	}
}

// options holds all extra configs needed when getting a value from the memo.
type getOptions[K comparable, V any] struct {
	// Load a value by key when is not found.
	loader Loader[K, V]
	// Expiration for the value to be loaded.
	expiration time.Duration
}

// GetOption specifies the option when getting a value from the memo.
type GetOption[K comparable, V any] func(*getOptions[K, V])

func (base *options[K, V]) newGetOptions(opts ...GetOption[K, V]) getOptions[K, V] {
	o := getOptions[K, V]{loader: base.loader, expiration: base.expiration}
	for _, opt := range opts {
		opt((*getOptions[K, V])(noescape(unsafe.Pointer(&o))))
	}

	if o.expiration < 0 {
		panic(ErrInvalidExpiration)
	}

	return o
}

// GetWithLoader provides a loader option when getting a value from the memo.
func GetWithLoader[K comparable, V any](loader Loader[K, V]) GetOption[K, V] {
	return func(o *getOptions[K, V]) {
		o.loader = loader
	}
}

// GetWithExpiration provides an expiration option when getting a value from the memo.
func GetWithExpiration[K comparable, V any](expiration time.Duration) GetOption[K, V] {
	return func(o *getOptions[K, V]) {
		o.expiration = expiration
	}
}

// options holds all extra configs needed when setting a value to the memo.
type setOptions[K comparable, V any] struct {
	// Expiration for the value to be set.
	expiration time.Duration
}

// SetOption specifies the option when setting a value to the memo.
type SetOption[K comparable, V any] func(*setOptions[K, V])

func (base *options[K, V]) newSetOptions(opts ...SetOption[K, V]) setOptions[K, V] {
	o := setOptions[K, V]{expiration: base.expiration}
	for _, opt := range opts {
		opt((*setOptions[K, V])(noescape(unsafe.Pointer(&o))))
	}

	if o.expiration < 0 {
		panic(ErrInvalidExpiration)
	}

	return o
}

// SetWithExpiration provides an expiration option when setting a value to the memo.
func SetWithExpiration[K comparable, V any](expiration time.Duration) SetOption[K, V] {
	return func(o *setOptions[K, V]) {
		o.expiration = expiration
	}
}
