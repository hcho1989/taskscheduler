package period

import (
	"time"
)

type PeriodInterface interface {
	GetPeriodStart() time.Time
	GetPeriodEnd() time.Time
}
type Period struct {
	Start time.Time
	End   time.Time
}

func (p *Period) GetPeriodStart() time.Time {
	return p.Start
}
func (p *Period) GetPeriodEnd() time.Time {
	return p.End
}

// type Week struct {
// 	CurrentTime time.Time
// }

// type Day struct {
// 	CurrentTime time.Time
// }

// func (p *Day) ResolveStart() time.Time {
// 	return now.New(p.CurrentTime).BeginningOfDay()
// }
// func (p *Day) ResolveEnd() time.Time {
// 	return now.New(p.CurrentTime).EndOfDay()
// }

// type EquiLength struct {
// 	CurrentTime time.Time
// 	Start       time.Time // WARNING: start must be a past time, otherwise panic
// 	Duration    time.Duration
// }

// func (p *EquiLength) ResolveStart() time.Time {
// 	if p.CurrentTime.Before(p.Start) {
// 		panic(errors.New("Start must be a past time"))
// 	}
// 	var m time.Duration
// 	m = time.Duration((int64(p.CurrentTime.Sub(p.Start)) / int64(p.Duration))) * p.Duration

// 	return p.Start.Add(m)
// }
// func (p *EquiLength) ResolveEnd() time.Time {
// 	return p.ResolveStart().Add(p.Duration)
// }

// type Static struct {
// 	Start time.Time
// 	End   time.Time
// }

// func (p *Static) ResolveStart() time.Time {
// 	return p.Start
// }
// func (p *Static) ResolveEnd() time.Time {
// 	return p.End
// }

// func BuildCalendarPeriod(p map[string]interface{}, n time.Time) (PeriodInterface, error) {
// 	if t, ok := p["Period"]; ok {
// 		switch t {
// 		case WEEK:
// 			return &Week{
// 				CurrentTime: n,
// 			}, nil
// 		case DAY:
// 			fallthrough
// 		default:
// 			return &Day{
// 				CurrentTime: n,
// 			}, nil
// 		}
// 	}

// 	return nil, errors.New("Fail to build calendar period")
// }

// func BuildEquilengthPeriod(p map[string]interface{}, n time.Time) (PeriodInterface, error) {
// 	start, ok := p["Start"]
// 	if !ok {
// 		return nil, errors.New("Fail to parse equilength start")
// 	}
// 	duration, ok := p["Duration"]
// 	if !ok {
// 		return nil, errors.New("Fail to parse equilength duration")
// 	}

// 	s, err := time.Parse(time.RFC3339, start.(string))
// 	if err != nil {
// 		return nil, errors.New("Fail to parse equilength start")
// 	}
// 	d, err := time.ParseDuration(duration.(string))
// 	if err != nil {
// 		return nil, errors.New("Fail to parse equilength duration")
// 	}

// 	return &EquiLength{
// 		CurrentTime: n,
// 		Start:       s,
// 		Duration:    d,
// 	}, nil
// }

// func BuildStaticPeriod(p map[string]interface{}) (PeriodInterface, error) {
// 	start, ok := p["Start"]
// 	if !ok {
// 		return nil, errors.New("Fail to parse equilength start")
// 	}
// 	end, ok := p["End"]
// 	if !ok {
// 		return nil, errors.New("Fail to parse equilength duration")
// 	}

// 	s, err := time.Parse(time.RFC3339, start.(string))
// 	if err != nil {
// 		return nil, errors.New("Fail to parse static start")
// 	}
// 	e, err := time.Parse(time.RFC3339, end.(string))
// 	if err != nil {
// 		return nil, errors.New("Fail to parse static end")
// 	}

// 	return &Static{
// 		Start: s,
// 		End:   e,
// 	}, nil
// }
