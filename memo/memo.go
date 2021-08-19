package memo

import (
	"container/heap"
	"sync"
)

// Memo is an in-memory k-v storage, which supports concurrently
// get/set/delete k-v pairs. The most special place is that it
// can load value if not found, and can set an expiration time.
type Memo struct {
	mu sync.Mutex
	o  options
	c  *cache
}

// New creates a memo with options.
func New(opts ...Option) *Memo {
	return &Memo{o: newOptions(opts...), c: newCache()}
}

// Get returns the associated value of the key.
// If the value is not found(or expired) but a loader is provided,
// the loader will be invoked to get a new value.
// If a new value is loaded and an expiration option is provided,
// the expiration option will act on the new value.
func (m *Memo) Get(k Key, opts ...GetOption) (Value, error) {
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

		return nil, ErrNotFound
	}

	e = newEntry()
	m.c.dictSet(k, e)
	m.c.heapPush(node{key: k, expireAt: expireAt})

	e.mu.Lock()
	m.mu.Unlock()
	defer e.mu.Unlock()
	e.value, e.err = o.loader(k)

	return e.value, e.err
}

// Set inserts a key-value pair into the memo, if the key
// already exists, update the associated value directly.
// If an expiration is provided, it will act on the pair.
func (m *Memo) Set(k Key, v Value, opts ...SetOption) {
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
		e = newEntry()
		e.value = v
		m.c.dictSet(k, e)
		m.c.heapPush(node{key: k, expireAt: expireAt})
		m.mu.Unlock()

		return
	}

	m.c.heapFix(e.position, node{key: k, expireAt: expireAt})

	m.mu.Unlock()
	e.mu.Lock()
	defer e.mu.Unlock()
	e.value, e.err = v, nil
}

// Del removes the key-value pair from the memo.
func (m *Memo) Del(k Key) {
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

func (m *Memo) cleanup(now int64) {
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
type cache struct {
	// A dict supports lookup value by key fastly.
	dict map[Key]*entry
	// A heap to hold all expiration time of keys.
	heap []node
	// The size of the heap.
	heapSize int
}

func newCache() *cache {
	return &cache{dict: make(map[Key]*entry)}
}

const zeroPosition = -1

type entry struct {
	mu       sync.Mutex
	position int
	value    Value
	err      error
}

func newEntry() *entry {
	return &entry{position: zeroPosition}
}

const zeroExpireAt = 0

type node struct {
	key      Key
	expireAt int64
}

func newNode() node {
	return node{}
}

func (c *cache) dictGet(k Key) *entry {
	return c.dict[k]
}

func (c *cache) dictSet(k Key, e *entry) {
	c.dict[k] = e
}

func (c *cache) dictDel(k Key) {
	delete(c.dict, k)
}

func (c *cache) heapEmpty() bool {
	return c.heapSize == 0
}

func (c *cache) heapTop() node {
	return c.heap[0]
}

func (c *cache) heapPop() {
	heap.Pop(c)
}

func (c *cache) heapPush(n node) {
	c.heapFix(zeroPosition, n)
}

func (c *cache) heapRemove(i int) {
	c.heapFix(i, node{})
}

func (c *cache) heapFix(i int, n node) {
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

func (c *cache) Len() int {
	return c.heapSize
}

func (c *cache) Less(i, j int) bool {
	return c.heap[i].expireAt < c.heap[j].expireAt
}

func (c *cache) Swap(i, j int) {
	if i != j {
		c.heap[i], c.heap[j] = c.heap[j], c.heap[i]
		c.dict[c.heap[i].key].position = i
		c.dict[c.heap[j].key].position = j
	}
}

func (c *cache) Push(n interface{}) {
	if c.heapSize == len(c.heap) {
		c.heap = append(c.heap, newNode())
	}

	c.heap[c.heapSize] = n.(node)
	c.dict[c.heap[c.heapSize].key].position = c.heapSize
	c.heapSize++
}

func (c *cache) Pop() interface{} {
	c.heapSize--
	c.dict[c.heap[c.heapSize].key].position = zeroPosition

	return c.heap[c.heapSize]
}
