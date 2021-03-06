package task

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

var taskRegistry = map[string]reflect.Type{}

type TaskInterface interface {
	Execute(instance time.Time) (bool, error)
}

func Register(name string, t TaskInterface) {
	taskRegistry[name] = reflect.TypeOf(t)
}

func Dispatch(name string) (TaskInterface, error) {
	regType := taskRegistry[name]
	inter := reflect.New(regType).Interface()
	if taskInter, ok := inter.(TaskInterface); ok {
		return taskInter, nil
	} else {
		err := errors.New("Fail to dispatch task")
		fmt.Printf("Fail to dispatch task, task: %v\n", name)
		return nil, err
	}
}
