package task

import (
	"errors"
	"reflect"

	"project.scmp.tech/technology/newsroom-system/assignment/cmd/cronjob/lib/period"
	"project.scmp.tech/technology/newsroom-system/assignment/pkg"
)

var taskRegistry = map[string]reflect.Type{}

type TaskInterface interface {
	Execute(p period.PeriodInterface) (bool, error)
}

type TaskConfig struct {
	Name   string
	Params map[string]interface{}
}

func Register(name string, t TaskInterface) {
	taskRegistry[name] = reflect.TypeOf(t)
}

func Dispatch(name string) (*TaskInterface, error) {
	regType := taskRegistry[name]
	inter := reflect.New(regType).Interface()
	if taskInter, ok := inter.(TaskInterface); ok {
		return &taskInter, nil
	} else {
		err := errors.New("Fail to dispatch task")
		pkg.GetLogger().WithError(err).Errorf("Fail to dispatch task, task: %v", name)
		return nil, err
	}
}
