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

type (
	// Key specifies the key type of memo.
	Key = interface{}
	// Value specifies the value type of memo.
	Value = interface{}
	// A Loader returns the value of the key.
	Loader func(Key) (Value, error)
)

// options holds all extra configs needed when creating a new memo.
type options struct {
	// The clock provides the current time in nanoseconds.
	clock Clock
	// Default loader used in memo.Get method.
	loader Loader
	// Default expiration used in memo.Get and memo.Set method.
	expiration time.Duration
}

// Option specifies the option when creating a new memo.
type Option func(*options)

func newOptions(opts ...Option) options {
	o := options{clock: NewRealClock()}
	for _, opt := range opts {
		opt((*options)(noescape(unsafe.Pointer(&o))))
	}

	if o.expiration < 0 {
		panic(ErrInvalidExpiration)
	}

	return o
}

// WithClock provides a clock option when creating a new memo.
func WithClock(clock Clock) Option {
	return func(o *options) {
		o.clock = clock
	}
}

// WithLoader provides a loader option when creating a new memo.
func WithLoader(loader Loader) Option {
	return func(o *options) {
		o.loader = loader
	}
}

// WithExpiration provides an expiration option when creating a new memo.
func WithExpiration(expiration time.Duration) Option {
	return func(o *options) {
		o.expiration = expiration
	}
}

// options holds all extra configs needed when getting a value from the memo.
type getOptions struct {
	// Load a value by key when is not found.
	loader Loader
	// Expiration for the value to be loaded.
	expiration time.Duration
}

// GetOption specifies the option when getting a value from the memo.
type GetOption func(*getOptions)

func (base *options) newGetOptions(opts ...GetOption) getOptions {
	o := getOptions{loader: base.loader, expiration: base.expiration}
	for _, opt := range opts {
		opt((*getOptions)(noescape(unsafe.Pointer(&o))))
	}

	if o.expiration < 0 {
		panic(ErrInvalidExpiration)
	}

	return o
}

// GetWithLoader provides a loader option when getting a value from the memo.
func GetWithLoader(loader Loader) GetOption {
	return func(o *getOptions) {
		o.loader = loader
	}
}

// GetWithExpiration provides an expiration option when getting a value from the memo.
func GetWithExpiration(expiration time.Duration) GetOption {
	return func(o *getOptions) {
		o.expiration = expiration
	}
}

// options holds all extra configs needed when setting a value to the memo.
type setOptions struct {
	// Expiration for the value to be set.
	expiration time.Duration
}

// SetOption specifies the option when setting a value to the memo.
type SetOption func(*setOptions)

func (base *options) newSetOptions(opts ...SetOption) setOptions {
	o := setOptions{expiration: base.expiration}
	for _, opt := range opts {
		opt((*setOptions)(noescape(unsafe.Pointer(&o))))
	}

	if o.expiration < 0 {
		panic(ErrInvalidExpiration)
	}

	return o
}

// SetWithExpiration provides an expiration option when setting a value to the memo.
func SetWithExpiration(expiration time.Duration) SetOption {
	return func(o *setOptions) {
		o.expiration = expiration
	}
}
