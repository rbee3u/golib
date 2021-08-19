package runner_test

import (
	"context"
	"errors"
	"fmt"
	"go/build"
	"syscall"
	"time"

	"github.com/rbee3u/golib/runner"
	"github.com/thejerf/suture/v4"
)

func ExampleSystemSignal() {
	// The two tasks are all terminated by themselves, and
	// the supervisor is terminated by system signal(kill).
	time.AfterFunc(millis(200), func() {
		_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	})

	err := runner.New("main").Run(context.Background(),
		&task{name: "task", wait: millis(500), err: suture.ErrDoNotRestart},
		&task{name: "task", wait: millis(800), err: suture.ErrDoNotRestart},
	)

	fmt.Printf("err: %v", err)

	// Output:
	// <task> stop
	// <task> stop
	// err: context canceled
}

func ExampleTerminateAll() {
	// The task-1 is first terminated by itself, then it will
	// terminate the supervisor immediately, including task-2.
	err := runner.New("main").Run(context.Background(),
		&task{name: "task-1", wait: millis(200), err: suture.ErrTerminateSupervisorTree},
		&task{name: "task-2", wait: millis(500), err: suture.ErrDoNotRestart},
	)

	fmt.Printf("err: %v", err)

	// Output:
	// <task-1> exec
	// <task-2> stop
	// err: tree should be terminated
}

func ExampleTerminateOne() {
	// The task-1 is first terminated by itself, the task-2
	// is terminated later(the supervisor is also terminated).
	err := runner.New("main").Run(context.Background(),
		&task{name: "task-1", wait: millis(200), err: suture.ErrDoNotRestart},
		&task{name: "task-2", wait: millis(500), err: suture.ErrTerminateSupervisorTree},
	)

	fmt.Printf("err: %v", err)

	// Output:
	// <task-1> exec
	// <task-2> exec
	// err: tree should be terminated
}

func ExampleRestart() {
	// The task-2 will be restarted twice, and then, the
	// whole supervisor tree will be terminated by task-1.
	err := runner.New("main").Run(context.Background(),
		&task{name: "task-1", wait: millis(500), err: suture.ErrTerminateSupervisorTree},
		&task{name: "task-2", wait: millis(200), err: errors.New("should be restarted")},
	)

	fmt.Printf("err: %v", err)

	// Output:
	// <task-2> exec
	// <task-2> exec
	// <task-1> exec
	// <task-2> stop
	// err: tree should be terminated
}

func ExampleBackoff() {
	// The task-2 will be restarted twice in 400ms, and the
	// whole supervisor tree will wait for 200ms. Then the
	// task-2 will be restarted once again, at the same time,
	// the whole supervisor will be terminated by task-1.
	err := runner.New("main",
		runner.WithFailureThreshold(1),
		runner.WithFailureBackoff(millis(200)),
		runner.WithBackoffJitter(suture.NoJitter{}),
	).Run(context.Background(),
		&task{name: "task-1", wait: millis(900), err: suture.ErrTerminateSupervisorTree},
		&task{name: "task-2", wait: millis(200), err: errors.New("should be restarted")},
	)

	fmt.Printf("err: %v", err)

	// Output:
	// <task-2> exec
	// <task-2> exec
	// <task-2> exec
	// <task-1> exec
	// <task-2> stop
	// err: tree should be terminated
}

type task struct {
	name string
	wait time.Duration
	err  error
}

func (t *task) String() string {
	return t.name
}

func (t *task) Serve(ctx context.Context) error {
	select {
	case <-ctx.Done():
		fmt.Printf("<%s> stop\n", t)
		return suture.ErrDoNotRestart
	case <-time.After(t.wait):
		fmt.Printf("<%s> exec\n", t)
		return t.err
	}
}

func millis(x time.Duration) time.Duration {
	x *= time.Millisecond

	if build.Default.GOOS == "darwin" && build.Default.GOPATH == "/Users/runner/go" {
		x *= 10
	}

	return x
}
