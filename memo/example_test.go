package memo_test

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rbee3u/golib/memo"
)

func ExampleSimple() {
	m := memo.New()
	m.Set("x", 1)
	fmt.Println(m.Get("x"))
	m.Del("x")
	fmt.Println(m.Get("x"))

	// Output:
	// 1 <nil>
	// <nil> memo: not found
}

func ExampleLoader() {
	fmt.Println("Scene: default loader only")
	m := memo.New(memo.WithLoader(length))
	fmt.Println(m.Get("x"))
	fmt.Println(m.Get(233))

	fmt.Println("Scene: get loader only")
	m = memo.New()
	fmt.Println(m.Get("x", memo.GetWithLoader(length)))
	fmt.Println(m.Get(233, memo.GetWithLoader(length)))

	fmt.Println("Scene: get loader overwrites default loader")
	m = memo.New(memo.WithLoader(length))
	fmt.Println(m.Get("x", memo.GetWithLoader(doubleLength)))

	// Output:
	// Scene: default loader only
	// 1 <nil>
	// <nil> reflect: call of reflect.Value.Len on int Value
	// Scene: get loader only
	// 1 <nil>
	// <nil> reflect: call of reflect.Value.Len on int Value
	// Scene: get loader overwrites default loader
	// 2 <nil>
}

func ExampleExpiration() {
	const (
		expiration       = 300 * time.Millisecond
		doubleExpiration = 600 * time.Millisecond
	)

	fmt.Println("Scene: default expiration only")
	m := memo.New(memo.WithExpiration(expiration))
	m.Set("x", 1)
	fmt.Println(m.Get("x"))
	time.Sleep(expiration)
	fmt.Println(m.Get("x"))

	fmt.Println("Scene: set expiration only")
	m = memo.New()
	m.Set("x", 1, memo.SetWithExpiration(expiration))
	fmt.Println(m.Get("x"))
	time.Sleep(expiration)
	fmt.Println(m.Get("x"))

	fmt.Println("Scene: set expiration overwrites default expiration")
	m = memo.New(memo.WithExpiration(expiration))
	m.Set("x", 1, memo.SetWithExpiration(doubleExpiration))
	fmt.Println(m.Get("x"))
	time.Sleep(expiration)
	fmt.Println(m.Get("x"))
	time.Sleep(expiration)
	fmt.Println(m.Get("x"))

	// Output:
	// Scene: default expiration only
	// 1 <nil>
	// <nil> memo: not found
	// Scene: set expiration only
	// 1 <nil>
	// <nil> memo: not found
	// Scene: set expiration overwrites default expiration
	// 1 <nil>
	// 1 <nil>
	// <nil> memo: not found
}

func ExampleLoaderAndExpiration() {
	const expiration = 30 * time.Millisecond

	m := memo.New(memo.WithExpiration(expiration))
	fmt.Println(m.Get("x", memo.GetWithLoader(length)))
	time.Sleep(expiration)
	fmt.Println(m.Get("x"))

	// Output:
	// 1 <nil>
	// <nil> memo: not found
}

func ExampleConcurrency() {
	var counter int32

	wrappedLoader := func(k memo.Key) (memo.Value, error) {
		atomic.AddInt32(&counter, 1)
		return slowLength(k)
	}

	m := memo.New(memo.WithLoader(wrappedLoader))

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(m.Get("x"))
		}()
	}
	wg.Wait()

	fmt.Println("counter:", atomic.LoadInt32(&counter))

	// Output:
	// 1 <nil>
	// 1 <nil>
	// 1 <nil>
	// counter: 1
}

func length(k memo.Key) (memo.Value, error) {
	switch v := reflect.ValueOf(k); v.Kind() {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Map, reflect.Chan:
		return v.Len(), nil
	default:
		return nil, &reflect.ValueError{Method: "reflect.Value.Len", Kind: v.Kind()}
	}
}

func doubleLength(k memo.Key) (memo.Value, error) {
	switch v := reflect.ValueOf(k); v.Kind() {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Map, reflect.Chan:
		return 2 * v.Len(), nil
	default:
		return nil, &reflect.ValueError{Method: "reflect.Value.Len", Kind: v.Kind()}
	}
}

func slowLength(k memo.Key) (memo.Value, error) {
	time.Sleep(100 * time.Millisecond)

	return length(k)
}
