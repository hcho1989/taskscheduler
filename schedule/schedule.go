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

	for i, offset := range s.Instances {
		if currentTime.After(s.EndAt) {
			fmt.Println("Time now is after schedule.EndAt, skipped.", s.EndAt, currentTime)
			continue
		}
		if currentTime.Before(s.StartFrom) {
			fmt.Println("Time now is before schedule.StartFrom, skipped.", s.StartFrom, currentTime)
			continue
		}
		success := false
		fmt.Printf("Checking Plan %v offset %d\n", planName, i)

		if s.Pattern.IsBeyondPattern(currentTime) {
			fmt.Println("current time lies beyond the defined pattern, skipped.")
			continue
		}
		instance := s.Pattern.ResolveInstance(currentTime)
		offsetDur, err := time.ParseDuration(offset)
		if err != nil {
			fmt.Printf("Fail to parse start offset %s, skipped, error: %s\n", offset, err.Error())
			continue
		}
		scheduleTime := instance.Add(offsetDur)

		ok, err := beforeExecute(planName, scheduleTime, currentTime)
		if ok {
			fmt.Printf("Running Task: %s\n", planName)
			success, err = t.Execute(scheduleTime)
			if err != nil {
				fmt.Printf("Error when executing task, error: %s\n", err.Error())
			}
			fmt.Printf("%s Finished, success: %v\n", planName, success)
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
