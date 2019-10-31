package schedule

import (
	"fmt"
	"time"

	"github.com/hcho1989/taskscheduler/pattern"
	"github.com/hcho1989/taskscheduler/task"
)

var beforeExecute = (func(a string, b, c time.Time) (bool, error) { return true, nil })
var afterExecute = (func(string, bool) { return })

type ScheduleInterface interface {
	Execute(planName string, t task.TaskInterface, currentTime time.Time)
}

type Schedule struct {
	StartFrom time.Time
	EndAt     time.Time
	Pattern   pattern.PatternInterface
	Instances []string // Duration from beginning of period
}

func (s Schedule) Execute(planName string, t task.TaskInterface, currentTime time.Time) {

	for i, instance := range s.Instances {
		if currentTime.After(s.EndAt) {
			fmt.Println("Time now is after schedule.EndAt, skipped", s.EndAt, currentTime)
			continue
		}
		if currentTime.Before(s.StartFrom) {
			fmt.Println("Time now is before schedule.StartFrom, skipped", s.StartFrom, currentTime)
			continue
		}
		success := false
		fmt.Printf("Checking Plan %v instance %d\n", planName, i)

		if s.Pattern.IsBeyondPattern(currentTime) {
			fmt.Println("current time lies beyond the defined pattern")
			continue
		}
		period, err := s.Pattern.ResolveCurrentPeriod(currentTime)
		if err != nil {
			fmt.Printf("Error when resolving period, skip, error: %s\n", err.Error())
			continue
		}
		durationFromPstart, err := time.ParseDuration(instance)
		if err != nil {
			fmt.Printf("Fail to parse start instance %s, skip, error: %s\n", instance, err.Error())
			continue
		}
		scheduleTime := period.GetPeriodStart().Add(durationFromPstart)

		ok, err := beforeExecute(planName, scheduleTime, currentTime)
		if ok {
			fmt.Println("Running Task")
			success, err = t.Execute(period)
			if err != nil {
				fmt.Printf("Error when executing task, error: %s\n", err.Error())
			}
			fmt.Printf("Finished, success: %v\n", success)
			afterExecute(planName, success)
		} else {
			fmt.Println("Skipped")
		}
	}
}

func SetBeforeExecute(f func(string, time.Time, time.Time) (bool, error)) {
	beforeExecute = f
}

func SetAfterExecute(f func(string, bool)) {
	afterExecute = f
}
