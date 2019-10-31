package task

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/hcho1989/taskscheduler/period"
)

type ExampleTask struct{}

func (e ExampleTask) Execute(p period.PeriodInterface) (bool, error) {
	return true, nil
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestRegister_Dispatch(t *testing.T) {

	Register("ExampleTask", reflect.TypeOf(ExampleTask{}))

	_, err := Dispatch(TaskConfig{
		Name: "ExampleTask",
	})
	// _, ok := (*expected).(ExampleTask)
	// assert.True(t, ok)
	assert.Nil(t, err)
}
