package taskscheduler

import (
	"time"

	"github.com/hcho1989/taskscheduler/plan"
)

func Execute(plans []plan.Plan, currentTime time.Time) {
	for _, p := range plans {
		p.Execute(currentTime)
	}
}
