package ibch_test

import (
	"fmt"

	"github.com/rbee3u/golib/ibch"
)

func Example() {
	ibCh := ibch.New[int]()

	for i := 0; i < 5; i++ {
		ibCh.SendCh() <- i
		fmt.Printf("send: %v\n", i)
	}

	ibCh.Close()

	for i := range ibCh.RecvCh() {
		fmt.Printf("recv: %v\n", i)
	}

	// Output:
	// send: 0
	// send: 1
	// send: 2
	// send: 3
	// send: 4
	// recv: 0
	// recv: 1
	// recv: 2
	// recv: 3
	// recv: 4
}
