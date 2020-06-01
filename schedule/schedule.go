package schedule

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hcho1989/taskscheduler/pattern"
	"github.com/hcho1989/taskscheduler/task"
)

type ScheduleInterface interface {
	SetBeforeExecute(f func(string, time.Time, time.Time) (bool, error))
	SetAfterExecute(f func(string, bool))
	Execute(planName string, t task.TaskInterface, currentTime time.Time)
}

type Schedule struct {
	StartFrom     time.Time
	EndAt         time.Time
	Pattern       pattern.PatternInterface
	Offset        *string // Duration from beginning of period, only support format "%sd", d stands for days
	beforeExecute func(a string, b, c time.Time) (bool, error)
	afterExecute  func(string, bool)
}

var TIMEZONE, _ = time.LoadLocation("Asia/Hong_Kong")

func (s Schedule) Execute(planName string, t task.TaskInterface, currentTime time.Time) {

	if currentTime.After(s.EndAt) {
		fmt.Println("Time now is after schedule.EndAt, skipped.", s.EndAt, currentTime)
		continue
	}
	if currentTime.Before(s.StartFrom) {
		fmt.Println("Time now is before schedule.StartFrom, skipped.", s.StartFrom, currentTime)
		continue
	}
	success := false
	fmt.Printf("Checking Plan %v offset %d\n", planName, s.Offset)

	if s.Pattern.IsBeyondPattern(currentTime) {
		fmt.Println("current time lies beyond the defined pattern, skipped.")
		continue
	}
	instance := s.Pattern.ResolveInstance(currentTime)
	offsetDur, err = calulateDuration(s.Offset)
	if err != nil {
		fmt.Printf("Fail to parse start offset %s, skipped, error: %s\n", offset, err.Error())
		continue
	}
	scheduleTime := instance.Add(offsetDur)

	ok, err := s.beforeExecute(planName, scheduleTime, currentTime)
	if ok {
		fmt.Printf("Running Task: %s\n", planName)
		success, err = t.Execute(scheduleTime)
		if err != nil {
			fmt.Printf("Error when executing task, error: %s\n", err.Error())
		}
		fmt.Printf("%s Finished, success: %v\n", planName, success)
		s.afterExecute(planName, success)
	} else {
		fmt.Println("Skipped")
	}
}

func (s *Schedule) SetBeforeExecute(f func(string, time.Time, time.Time) (bool, error)) {
	s.beforeExecute = f
}

func (s *Schedule) SetAfterExecute(f func(string, bool)) {
	s.afterExecute = f
}

func (s *Schedule) UnmarshalJSON(b []byte) error {
	var j scheduleJSON
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	var start time.Time
	var end time.Time
	var err error
	if len(j.StartFrom) > 0 {
		start, err = time.Parse(time.RFC3339, j.StartFrom)
		if err != nil {
			fmt.Println("Fail to Parse start Time")
			return err
		}
	} else {
		start, _ = time.Parse(time.RFC3339, "0000-01-01T00:00:00+00:00")
	}
	if len(j.EndAt) > 0 {
		end, err = time.Parse(time.RFC3339, j.EndAt)
		if err != nil {
			fmt.Println("Fail to Parse end Time")
			return err
		}

	} else {
		end, _ = time.Parse(time.RFC3339, "3000-12-31T23:59:59+00:00")
	}
	var pr pattern.PatternInterface
	switch j.Pattern.Type {

	case pattern.WEEK:
		pr = pattern.BuildWeekPattern(time.Now().In(TIMEZONE))
	case pattern.DAY:
		pr = pattern.BuildDayPattern(time.Now().In(TIMEZONE))
	case pattern.EQUILENGTH:
		_start, err := time.Parse(time.RFC3339, j.Pattern.Params.Start)
		if err != nil {
			return err
		}
		_duration, err := time.ParseDuration(j.Pattern.Params.Duration)
		if err != nil {
			return err
		}
		pr = pattern.BuildEquilengthPattern(_start, _duration)
	case pattern.STATIC:
		fallthrough
	default:
		_start, err := time.Parse(time.RFC3339, j.Pattern.Params.Start)
		if err != nil {
			return err
		}
		_end, err := time.Parse(time.RFC3339, j.Pattern.Params.End)
		if err != nil {
			return err
		}
		_duration, err := time.ParseDuration(j.Pattern.Params.Duration)
		if err != nil {
			return err
		}
		pr = pattern.BuildStaticPattern(_start, _end, _duration)
	}
	if err != nil {
		fmt.Printf("Fail to parse %v pattern\n", j.Pattern.Type)
		return err
	}
	s.StartFrom = start
	s.EndAt = end
	s.Pattern = pr
	s.Instances = j.Instances
	return nil
}

func SetTimeZone(tz string) error {
	t, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Println("Fail to set timezone", err)
		return err
	}
	TIMEZONE = t
	return nil
}

type scheduleJSON struct {
	Pattern   patternJSON
	StartFrom string
	EndAt     string
	Instances []string
}

type patternJSON struct {
	Type   string
	Params struct {
		Start    string
		End      string
		Duration string
	}
}

func calulateDuration(a *string) (*time.Duration, error) {
	if a == nil {
		r, _ := time.ParseDuration("0h")
		return &r, nil
	}
	if len(a) > 0 && strings.Index(a, "d") == len(a)-1 {
		// fmt.Println(int(a[:len(a)-1]))
		i, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return nil, err
		}
		r, err := time.ParseDuration(fmt.Printf("%sh", string(i*24)))
		if err != nil {
			return nil, err
		}
		return &r, nil
	}
}
