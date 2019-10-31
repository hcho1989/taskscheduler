package schedule

import (
	"errors"
	"fmt"
	"time"

	"project.scmp.tech/technology/newsroom-system/assignment/cmd/cronjob/lib/period"

	"project.scmp.tech/technology/newsroom-system/assignment/cmd/cronjob/lib/pattern"
	"project.scmp.tech/technology/newsroom-system/assignment/cmd/cronjob/lib/task"
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
		fmt.Println("Checking Plan %v instance %d", planName, i)
		period, err := s.ResolveCurrentPeriod(currentTime)
		if err != nil {
			fmt.Errorf("Error when resolving period, skip, error: %s", err.Error())
			continue
		}
		durationFromPstart, err := time.ParseDuration(instance)
		if err != nil {
			fmt.Errorf("Fail to parse start instance, skip, error: %s", instance, err.Error())
			continue
		}
		scheduleTime := period.GetPeriodStart().Add(durationFromPstart)

		ok, err := beforeExecute(planName, scheduleTime, currentTime)
		if ok {
			fmt.Println("Running Task")
			success, err = t.Execute(period)
			if err != nil {
				fmt.Errorf("Error when executing task, error: %s", err.Error())
			}
			fmt.Println("Finished, success: %v", success)
			afterExecute(planName, success)
		} else {
			fmt.Println("Skipped")
		}
	}
}

func (s Schedule) ResolveCurrentPeriod(currentTime time.Time) (period.PeriodInterface, error) {
	p, err := s.Pattern.ResolveCurrentPeriod(currentTime)
	if err != nil {
		return nil, err
	}

	pEnd := p.GetPeriodEnd()
	if currentTime.After(pEnd) {
		return nil, errors.New("Time now is after period end")
	}
	return p, nil
}

func SetBeforeExecute(f func(string, time.Time, time.Time) (bool, error)) {
	beforeExecute = f
}

func SetAfterExecute(f func(string, bool)) {
	afterExecute = f
}
