package plan

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hcho1989/taskscheduler/schedule"
	"github.com/hcho1989/taskscheduler/task"
	"github.com/mitchellh/mapstructure"
)

type Plan struct {
	Name     string
	Schedule schedule.ScheduleInterface
	Task     task.TaskInterface
}

func (p *Plan) UnmarshalJSON(b []byte) error {
	var s planJSON
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, err := task.Dispatch(s.Task.Name)
	mapstructure.Decode(s.Task.Params, t)
	if err != nil {
		fmt.Println("Fail to Dispatch Task")
		return err
	}
	p.Name = s.Name
	p.Schedule = &s.Schedule
	p.Task = t
	return nil
}

func (p *Plan) Execute(currentTime time.Time) {
	p.Schedule.Execute(p.Name, p.Task, curentTime)
}

type planJSON struct {
	Name     string
	Schedule schedule.Schedule
	Task     taskJSON
}
type taskJSON struct {
	Name   string
	Params map[string]interface{}
}
