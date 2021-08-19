package runner_test

import (
	"testing"
	"time"

	"github.com/rbee3u/golib/runner"
	"github.com/thejerf/suture/v4"
)

func TestWithEventHook(t *testing.T) {
	var spec suture.Spec
	if spec.EventHook != nil {
		t.Errorf("expect spec.EventHook to be nil")
	}
	runner.WithEventHook(func(suture.Event) {})(&spec)
	if spec.EventHook == nil {
		t.Errorf("expect spec.EventHook not to be nil")
	}
	runner.WithEventHook(nil)(&spec)
	if spec.EventHook != nil {
		t.Errorf("expect spec.EventHook to be nil")
	}
}

func TestWithFailureDecay(t *testing.T) {
	var spec suture.Spec
	if spec.FailureDecay != 0 {
		t.Errorf("expect spec.FailureDecay(%v) to be 0", spec.FailureDecay)
	}
	runner.WithFailureDecay(1)(&spec)
	if spec.FailureDecay != 1 {
		t.Errorf("expect spec.FailureDecay(%v) to be 1", spec.FailureDecay)
	}
	runner.WithFailureDecay(0)(&spec)
	if spec.FailureDecay != 0 {
		t.Errorf("expect spec.FailureDecay(%v) to be 0", spec.FailureDecay)
	}
}

func TestWithFailureThreshold(t *testing.T) {
	var spec suture.Spec
	if spec.FailureThreshold != 0 {
		t.Errorf("expect spec.FailureThreshold(%v) to be 0", spec.FailureThreshold)
	}
	runner.WithFailureThreshold(1)(&spec)
	if spec.FailureThreshold != 1 {
		t.Errorf("expect spec.FailureThreshold(%v) to be 1", spec.FailureThreshold)
	}
	runner.WithFailureThreshold(0)(&spec)
	if spec.FailureThreshold != 0 {
		t.Errorf("expect spec.FailureThreshold(%v) to be 0", spec.FailureThreshold)
	}
}

func TestWithFailureBackoff(t *testing.T) {
	var spec suture.Spec
	if spec.FailureBackoff != 0*time.Second {
		t.Errorf("expect spec.FailureBackoff(%v) to be 0*time.Second", spec.FailureBackoff)
	}
	runner.WithFailureBackoff(1 * time.Second)(&spec)
	if spec.FailureBackoff != 1*time.Second {
		t.Errorf("expect spec.FailureBackoff(%v) to be 1*time.Second", spec.FailureBackoff)
	}
	runner.WithFailureBackoff(0 * time.Second)(&spec)
	if spec.FailureBackoff != 0*time.Second {
		t.Errorf("expect spec.FailureBackoff(%v) to be 0*time.Second", spec.FailureBackoff)
	}
}

func TestWithBackoffJitter(t *testing.T) {
	var spec suture.Spec
	if spec.BackoffJitter != nil {
		t.Errorf("expect spec.BackoffJitter to be nil")
	}
	runner.WithBackoffJitter(&suture.DefaultJitter{})(&spec)
	if spec.BackoffJitter == nil {
		t.Errorf("expect spec.BackoffJitter not to be nil")
	}
	runner.WithBackoffJitter(nil)(&spec)
	if spec.BackoffJitter != nil {
		t.Errorf("expect spec.BackoffJitter to be nil")
	}
}

func TestWithTimeout(t *testing.T) {
	var spec suture.Spec
	if spec.Timeout != 0*time.Second {
		t.Errorf("expect spec.Timeout(%v) to be 0*time.Second", spec.Timeout)
	}
	runner.WithTimeout(1 * time.Second)(&spec)
	if spec.Timeout != 1*time.Second {
		t.Errorf("expect spec.Timeout(%v) to be 1*time.Second", spec.Timeout)
	}
	runner.WithTimeout(0 * time.Second)(&spec)
	if spec.Timeout != 0*time.Second {
		t.Errorf("expect spec.Timeout(%v) to be 0*time.Second", spec.Timeout)
	}
}

func TestWithPassThroughPanics(t *testing.T) {
	var spec suture.Spec
	if spec.PassThroughPanics {
		t.Errorf("expect spec.PassThroughPanics(%v) to be false", spec.PassThroughPanics)
	}
	runner.WithPassThroughPanics(true)(&spec)
	if !spec.PassThroughPanics {
		t.Errorf("expect spec.PassThroughPanics(%v) to be true", spec.PassThroughPanics)
	}
	runner.WithPassThroughPanics(false)(&spec)
	if spec.PassThroughPanics {
		t.Errorf("expect spec.PassThroughPanics(%v) to be false", spec.PassThroughPanics)
	}
}

func TestWithDontPropagateTermination(t *testing.T) {
	var spec suture.Spec
	if spec.DontPropagateTermination {
		t.Errorf("expect spec.DontPropagateTermination(%v) to be false", spec.DontPropagateTermination)
	}
	runner.WithDontPropagateTermination(true)(&spec)
	if !spec.DontPropagateTermination {
		t.Errorf("expect spec.DontPropagateTermination(%v) to be true", spec.DontPropagateTermination)
	}
	runner.WithDontPropagateTermination(false)(&spec)
	if spec.DontPropagateTermination {
		t.Errorf("expect spec.DontPropagateTermination(%v) to be false", spec.DontPropagateTermination)
	}
}
