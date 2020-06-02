package plan

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/hcho1989/taskscheduler/pattern"
	"github.com/hcho1989/taskscheduler/schedule"
	"github.com/hcho1989/taskscheduler/task"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PlanTestSuite struct {
	suite.Suite
}

func (suite *PlanTestSuite) SetupSuite() {
	task.Register("TaskAbc", TaskAbc{})
}

func (suite *PlanTestSuite) TestPlan_UnmarshalJSON() {
	t := suite.T()

	schedule.SetTimeZone("Asia/Hong_Kong")
	tz, _ := time.LoadLocation("Asia/Hong_Kong")
	t.Run("UnmarshalJSON construct week schedule and task", func(t *testing.T) {
		j := `{
			"Name": "ewrgioher",
			"Schedule": {
				"Pattern": {
					"Type": "WEEK"
				},
				"StartFrom": "2019-07-01T00:00:01+08:00",
				"EndAt": "2019-07-01T00:00:02+08:00",
				"Instances": ["0h", "1h"]
			},
			"Task": {
				"Name": "TaskAbc",
				"Params": {
					"Abc": 1,
					"Def": "abcdefg"
				}
			}
		}`
		startFrom, _ := time.Parse(time.RFC3339, "2019-07-01T00:00:01+08:00")
		endAt, _ := time.Parse(time.RFC3339, "2019-07-01T00:00:02+08:00")

		var p Plan
		expected := Plan{
			Name: "ewrgioher",
			Schedule: &schedule.Schedule{
				StartFrom: startFrom,
				EndAt:     endAt,
				Pattern:   pattern.BuildWeekPattern(time.Now().In(tz)),
				Instances: []string{"0h", "1h"},
			},
			Task: &TaskAbc{
				Abc: 1,
				Def: "abcdefg",
			},
		}
		err := json.Unmarshal([]byte(j), &p)
		assert.Nil(t, err)
		assert.EqualValues(t, expected, p)
	})
	t.Run("UnmarshalJSON construct day schedule and task", func(t *testing.T) {
		j := `{
			"Name": "ewrgioher",
			"Schedule": {
				"Pattern": {
					"Type": "DAY"
				},
				"StartFrom": "2019-07-01T00:00:01+08:00",
				"EndAt": "2019-07-01T00:00:02+08:00",
				"Instances": ["0h", "1h"]
			},
			"Task": {
				"Name": "TaskAbc",
				"Params": {
					"Abc": 1,
					"Def": "abcdefg"
				}
			}
		}`
		startFrom, _ := time.Parse(time.RFC3339, "2019-07-01T00:00:01+08:00")
		endAt, _ := time.Parse(time.RFC3339, "2019-07-01T00:00:02+08:00")

		var p Plan
		expected := Plan{
			Name: "ewrgioher",
			Schedule: &schedule.Schedule{
				StartFrom: startFrom,
				EndAt:     endAt,
				Pattern:   pattern.BuildDayPattern(time.Now().In(tz)),
				Instances: []string{"0h", "1h"},
			},
			Task: &TaskAbc{
				Abc: 1,
				Def: "abcdefg",
			},
		}
		err := json.Unmarshal([]byte(j), &p)
		assert.Nil(t, err)
		assert.EqualValues(t, expected, p)
	})
}

func TestPlanTestSuite(t *testing.T) {
	suite.Run(t, new(PlanTestSuite))
}

type TaskAbc struct {
	Abc int
	Def string
}

func (t TaskAbc) Execute(p time.Time) (bool, error) {
	return true, nil
}
