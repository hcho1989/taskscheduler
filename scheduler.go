package taskscheduler

import (
	"errors"
	"time"

	"github.com/hcho1989/taskscheduler/plan"
	"github.com/hcho1989/taskscheduler/task"
)

var defaultBeforeScheduleExecute = (func(a string, b, c time.Time) (bool, error) { return true, nil })
var defaultAfterScheduleExecute = (func(string, bool) { return })

type loaderInterface interface {
	LoadPlan() ([]plan.Plan, error)
}

type Scheduler struct {
	loader                loaderInterface
	beforeScheduleExecute func(a string, b, c time.Time) (bool, error)
	afterScheduleExecute  func(string, bool)
}

func (s *Scheduler) Execute(currentTime time.Time) {
	if s.loader == nil {
		panic(errors.New("Loader is nil"))
	}
	plans, err := s.loader.LoadPlan()
	if err != nil {
		panic(err)
	}
	for _, p := range plans {
		p.Schedule.SetBeforeExecute(s.beforeScheduleExecute)
		p.Schedule.SetAfterExecute(s.afterScheduleExecute)
		p.Execute(currentTime)
	}
}

func (s *Scheduler) init() {
	s.SetAfterScheduleExecute(defaultAfterScheduleExecute)
	s.SetBeforeScheduleExecute(defaultBeforeScheduleExecute)
}

func (s *Scheduler) SetBeforeScheduleExecute(f func(a string, b, c time.Time) (bool, error)) {
	s.beforeScheduleExecute = f
}

func (s *Scheduler) SetAfterScheduleExecute(f func(string, bool)) {
	s.afterScheduleExecute = f
}

func (s *Scheduler) SetLoader(l loaderInterface) {
	s.loader = l
}

func (s *Scheduler) RegisterTask(name string, t task.TaskInterface) {
	task.Register(name, t)
}
