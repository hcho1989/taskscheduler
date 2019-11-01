package pattern

import (
	"errors"
	"sort"
	"time"

	"github.com/jinzhu/now"
)

const WEEK = "WEEK"
const DAY = "DAY"
const EQUILENGTH = "EQUILENGTH"
const STATIC = "STATIC"

// Pattern can resolve the Start and End of the current period
type PatternInterface interface {
	IsBeyondPattern(n time.Time) bool
	ResolveInstance(n time.Time) time.Time
}

type InfiniteEquilength struct {
	Instance time.Time
	Duration time.Duration
}

func (i *InfiniteEquilength) IsBeyondPattern(n time.Time) bool {
	return false
}

func (i *InfiniteEquilength) ResolveInstance(n time.Time) time.Time {
	return resolveInstance(i.Instance, n, i.Duration)
}

type FiniteEquilength struct {
	Start    time.Time
	Duration time.Duration
	End      time.Time
}

func (f *FiniteEquilength) IsBeyondPattern(n time.Time) bool {
	return n.Before(f.Start) || n.After(f.End)
}

func (f *FiniteEquilength) ResolveInstance(n time.Time) time.Time {
	return resolveInstance(f.Start, n, f.Duration)
}

func BuildWeekPattern() PatternInterface {
	return &InfiniteEquilength{
		Instance: now.BeginningOfWeek(),
		Duration: time.Duration(7 * 24 * time.Hour),
	}
}

func BuildDayPattern() PatternInterface {
	return &InfiniteEquilength{
		Instance: now.BeginningOfDay(),
		Duration: time.Duration(24 * time.Hour),
	}
}

func BuildEquilengthPattern(start time.Time, duration time.Duration) PatternInterface {
	return &InfiniteEquilength{
		Instance: start,
		Duration: duration,
	}
}

func BuildStaticPattern(start, end time.Time, dur time.Duration) PatternInterface {
	return &FiniteEquilength{
		Start:    start,
		End:      end,
		Duration: end.Sub(start),
	}
}

func resolveInstance(instance, n time.Time, dur time.Duration) time.Time {
	secFromStart := int64(n.Sub(instance))
	secDur := int64(dur)
	periodFromStart := secFromStart / secDur
	if secFromStart < 0 {
		periodFromStart--
	}
	durFromStart := time.Duration(periodFromStart) * dur
	pStart := instance.Add(durFromStart)
	return pStart
}

type Specific struct {
	instances []int // Unixnano
}

func (s *Specific) IsBeyondPattern(n time.Time) bool {
	// assume instance length > 0
	if len(s.instances) == 0 {
		panic(errors.New("no instances"))
	}
	return int(n.UnixNano()) < s.instances[0]
}

func (s *Specific) ResolveInstance(n time.Time) time.Time {
	// assume instance length > 0
	// assume n > s.instance[0]
	_n := int(n.UnixNano())
	i := sort.SearchInts(s.instances, _n)
	if i < len(s.instances) && s.instances[i] == _n {
		return n
	} else if i > 0 {
		// n to be insert at position i, take the previous one.
		return time.Unix(0, int64(s.instances[i-1]))
	} else {
		panic(errors.New("time input is beyond pattern"))
	}
}

func (s *Specific) AddInstances(i ...time.Time) {
	// add
	for _, _i := range i {
		s.instances = append(s.instances, int(_i.UnixNano()))
	}
	// deduplicate
	_map := make(map[int]bool, 0)
	for _, _i := range s.instances {
		_map[_i] = true
	}
	result := make([]int, len(_map))
	for _k := range _map {
		result = append(result, _k)
	}
	sort.Ints(result)
	s.instances = result
}
