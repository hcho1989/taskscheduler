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
