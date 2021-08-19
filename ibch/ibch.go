package ibch

import (
	"sync"
)

// An IBCh represents a channel with infinite buffer, which consists of a
// sendCh and a recvCh, the sendCh for sending and the recvCh for receiving.
type IBCh struct {
	// The sendCh is a channel which used to send data.
	sendCh chan interface{}
	// The recvCh is a channel which used to receive data.
	recvCh chan interface{}
}

// New creates an IBCh.
func New() *IBCh {
	ibCh := &IBCh{
		sendCh: make(chan interface{}),
		recvCh: make(chan interface{}),
	}

	go ibCh.loop()

	return ibCh
}

// SendCh returns the sendCh.
func (ibCh *IBCh) SendCh() chan<- interface{} {
	return ibCh.sendCh
}

// RecvCh returns the recvCh.
func (ibCh *IBCh) RecvCh() <-chan interface{} {
	return ibCh.recvCh
}

// Close will close the whole ibCh by closing the sendCh.
func (ibCh *IBCh) Close() {
	close(ibCh.sendCh)
}

func (ibCh *IBCh) loop() {
	q := newQueue()

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

type queue struct {
	pool sync.Pool
	head *node
	tail *node
}

func newQueue() *queue {
	var q queue
	q.pool.New = func() interface{} {
		return new(node)
	}

	return &q
}

type node struct {
	next *node
	data interface{}
}

func (q *queue) empty() bool {
	return q.head == nil
}

func (q *queue) front() interface{} {
	return q.head.data
}

func (q *queue) popFront() {
	n := q.head

	q.head = n.next
	if q.head == nil {
		q.tail = nil
	}

	n.next = nil
	n.data = nil
	q.pool.Put(n)
}

func (q *queue) pushBack(data interface{}) {
	n := q.pool.Get().(*node)
	n.data = data

	if q.head == nil {
		q.head = n
	} else {
		q.tail.next = n
	}

	q.tail = n
}
