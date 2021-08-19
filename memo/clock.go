package memo

// A Clock represents the passage of time, it can provide the current time
// in nanoseconds, which could be a relative value, not an absolute value.
type Clock interface {
	// Now returns the current time in nanoseconds.
	Now() int64
}

// A RealClock can provide the real current time.
type RealClock struct{}

// NewRealClock creates a real clock.
func NewRealClock() RealClock {
	return RealClock{}
}

// Now returns the real current time in nanoseconds.
func (rc RealClock) Now() int64 {
	return nanotime()
}
