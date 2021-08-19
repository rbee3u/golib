// Package runner is a simple wrapper of [suture](https://github.com/thejerf/suture).
// The package suture has done a very good job for reimplementing
// [supervisor behaviour](https://erlang.org/doc/design_principles/sup_princ.html), which
// provides very powerful capabilities. While most of the time, we only want to use its
// basic capabilities. So this package does the following two things:
//
// - Wrap suture to make it easier to use
// - Trap system signals and graceful stop
//
// For details on how to use it, you can refer to the [example](example_test.go). By the
// way, if you want to use more advanced customization capabilities, I suggest you to use
// suture directly.
package runner

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/thejerf/suture/v4"
)

// Runner is a simple wrapper of [suture](https://github.com/thejerf/suture).
type Runner struct {
	// The name of runner, which is also the name of supervisor.
	name string
	// The spec of runner, which is also the spec of supervisor.
	spec suture.Spec
}

// New create a new runner based on name and opts.
func New(name string, opts ...Option) *Runner {
	r := &Runner{name: name}
	r.spec.EventHook = func(suture.Event) {}

	for _, opt := range opts {
		opt(&r.spec)
	}

	return r
}

// Run execute services by constructing a supervisor.
func (r *Runner) Run(ctx context.Context, services ...suture.Service) error {
	ctx, stop := signal.NotifyContext(ctx,
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	supervisor := suture.New(r.name, r.spec)
	for _, service := range services {
		supervisor.Add(service)
	}

	return <-supervisor.ServeBackground(ctx)
}
