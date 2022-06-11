package ibch

import (
	"sync"
)

// An IBCh represents a channel with infinite buffer, which consists of a
// sendCh and a recvCh, the sendCh for sending and the recvCh for receiving.
type IBCh[T any] struct {
	// The sendCh is a channel which used to send data.
	sendCh chan T
	// The recvCh is a channel which used to receive data.
	recvCh chan T
}

// New creates an IBCh.
func New[T any]() *IBCh[T] {
	ibCh := &IBCh[T]{
		sendCh: make(chan T),
		recvCh: make(chan T),
	}

	go ibCh.loop()

	return ibCh
}

// SendCh returns the sendCh.
func (ibCh *IBCh[T]) SendCh() chan<- T {
	return ibCh.sendCh
}

// RecvCh returns the recvCh.
func (ibCh *IBCh[T]) RecvCh() <-chan T {
	return ibCh.recvCh
}

// Close will close the whole ibCh by closing the sendCh.
func (ibCh *IBCh[T]) Close() {
	close(ibCh.sendCh)
}

func (ibCh *IBCh[T]) loop() {
	q := newQueue[T]()

	for {
		if q.empty() {
			data, unclosed := <-ibCh.sendCh
			if !unclosed {
				goto exit
			}

			q.pushBack(data)
		}

		select {
		case data, unclosed := <-ibCh.sendCh:
			if !unclosed {
				goto exit
			}

			q.pushBack(data)
		case ibCh.recvCh <- q.front():
			q.popFront()
		}
	}

exit:
	for ; !q.empty(); q.popFront() {
		ibCh.recvCh <- q.front()
	}

	close(ibCh.recvCh)
}

type queue[T any] struct {
	pool sync.Pool
	head *node[T]
	tail *node[T]
}

func newQueue[T any]() *queue[T] {
	var q queue[T]
	q.pool.New = func() any { return new(node[T]) }

	return &q
}

type node[T any] struct {
	next *node[T]
	data T
}

func (q *queue[T]) empty() bool {
	return q.head == nil
}

func (q *queue[T]) front() T {
	return q.head.data
}

func (q *queue[T]) popFront() {
	n := q.head

	q.head = n.next
	if q.head == nil {
		q.tail = nil
	}

	*n = node[T]{}
	q.pool.Put(n)
}

func (q *queue[T]) pushBack(data T) {
	n := q.pool.Get().(*node[T])
	n.data = data

	if q.head == nil {
		q.head = n
	} else {
		q.tail.next = n
	}

	q.tail = n
}
