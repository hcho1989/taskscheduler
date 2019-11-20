package tests

import (
	// "fmt"

	"errors"
	"testing"
	"time"

	"github.com/hcho1989/taskscheduler/task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TaskTestSuite struct {
	suite.Suite
}

func (suite *TaskTestSuite) SetupSuite() {
}
func (suite *TaskTestSuite) SetupTest() {
}

func (suite *TaskTestSuite) TestTask_Dispatch() {
	suiteTest := suite.T()
	suiteTest.Run("Should dispatch registered task", func(t *testing.T) {
		testTime, _ := time.Parse(time.RFC3339, "2019-10-31T00:00:01+08:00")
		testTime2, _ := time.Parse(time.RFC3339, "2019-10-31T00:00:21+08:00")

		someError := errors.New("some error occured")

		task.Register("SomeTask", someTask{})
		taskInterf, err := task.Dispatch("SomeTask")
		// taskIn
		assert.Nil(t, err, "Shall be nil")
		assert.NotNil(t, taskInterf, "Shall not be nil")
		assert.Implements(t, (*task.TaskInterface)(nil), *taskInterf, "Shall implement task.TaskInterface but it doesn't")

		success, err := (*taskInterf).Execute(testTime)
		assert.True(t, success, "Shall be True")
		assert.Nil(t, err, "Shall be True")
		success2, err2 := (*taskInterf).Execute(testTime2)
		assert.False(t, success2, "Shall be False")
		assert.EqualValues(t, someError, err2, "Shall be Equal")
	})
}

func TestTaskTestSuite(t *testing.T) {
	suite.Run(t, new(TaskTestSuite))
}

type someTask struct {
	SomeProperty string
}

func (s someTask) Execute(instance time.Time) (bool, error) {
	testTime, _ := time.Parse(time.RFC3339, "2019-10-31T00:00:01+08:00")

	if instance.Equal(testTime) {
		return true, nil
	}
	return false, errors.New("some error occured")
}
