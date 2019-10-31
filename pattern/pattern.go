package pattern

import (
	"errors"
	"time"

	"github.com/hcho1989/taskscheduler/period"
	"github.com/jinzhu/now"
)

const WEEK = "WEEK"
const DAY = "DAY"
const EQUILENGTH = "EQUILENGTH"
const STATIC = "STATIC"

// Pattern can resolve the Start and End of the current period
type PatternInterface interface {
	IsBeyondPattern(n time.Time) bool
	ResolveCurrentPeriod(n time.Time) (period.PeriodInterface, error)
}

type InfinitePattern struct {
	Start    time.Time
	Duration time.Duration
}

func (i *InfinitePattern) IsBeyondPattern(n time.Time) bool {
	return false
}

func (i *InfinitePattern) ResolveCurrentPeriod(n time.Time) (period.PeriodInterface, error) {
	p := resolveCurrentPeriod(i.Start, n, i.Duration)
	return p, nil
}

type FinitePattern struct {
	Start    time.Time
	Duration time.Duration
	End      time.Time
}

func (f *FinitePattern) IsBeyondPattern(n time.Time) bool {
	return n.Before(f.Start) || n.After(f.End)
}

func (f *FinitePattern) ResolveCurrentPeriod(n time.Time) (period.PeriodInterface, error) {
	p := resolveCurrentPeriod(f.Start, n, f.Duration)
	if p.GetPeriodEnd().After(f.End) {
		return nil, errors.New("current period exceeded pattern period")
	}
	return p, nil
}

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
