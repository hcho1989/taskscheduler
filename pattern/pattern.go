package pattern

import (
	"errors"
	"time"

	"github.com/jinzhu/now"
	"project.scmp.tech/technology/newsroom-system/assignment/cmd/cronjob/lib/period"
)

const WEEK = "WEEK"
const DAY = "DAY"
const EQUILENGTH = "EQUILENGTH"
const STATIC = "STATIC"

// Pattern can resolve the Start and End of the current period
type PatternInterface interface {
	IsPassed(n time.Time) bool
	ResolveCurrentPeriod(n time.Time) (period.PeriodInterface, error)
	// ResolveCurrentPeriodEnd(n time.Time) time.Time
}

type InfinitePattern struct {
	Start    time.Time
	Duration time.Duration
}

func (i *InfinitePattern) IsPassed(n time.Time) bool {
	return false
}

func (i *InfinitePattern) ResolveCurrentPeriod(n time.Time) (period.PeriodInterface, error) {
	p := resolveCurrentPeriod(i.Start, n, i.Duration)
	return p, nil
}

// func (i *InfinitePattern) ResolveCurrentPeriodEnd(n time.Time) time.Time {
// 	return i.ResolveCurrentPeriodStart(n).Add(i.Duration)
// }

type FinitePattern struct {
	Start    time.Time
	Duration time.Duration
	End      time.Time
}

func (f *FinitePattern) IsPassed(n time.Time) bool {
	return n.Before(f.Start) || n.After(f.End)
}

func (f *FinitePattern) ResolveCurrentPeriod(n time.Time) (period.PeriodInterface, error) {
	if f.IsPassed(n) {
		return nil, errors.New("current time lies beyond the defined pattern")
	}
	p := resolveCurrentPeriod(f.Start, n, f.Duration)
	if p.GetPeriodEnd().After(f.End) {
		return nil, errors.New("current period exceeded pattern period")
	}
	return p, nil
}

// func (f *FinitePattern) ResolveCurrentPeriodStart(n time.Time) time.Time {
// 	if f.IsPassed(n) {
// 		panic(errors.New("current time lies beyond the defined pattern"))
// 	}
// 	var m time.Duration
// 	m = time.Duration((int64(n.Sub(f.Start)) / int64(f.Duration))) * f.Duration

// 	return f.Start.Add(m)
// }
// func (f *FinitePattern) ResolveCurrentPeriodEnd(n time.Time) time.Time {
// 	if f.IsPassed(n) {
// 		panic(errors.New("current time lies beyond the defined pattern"))
// 	}
// 	return f.ResolveCurrentPeriodStart(n).Add(f.Duration)
// }

func BuildWeekPattern() PatternInterface {
	return &InfinitePattern{
		Start:    now.BeginningOfWeek(),
		Duration: time.Duration(7 * 24 * time.Hour),
	}
}

func BuildDayPattern() PatternInterface {
	return &InfinitePattern{
		Start:    now.BeginningOfDay(),
		Duration: time.Duration(24 * time.Hour),
	}
}

func BuildEquilengthPattern(start time.Time, duration time.Duration) PatternInterface {
	return &InfinitePattern{
		Start:    start,
		Duration: duration,
	}
}

func BuildStaticPattern(start, end time.Time, dur time.Duration) PatternInterface {
	return &FinitePattern{
		Start:    start,
		End:      end,
		Duration: end.Sub(start),
	}
}

func resolveCurrentPeriod(instance, n time.Time, dur time.Duration) period.PeriodInterface {
	secFromStart := int64(n.Sub(instance))
	secDur := int64(dur)
	periodFromStart := secFromStart / secDur
	if secFromStart < 0 {
		periodFromStart--
	}
	durFromStart := time.Duration(periodFromStart) * dur
	pStart := instance.Add(durFromStart)
	pEnd := pStart.Add(dur)
	return &period.Period{
		Start: pStart,
		End:   pEnd,
	}
}
