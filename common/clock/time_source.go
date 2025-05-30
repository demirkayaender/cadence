// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package clock

import (
	"context"
	"fmt"
	"time"

	"github.com/jonboulle/clockwork"
)

type (
	// TimeSource provides an interface that packages can use instead of directly using
	// the [time] module, so that chronology-related behavior can be tested.
	TimeSource interface {
		After(d time.Duration) <-chan time.Time
		Sleep(d time.Duration)
		Now() time.Time
		Since(t time.Time) time.Duration
		NewTicker(d time.Duration) Ticker
		NewTimer(d time.Duration) Timer
		AfterFunc(d time.Duration, f func()) Timer
	}

	// Ticker provides an interface which can be used instead of directly using
	// [time.Ticker]. The real-time ticker t provides ticks through t.C which
	// becomes t.Chan() to make this channel requirement definable in this
	// interface.
	Ticker interface {
		Chan() <-chan time.Time
		Reset(d time.Duration)
		Stop()
	}

	// Timer provides an interface which can be used instead of directly using
	// [time.Timer]. The real-time timer t provides events through t.C which becomes
	// t.Chan() to make this channel requirement definable in this interface.
	Timer interface {
		Chan() <-chan time.Time
		Reset(d time.Duration) bool
		Stop() bool
	}

	// clock serves real wall-clock time
	clock struct {
		clockwork.Clock
	}

	// fakeClock serves fake controlled time
	fakeClock struct {
		*clockwork.FakeClock
	}

	// MockedTimeSource provides an interface for a clock which can be manually advanced
	// through time.
	//
	// MockedTimeSource maintains a list of "waiters," which consists of all callers
	// waiting on the underlying clock (i.e. Tickers and Timers including callers of
	// Sleep or After). Users can call BlockUntil to block until the clock has an
	// expected number of waiters.
	MockedTimeSource interface {
		TimeSource
		// Advance advances the FakeClock to a new point in time, ensuring any existing
		// waiters are notified appropriately before returning.
		Advance(d time.Duration)
		// BlockUntil blocks until the FakeClock has at least the given number
		// of waiters running at the same time.
		//
		// Waiters are either time.Sleep, time.After[Func], time.Ticker, or time.Timer,
		// and they decrement the counter when they complete or are stopped.
		BlockUntil(waiters int)
	}
)

var _ TimeSource = (*clock)(nil)
var _ TimeSource = (*fakeClock)(nil)
var _ MockedTimeSource = (*fakeClock)(nil)

// NewRealTimeSource returns a time source that servers
// real wall clock time
func NewRealTimeSource() TimeSource {
	return &clock{
		Clock: clockwork.NewRealClock(),
	}
}

// NewMockedTimeSource returns a time source that servers
// fake controlled time
func NewMockedTimeSource() MockedTimeSource {
	return &fakeClock{
		FakeClock: clockwork.NewFakeClock(),
	}
}

// NewMockedTimeSourceAt returns a time source that servers
// fake controlled time. The initial time of the MockedTimeSource will be the given time.
func NewMockedTimeSourceAt(t time.Time) MockedTimeSource {
	return &fakeClock{
		FakeClock: clockwork.NewFakeClockAt(t),
	}
}

func (r *clock) NewTicker(d time.Duration) Ticker {
	return r.Clock.NewTicker(d)
}

func (r *clock) NewTimer(d time.Duration) Timer {
	return r.Clock.NewTimer(d)
}

func (r *clock) AfterFunc(d time.Duration, f func()) Timer {
	return r.Clock.AfterFunc(d, f)
}

func (c *fakeClock) NewTicker(d time.Duration) Ticker {
	return c.FakeClock.NewTicker(d)
}

func (c *fakeClock) NewTimer(d time.Duration) Timer {
	return c.FakeClock.NewTimer(d)
}

func (c *fakeClock) AfterFunc(d time.Duration, f func()) Timer {
	return c.FakeClock.AfterFunc(d, f)
}

// WithDeadline provides a way to add a context deadline with a time controlled fashion using time source.
func WithDeadline(ctx context.Context, ts TimeSource, deadline time.Time) (context.Context, context.CancelFunc) {
	switch typedTS := ts.(type) {
	case *clock:
		return clockwork.WithDeadline(ctx, typedTS.Clock, deadline)
	case *fakeClock:
		return clockwork.WithDeadline(ctx, typedTS.FakeClock, deadline)
	default:
		panic(fmt.Sprintf("unexpected type: %T", typedTS))
	}
}

// WithTimeout provides a way to add a context deadline with a time controlled fashion using time source.
func WithTimeout(ctx context.Context, ts TimeSource, timeout time.Duration) (context.Context, context.CancelFunc) {
	return WithDeadline(ctx, ts, time.Now().Add(timeout))
}
