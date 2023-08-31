package metrics

import (
	"sync"
	"time"
)

type TimerSnapshot interface {
	Count() int64
	Max() int64
	Mean() float64
	Min() int64
	Percentile(float64) float64
	Percentiles([]float64) []float64
	Rate1() float64
	Rate5() float64
	Rate15() float64
	RateMean() float64
	Sum() int64
	Variance() float64
	StdDev() float64
}

// Timers capture the duration and rate of events.
type Timer interface {
	Snapshot() TimerSnapshot
	Stop()
	Time(func())
	UpdateSince(time.Time)
	Update(time.Duration)
}

// GetOrRegisterTimer returns an existing Timer or constructs and registers a
// new StandardTimer.
// Be sure to unregister the meter from the registry once it is of no use to
// allow for garbage collection.
func GetOrRegisterTimer(name string, r Registry) Timer {
	if nil == r {
		r = DefaultRegistry
	}
	return r.GetOrRegister(name, NewTimer).(Timer)
}

// NewCustomTimer constructs a new StandardTimer from a Histogram and a Meter.
// Be sure to call Stop() once the timer is of no use to allow for garbage collection.
func NewCustomTimer(h Histogram, m Meter) Timer {
	if !Enabled {
		return NilTimer{}
	}
	return &StandardTimer{
		histogram: h,
		meter:     m,
	}
}

// NewRegisteredTimer constructs and registers a new StandardTimer.
// Be sure to unregister the meter from the registry once it is of no use to
// allow for garbage collection.
func NewRegisteredTimer(name string, r Registry) Timer {
	c := NewTimer()
	if nil == r {
		r = DefaultRegistry
	}
	r.Register(name, c)
	return c
}

// NewTimer constructs a new StandardTimer using an exponentially-decaying
// sample with the same reservoir size and alpha as UNIX load averages.
// Be sure to call Stop() once the timer is of no use to allow for garbage collection.
func NewTimer() Timer {
	if !Enabled {
		return NilTimer{}
	}
	return &StandardTimer{
		histogram: NewHistogram(NewExpDecaySample(1028, 0.015)),
		meter:     NewMeter(),
	}
}

// NilTimer is a no-op Timer.
type NilTimer struct{}

// Count is a no-op.
func (NilTimer) Count() int64 { return 0 }

// Max is a no-op.
func (NilTimer) Max() int64 { return 0 }

// Mean is a no-op.
func (NilTimer) Mean() float64 { return 0.0 }

// Min is a no-op.
func (NilTimer) Min() int64 { return 0 }

// Percentile is a no-op.
func (NilTimer) Percentile(p float64) float64 { return 0.0 }

// Percentiles is a no-op.
func (NilTimer) Percentiles(ps []float64) []float64 {
	return make([]float64, len(ps))
}

// Rate1 is a no-op.
func (NilTimer) Rate1() float64 { return 0.0 }

// Rate5 is a no-op.
func (NilTimer) Rate5() float64 { return 0.0 }

// Rate15 is a no-op.
func (NilTimer) Rate15() float64 { return 0.0 }

// RateMean is a no-op.
func (NilTimer) RateMean() float64 { return 0.0 }

// Snapshot is a no-op.
func (NilTimer) Snapshot() TimerSnapshot { return NilTimer{} }

// StdDev is a no-op.
func (NilTimer) StdDev() float64 { return 0.0 }

// Stop is a no-op.
func (NilTimer) Stop() {}

// Sum is a no-op.
func (NilTimer) Sum() int64 { return 0 }

// Time is a no-op.
func (NilTimer) Time(f func()) { f() }

// Update is a no-op.
func (NilTimer) Update(time.Duration) {}

// UpdateSince is a no-op.
func (NilTimer) UpdateSince(time.Time) {}

// Variance is a no-op.
func (NilTimer) Variance() float64 { return 0.0 }

// StandardTimer is the standard implementation of a Timer and uses a Histogram
// and Meter.
type StandardTimer struct {
	histogram Histogram
	meter     Meter
	mutex     sync.Mutex
}

// Snapshot returns a read-only copy of the timer.
func (t *StandardTimer) Snapshot() TimerSnapshot {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return &timerSnapshot{
		histogram: t.histogram.Snapshot().(*histogramSnapshot),
		meter:     t.meter.Snapshot().(*meterSnapshot),
	}
}

// Stop stops the meter.
func (t *StandardTimer) Stop() {
	t.meter.Stop()
}

// Record the duration of the execution of the given function.
func (t *StandardTimer) Time(f func()) {
	ts := time.Now()
	f()
	t.Update(time.Since(ts))
}

// Record the duration of an event.
func (t *StandardTimer) Update(d time.Duration) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.histogram.Update(int64(d))
	t.meter.Mark(1)
}

// Record the duration of an event that started at a time and ends now.
func (t *StandardTimer) UpdateSince(ts time.Time) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.histogram.Update(int64(time.Since(ts)))
	t.meter.Mark(1)
}

// timerSnapshot is a read-only copy of another Timer.
type timerSnapshot struct {
	histogram *histogramSnapshot
	meter     *meterSnapshot
}

// Count returns the number of events recorded at the time the snapshot was
// taken.
func (t *timerSnapshot) Count() int64 { return t.histogram.Count() }

// Max returns the maximum value at the time the snapshot was taken.
func (t *timerSnapshot) Max() int64 { return t.histogram.Max() }

// Mean returns the mean value at the time the snapshot was taken.
func (t *timerSnapshot) Mean() float64 { return t.histogram.Mean() }

// Min returns the minimum value at the time the snapshot was taken.
func (t *timerSnapshot) Min() int64 { return t.histogram.Min() }

// Percentile returns an arbitrary percentile of sampled values at the time the
// snapshot was taken.
func (t *timerSnapshot) Percentile(p float64) float64 {
	return t.histogram.Percentile(p)
}

// Percentiles returns a slice of arbitrary percentiles of sampled values at
// the time the snapshot was taken.
func (t *timerSnapshot) Percentiles(ps []float64) []float64 {
	return t.histogram.Percentiles(ps)
}

// Rate1 returns the one-minute moving average rate of events per second at the
// time the snapshot was taken.
func (t *timerSnapshot) Rate1() float64 { return t.meter.Rate1() }

// Rate5 returns the five-minute moving average rate of events per second at
// the time the snapshot was taken.
func (t *timerSnapshot) Rate5() float64 { return t.meter.Rate5() }

// Rate15 returns the fifteen-minute moving average rate of events per second
// at the time the snapshot was taken.
func (t *timerSnapshot) Rate15() float64 { return t.meter.Rate15() }

// RateMean returns the meter's mean rate of events per second at the time the
// snapshot was taken.
func (t *timerSnapshot) RateMean() float64 { return t.meter.RateMean() }

// StdDev returns the standard deviation of the values at the time the snapshot
// was taken.
func (t *timerSnapshot) StdDev() float64 { return t.histogram.StdDev() }

// Sum returns the sum at the time the snapshot was taken.
func (t *timerSnapshot) Sum() int64 { return t.histogram.Sum() }

// Variance returns the variance of the values at the time the snapshot was
// taken.
func (t *timerSnapshot) Variance() float64 { return t.histogram.Variance() }
