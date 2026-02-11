package circuitbreaker

import (
	"sync"
	"time"
)

type State int

const(
	Closed State = iota
	Open
	HalfOpen
)

type Breaker struct{
	mu sync.Mutex
	state State
	failures int
	failureLimit int
	lastFailure time.Time
	resetAfter time.Duration
}

func New(failureLimit int, resetAfter time.Duration) *Breaker {
	return &Breaker{
		state: Closed,
		failureLimit: failureLimit,
		resetAfter: resetAfter,
	}
}

func (b *Breaker) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.state == Open{
		if time.Since(b.lastFailure) > b.resetAfter{
			b.state= HalfOpen
			return true
		}
		return false
	}
	return true
}

func(b *Breaker)Success(){
	b.mu.Lock()
	defer b.mu.Unlock()

	b.failures = 0
	b.state = Closed
}

func(b *Breaker) Failure(){
	b.mu.Lock()
	defer b.mu.Unlock()

	b.failures++
	b.lastFailure = time.Now()

	if b.failures >= b.failureLimit {
		b.state = Open
	}
}