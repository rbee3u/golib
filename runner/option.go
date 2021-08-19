package runner

import (
	"time"

	"github.com/thejerf/suture/v4"
)

type Option func(*suture.Spec)

func WithEventHook(eventHook suture.EventHook) Option {
	return func(s *suture.Spec) {
		s.EventHook = eventHook
	}
}

func WithFailureDecay(failureDecay float64) Option {
	return func(s *suture.Spec) {
		s.FailureDecay = failureDecay
	}
}

func WithFailureThreshold(failureThreshold float64) Option {
	return func(s *suture.Spec) {
		s.FailureThreshold = failureThreshold
	}
}

func WithFailureBackoff(failureBackoff time.Duration) Option {
	return func(s *suture.Spec) {
		s.FailureBackoff = failureBackoff
	}
}

func WithBackoffJitter(backoffJitter suture.Jitter) Option {
	return func(s *suture.Spec) {
		s.BackoffJitter = backoffJitter
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *suture.Spec) {
		s.Timeout = timeout
	}
}

func WithPassThroughPanics(passThroughPanics bool) Option {
	return func(s *suture.Spec) {
		s.PassThroughPanics = passThroughPanics
	}
}

func WithDontPropagateTermination(dontPropagateTermination bool) Option {
	return func(s *suture.Spec) {
		s.DontPropagateTermination = dontPropagateTermination
	}
}
