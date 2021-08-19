package bloomfilter_test

import (
	"fmt"

	"github.com/rbee3u/golib/bloomfilter"
)

func Example() {
	bf, err := bloomfilter.NewWithEstimate(1000000, 0.03)
	if err != nil {
		fmt.Printf("failed to new Bloom filter: %v\n", err)
		return
	}

	bf.Add([]byte("foo"))
	for _, item := range []string{"foo", "bar"} {
		if bf.Contains([]byte(item)) {
			fmt.Printf("%s is possibly in the set\n", item)
		} else {
			fmt.Printf("%s is definitely not in the set\n", item)
		}
	}

	// Output:
	// foo is possibly in the set
	// bar is definitely not in the set
}
