package plan

import (
	"os"
	"testing"
	"time"

	"github.com/hcho1989/taskscheduler/mocks"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestExecute(t *testing.T) {
	smock := &mocks.ScheduleInterface{}
	tmock := &mocks.TaskInterface{}
	pName1 := "Example 1"
	pName2 := "Example 2"
	forceName := "uer9gh7"
	forceInstance := 3
	noRecord := 1
	plans := []Plan{
		{
			Name:     pName1,
			Schedule: smock,
			Task:     tmock,
		},
		{
			Name:     pName2,
			Schedule: smock,
			Task:     tmock,
		},
	}
	now := time.Now()
	smock.On("Execute", pName1, tmock, forceName, forceInstance, noRecord, now).Return(nil)
	smock.On("Execute", pName2, tmock, forceName, forceInstance, noRecord, now).Return(nil)

	t.Run("plan.Execute executes all schedule", func(t *testing.T) {
		Execute(plans, forceName, forceInstance, noRecord, now)
		smock.AssertCalled(t, "Execute", pName1, tmock, forceName, forceInstance, noRecord, now)
		smock.AssertCalled(t, "Execute", pName2, tmock, forceName, forceInstance, noRecord, now)
	})
}
