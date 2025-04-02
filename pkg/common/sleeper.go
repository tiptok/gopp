package common

import (
	"time"
)

const (
	// DefaultSleepMin is the default minimum sleep time used by NewDefaultSleeper().
	DefaultSleepMin = time.Millisecond * 100

	// DefaultSleepMax is the default maximum sleep time used by NewDefaultSleeper().
	DefaultSleepMax = time.Second * 30
)

// Sleeper is an exponential backoff sleeper.
type Sleeper struct {
	done <-chan struct{}
	cur  time.Duration
	min  time.Duration
	max  time.Duration
}

// NewDefaultSleeper creates a new sleeper with the default sleep values.
func NewDefaultSleeper(done <-chan struct{}) *Sleeper {
	return NewSleeper(done, DefaultSleepMin, DefaultSleepMax)
}

// NewSleeper creates a new sleeper which starts with a value of min and
// on each call to Sleep increases the time sleep exponentially up to max.
func NewSleeper(done <-chan struct{}, min, max time.Duration) *Sleeper {
	s := &Sleeper{
		done: done,
		cur:  min,
		min:  min,
		max:  max,
	}

	return s
}

func (s *Sleeper) Sleep() {
	t := time.NewTimer(s.cur)
	select {
	case <-s.done:
		t.Stop()
		return
	case <-t.C:
	}

	s.cur = s.next()
}

// Reset resets the sleeper current sleep value back to min.
func (s *Sleeper) Reset() {
	s.cur = s.min
}

// next returns the next calculated sleep duration.
func (s *Sleeper) next() time.Duration {
	// Initial exponential back off.
	nc := s.cur * 2

	return min(s.max, nc)
}

func min(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}

	return b
}

// Value returns the duration the next Sleep() call would sleep for.
func (s *Sleeper) Value() time.Duration {
	return s.cur
}
