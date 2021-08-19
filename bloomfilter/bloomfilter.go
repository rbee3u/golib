package bloomfilter

import (
	"errors"
	"fmt"
	"math"
	"sync"

	"github.com/spaolacci/murmur3"
)

// BloomFilter implements [Bloom filter](https://en.wikipedia.org/wiki/Bloom_filter),
// a space-efficient probabilistic data structure used to test whether an item is a
// member of a set. False positive matches are possible, but false negatives are not.
// In other words, a query returns either "possibly in set" or "definitely not in set".
// Elements can be added to the set, but not removed (though this can be addressed
// with the counting Bloom filter variant); the more elements that are added to the
// set, the larger the probability of false positives.
type BloomFilter struct {
	// A mutex to let concurrency.
	sync.RWMutex
	// The capacity of bitset.
	m uint64
	// The storage of bitset.
	b bitset
	// Number of hash functions.
	k uint64
	// Hasher to generate hashes.
	h Hasher
}

// A Hasher transform a byte slice into two uint64(128 bits).
type Hasher func([]byte) (uint64, uint64)

// Option represents the option when creating a Bloom filter.
type Option func(*BloomFilter)

// WithHasher creates an option of hasher.
func WithHasher(h Hasher) Option {
	return func(bf *BloomFilter) {
		bf.h = h
	}
}

// ErrInvalidArgument represents an invalid argument error.
var ErrInvalidArgument = errors.New("invalid argument")

// NewWithEstimate creates a Bloom filter for about `n` items with `p` false
// positive possibility.
func NewWithEstimate(n uint64, p float64, opts ...Option) (*BloomFilter, error) {
	if n == 0 {
		return nil, fmt.Errorf("%w: n(%v)", ErrInvalidArgument, n)
	}

	if p <= 0 || p >= 1 {
		return nil, fmt.Errorf("%w: p(%v)", ErrInvalidArgument, p)
	}

	m, k := EstimateParameters(n, p)

	return New(m, k, opts...)
}

// EstimateParameters estimates the capacity of bitset `m` and the number of
// hash functions `k` for about `n` items with `p` false positive possibility.
func EstimateParameters(n uint64, p float64) (uint64, uint64) {
	m := math.Ceil(math.Log2E * math.Log2(1/p) * float64(n))
	k := math.Ceil(math.Ln2 * m / float64(n))

	return uint64(m), uint64(k)
}

// New creates a Bloom filter with `m` bits storage and `k` hash functions.
func New(m uint64, k uint64, opts ...Option) (*BloomFilter, error) {
	if m == 0 {
		return nil, fmt.Errorf("%w: m(%v)", ErrInvalidArgument, m)
	}

	if k == 0 {
		return nil, fmt.Errorf("%w: k(%v)", ErrInvalidArgument, k)
	}

	bf := &BloomFilter{
		m: m,
		b: newBitset(m),
		k: k,
		h: murmur3.Sum128,
	}

	for _, opt := range opts {
		opt(bf)
	}

	if bf.h == nil {
		return nil, fmt.Errorf("%w: nil hasher", ErrInvalidArgument)
	}

	return bf, nil
}

// Add adds item to the Bloom filter.
func (bf *BloomFilter) Add(item []byte) {
	bf.Lock()
	defer bf.Unlock()

	bf.AddWithoutLock(item)
}

// AddWithoutLock is same with Add, but without lock.
func (bf *BloomFilter) AddWithoutLock(item []byte) {
	a, b := bf.h(item)
	a, b = a%bf.m, b%bf.m

	for i := uint64(0); i < bf.k; i++ {
		bf.b.mark((a + i*b) % bf.m)
	}
}

// Contains returns true if the item is in the Bloom filter, false otherwise.
// If true, the result might be a false positive.
// If false, the item is definitely not in the set.
func (bf *BloomFilter) Contains(item []byte) bool {
	bf.RLock()
	defer bf.RUnlock()

	return bf.ContainsWithoutLock(item)
}

// ContainsWithoutLock is same with Contains, but without lock.
func (bf *BloomFilter) ContainsWithoutLock(item []byte) bool {
	a, b := bf.h(item)
	a, b = a%bf.m, b%bf.m

	for i := uint64(0); i < bf.k; i++ {
		if bf.b.test((a + i*b) % bf.m) {
			return false
		}
	}

	return true
}

type bitset []uint8

func newBitset(m uint64) bitset {
	return make(bitset, (m+8-1)/8)
}

func (b bitset) mark(p uint64) {
	b[p/8] |= 1 << (p % 8)
}

func (b bitset) test(p uint64) bool {
	return b[p/8]&(1<<(p%8)) == 0
}
