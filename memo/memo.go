package memo

import (
	"container/heap"
	"sync"
)

// Memo is an in-memory k-v storage, which supports concurrently
// get/set/delete k-v pairs. The most special place is that it
// can load value if not found, and can set an expiration time.
type Memo[K comparable, V any] struct {
	mu sync.Mutex
	o  options[K, V]
	c  *cache[K, V]
}

// New creates a memo with options.
func New[K comparable, V any](opts ...Option[K, V]) *Memo[K, V] {
	return &Memo[K, V]{
		o: newOptions[K, V](opts...),
		c: newCache[K, V](),
	}
}

// Get returns the associated value of the key.
// If the value is not found(or expired) but a loader is provided,
// the loader will be invoked to get a new value.
// If a new value is loaded and an expiration option is provided,
// the expiration option will act on the new value.
func (m *Memo[K, V]) Get(k K, opts ...GetOption[K, V]) (V, error) {
	o := m.o.newGetOptions(opts...)
	now := m.o.clock.Now()

	var expireAt int64
	if o.expiration != 0 {
		expireAt = now + int64(o.expiration)
	}

	m.mu.Lock()
	m.cleanup(now)

	e := m.c.dictGet(k)
	if e != nil {
		m.mu.Unlock()
		e.mu.Lock()
		defer e.mu.Unlock()

		return e.value, e.err
	}

	if o.loader == nil {
		m.mu.Unlock()

		var zero V

		return zero, ErrNotFound
	}

	e = newEntry[V]()
	m.c.dictSet(k, e)
	m.c.heapPush(node[K]{key: k, expireAt: expireAt})

	e.mu.Lock()
	m.mu.Unlock()
	defer e.mu.Unlock()
	e.value, e.err = o.loader(k)

	return e.value, e.err
}

// Set inserts a key-value pair into the memo, if the key
// already exists, update the associated value directly.
// If an expiration is provided, it will act on the pair.
func (m *Memo[K, V]) Set(k K, v V, opts ...SetOption[K, V]) {
	o := m.o.newSetOptions(opts...)
	now := m.o.clock.Now()

	var expireAt int64
	if o.expiration != 0 {
		expireAt = now + int64(o.expiration)
	}

	m.mu.Lock()
	m.cleanup(now)

	e := m.c.dictGet(k)
	if e == nil {
		e = newEntry[V]()
		e.value = v
		m.c.dictSet(k, e)
		m.c.heapPush(node[K]{key: k, expireAt: expireAt})
		m.mu.Unlock()

		return
	}

	m.c.heapFix(e.position, node[K]{key: k, expireAt: expireAt})

	m.mu.Unlock()
	e.mu.Lock()
	defer e.mu.Unlock()
	e.value, e.err = v, nil
}

// Del removes the key-value pair from the memo.
func (m *Memo[K, V]) Del(k K) {
	now := m.o.clock.Now()
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cleanup(now)

	e := m.c.dictGet(k)
	if e == nil {
		return
	}

	m.c.heapRemove(e.position)
	m.c.dictDel(k)
}

func (m *Memo[K, V]) cleanup(now int64) {
	for !m.c.heapEmpty() {
		top := m.c.heapTop()
		if top.expireAt > now {
			break
		}

		m.c.heapPop()
		m.c.dictDel(top.key)
	}
}

// cache is the actual storage layer of memo.
type cache[K comparable, V any] struct {
	// A dict supports lookup value by key quickly.
	dict map[K]*entry[V]
	// A heap to hold all expiration time of keys.
	heap []node[K]
	// The size of the heap.
	heapSize int
}

func newCache[K comparable, V any]() *cache[K, V] {
	return &cache[K, V]{dict: make(map[K]*entry[V])}
}

const zeroPosition = -1

type entry[V any] struct {
	mu       sync.Mutex
	position int
	value    V
	err      error
}

func newEntry[V any]() *entry[V] {
	return &entry[V]{position: zeroPosition}
}

const zeroExpireAt = 0

type node[K comparable] struct {
	key      K
	expireAt int64
}

func newNode[K comparable]() node[K] {
	return node[K]{}
}

func (c *cache[K, V]) dictGet(k K) *entry[V] {
	return c.dict[k]
}

func (c *cache[K, V]) dictSet(k K, e *entry[V]) {
	c.dict[k] = e
}

func (c *cache[K, V]) dictDel(k K) {
	delete(c.dict, k)
}

func (c *cache[K, V]) heapEmpty() bool {
	return c.heapSize == 0
}

func (c *cache[K, V]) heapTop() node[K] {
	return c.heap[0]
}

func (c *cache[K, V]) heapPop() {
	heap.Pop(c)
}

func (c *cache[K, V]) heapPush(n node[K]) {
	c.heapFix(zeroPosition, n)
}

func (c *cache[K, V]) heapRemove(i int) {
	c.heapFix(i, node[K]{})
}

func (c *cache[K, V]) heapFix(i int, n node[K]) {
	switch {
	case i == zeroPosition && n.expireAt != zeroExpireAt:
		heap.Push(c, n)
	case i != zeroPosition && n.expireAt == zeroExpireAt:
		heap.Remove(c, i)
	case i != zeroPosition && c.heap[i].expireAt != n.expireAt:
		c.heap[i].expireAt = n.expireAt
		heap.Fix(c, i)
	}
}

func (c *cache[K, V]) Len() int {
	return c.heapSize
}

func (c *cache[K, V]) Less(i, j int) bool {
	return c.heap[i].expireAt < c.heap[j].expireAt
}

func (c *cache[K, V]) Swap(i, j int) {
	if i != j {
		c.heap[i], c.heap[j] = c.heap[j], c.heap[i]
		c.dict[c.heap[i].key].position = i
		c.dict[c.heap[j].key].position = j
	}
}

func (c *cache[K, V]) Push(n interface{}) {
	if c.heapSize == len(c.heap) {
		c.heap = append(c.heap, newNode[K]())
	}

	c.heap[c.heapSize] = n.(node[K])
	c.dict[c.heap[c.heapSize].key].position = c.heapSize
	c.heapSize++
}

func (c *cache[K, V]) Pop() interface{} {
	c.heapSize--
	c.dict[c.heap[c.heapSize].key].position = zeroPosition

	return c.heap[c.heapSize]
}
