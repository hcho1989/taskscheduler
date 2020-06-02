package schedule

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/hcho1989/taskscheduler/mocks"
	"github.com/hcho1989/taskscheduler/pattern"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ScheduleTestSuite struct {
	suite.Suite
}

func (suite *ScheduleTestSuite) SetupSuite() {
	// task.Register("TaskAbc", TaskAbc{})
}

func (suite *ScheduleTestSuite) TestSchedule_UnmarshalJSON() {
	suiteTest := suite.T()

	SetTimeZone("Asia/Hong_Kong")
	tz, _ := time.LoadLocation("Asia/Hong_Kong")
	suiteTest.Run("UnmarshalJSON with instances", func(t *testing.T) {
		j := `{
			"Pattern": {
				"Type": "WEEK"
			},
			"StartFrom": "2019-07-01T00:00:01+08:00",
			"EndAt": "2019-07-01T00:00:02+08:00",
			"Instances": ["0h", "1h"]
		}`
		startFrom, _ := time.Parse(time.RFC3339, "2019-07-01T00:00:01+08:00")
		endAt, _ := time.Parse(time.RFC3339, "2019-07-01T00:00:02+08:00")

		var p Schedule
		expected := Schedule{
			StartFrom: startFrom,
			EndAt:     endAt,
			Pattern:   pattern.BuildWeekPattern(time.Now().In(tz)),
			Instances: []string{"0h", "1h"},
		}
		err := json.Unmarshal([]byte(j), &p)
		assert.Nil(t, err)
		assert.EqualValues(t, expected, p)
	})
	suiteTest.Run("UnmarshalJSON with offset", func(t *testing.T) {
		j := `{
			"Pattern": {
				"Type": "WEEK"
			},
			"StartFrom": "2019-07-01T00:00:01+08:00",
			"EndAt": "2019-07-01T00:00:02+08:00",
			"Offset": "1d20h"
		}`
		startFrom, _ := time.Parse(time.RFC3339, "2019-07-01T00:00:01+08:00")
		endAt, _ := time.Parse(time.RFC3339, "2019-07-01T00:00:02+08:00")

		var p Schedule
		expected := Schedule{
			StartFrom: startFrom,
			EndAt:     endAt,
			Pattern:   pattern.BuildWeekPattern(time.Now().In(tz)),
			Offset:    "1d20h",
		}
		err := json.Unmarshal([]byte(j), &p)
		assert.Nil(t, err)
		assert.EqualValues(t, expected, p)
	})
}

func (suite *ScheduleTestSuite) TestSchedule_Execute() {
	suiteTest := suite.T()

	suiteTest.Run("Execute once", func(t *testing.T) {
		j := `{
			"Pattern": {
				"Type": "DAY"
			},
			"Offset": "1d20h"
		}`
		var p Schedule
		currentTime, _ := time.Parse(time.RFC3339, "2019-07-01T00:00:01+08:00")
		scehduleTime, _ := time.Parse(time.RFC3339, "2019-07-02T20:00:00+08:00")
		m := &mocks.TaskInterface{}
		m.On("Execute", scehduleTime.In(TIMEZONE)).Return(true, nil)
		json.Unmarshal([]byte(j), &p)
		p.SetBeforeExecute(func(a string, b, c time.Time) (bool, error) { return true, nil })
		p.SetAfterExecute(func(a string, b bool) { return })
		p.Execute("test", m, currentTime)
		m.AssertCalled(t, "Execute", scehduleTime.In(TIMEZONE))
	})
	suiteTest.Run("Execute once with instance", func(t *testing.T) {
		j := `{
			"Pattern": {
				"Type": "DAY"
			},
			"Instances": ["1d20h"]
		}`
		var p Schedule
		currentTime, _ := time.Parse(time.RFC3339, "2019-07-01T00:00:01+08:00")
		scehduleTime, _ := time.Parse(time.RFC3339, "2019-07-02T20:00:00+08:00")
		m := &mocks.TaskInterface{}
		m.On("Execute", scehduleTime.In(TIMEZONE)).Return(true, nil)
		json.Unmarshal([]byte(j), &p)
		p.SetBeforeExecute(func(a string, b, c time.Time) (bool, error) { return true, nil })
		p.SetAfterExecute(func(a string, b bool) { return })
		p.Execute("test", m, currentTime)
		m.AssertCalled(t, "Execute", scehduleTime.In(TIMEZONE))
	})
	suiteTest.Run("Execute twice with instance", func(t *testing.T) {
		j := `{
			"Pattern": {
				"Type": "DAY"
			},
			"Instances": ["0h","1d20h"]
		}`
		var p Schedule
		currentTime, _ := time.Parse(time.RFC3339, "2019-07-01T00:00:01+08:00")
		scehduleTime, _ := time.Parse(time.RFC3339, "2019-07-01T00:00:00+08:00")
		scehduleTime2, _ := time.Parse(time.RFC3339, "2019-07-02T20:00:00+08:00")
		m := &mocks.TaskInterface{}
		m.On("Execute", scehduleTime.In(TIMEZONE)).Return(true, nil)
		m.On("Execute", scehduleTime2.In(TIMEZONE)).Return(true, nil)
		json.Unmarshal([]byte(j), &p)
		p.SetBeforeExecute(func(a string, b, c time.Time) (bool, error) { return true, nil })
		p.SetAfterExecute(func(a string, b bool) { return })
		p.Execute("test", m, currentTime)
		m.AssertCalled(t, "Execute", scehduleTime.In(TIMEZONE))
		m.AssertCalled(t, "Execute", scehduleTime2.In(TIMEZONE))
	})
	suiteTest.Run("Execute once with offset only if both offset and instances are present", func(t *testing.T) {
		j := `{
			"Pattern": {
				"Type": "DAY"
			},
			"Instances": ["0h","1d20h"],
			"Offset": "2d20h"
		}`
		var p Schedule
		currentTime, _ := time.Parse(time.RFC3339, "2019-07-01T00:00:01+08:00")
		scehduleTime2, _ := time.Parse(time.RFC3339, "2019-07-03T20:00:00+08:00")
		m := &mocks.TaskInterface{}
		m.On("Execute", scehduleTime2.In(TIMEZONE)).Return(true, nil)
		json.Unmarshal([]byte(j), &p)
		p.SetBeforeExecute(func(a string, b, c time.Time) (bool, error) { return true, nil })
		p.SetAfterExecute(func(a string, b bool) { return })
		p.Execute("test", m, currentTime)
		m.AssertCalled(t, "Execute", scehduleTime2.In(TIMEZONE))
	})
}

func TestScheduleTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleTestSuite))
}

type TaskAbc struct {
	Abc int
	Def string
}

func (t TaskAbc) Execute(p time.Time) (bool, error) {
	return true, nil
}
