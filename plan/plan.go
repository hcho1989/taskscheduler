package plan

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hcho1989/taskscheduler/pattern"
	"github.com/hcho1989/taskscheduler/schedule"
	"github.com/hcho1989/taskscheduler/task"
	"github.com/mitchellh/mapstructure"
)

var TIMEZONE, _ = time.LoadLocation("Asia/Hong_Kong")

type Plan struct {
	Name     string
	Schedule schedule.ScheduleInterface
	Task     task.TaskInterface
}

type planJSON struct {
	Name     string
	Schedule struct {
		Pattern struct {
			Type   string
			Params struct {
				Start    string
				End      string
				Duration string
			}
		}
		StartFrom string
		EndAt     string
		Instances []string
	}
	Task task.TaskConfig
}

func (p *Plan) UnmarshalJSON(b []byte) error {
	var s planJSON
	var start time.Time
	var end time.Time
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, err := task.Dispatch(s.Task.Name)
	mapstructure.Decode(s.Task.Params, t)
	if err != nil {
		fmt.Println("Fail to Dispatch Task")
		return err
	}
	if len(s.Schedule.StartFrom) > 0 {
		start, err = time.Parse(time.RFC3339, s.Schedule.StartFrom)
		if err != nil {
			fmt.Println("Fail to Parse start Time")
			return err
		}
	} else {
		start, _ = time.Parse(time.RFC3339, "0000-01-01T00:00:00+00:00")
	}
	if len(s.Schedule.EndAt) > 0 {
		end, err = time.Parse(time.RFC3339, s.Schedule.EndAt)
		if err != nil {
			fmt.Println("Fail to Parse end Time")
			return err
		}

	} else {
		end, _ = time.Parse(time.RFC3339, "3000-12-31T23:59:59+00:00")
	}
	var pr pattern.PatternInterface
	switch s.Schedule.Pattern.Type {

	case pattern.WEEK:
		pr = pattern.BuildWeekPattern()
	case pattern.DAY:
		pr = pattern.BuildDayPattern()
	case pattern.EQUILENGTH:
		start, err := time.Parse(time.RFC3339, s.Schedule.Pattern.Params.Start)
		if err != nil {
			return err
		}
		duration, err := time.ParseDuration(s.Schedule.Pattern.Params.Duration)
		if err != nil {
			return err
		}
		pr = pattern.BuildEquilengthPattern(start, duration)
	case pattern.STATIC:
		fallthrough
	default:
		start, err := time.Parse(time.RFC3339, s.Schedule.Pattern.Params.Start)
		if err != nil {
			return err
		}
		end, err := time.Parse(time.RFC3339, s.Schedule.Pattern.Params.End)
		if err != nil {
			return err
		}
		duration, err := time.ParseDuration(s.Schedule.Pattern.Params.Duration)
		if err != nil {
			return err
		}
		pr = pattern.BuildStaticPattern(start, end, duration)
	}
	if err != nil {
		fmt.Println("Fail to parse %v pattern\n", s.Schedule.Pattern.Type)
		return err
	}

	*p = Plan{
		Name: s.Name,
		Schedule: schedule.Schedule{
			StartFrom: start,
			EndAt:     end,
			Pattern:   pr,
			Instances: s.Schedule.Instances,
		},
		Task: *t,
	}
	return nil
}

func (p *Plan) Execute(currentTime time.Time) {
	p.Schedule.Execute(p.Name, p.Task, currentTime)
}
