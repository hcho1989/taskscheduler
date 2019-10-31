package schedule

import (
	"os"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	testify "github.com/stretchr/testify/mock"

	"project.scmp.tech/technology/newsroom-system/assignment/cmd/cronjob/mocks"
	"project.scmp.tech/technology/newsroom-system/assignment/cmd/cronjob/period"
	"project.scmp.tech/technology/newsroom-system/assignment/pkg/lib/models"
	repoMocks "project.scmp.tech/technology/newsroom-system/assignment/pkg/lib/repositories/mocks"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSchedule_Execute(t *testing.T) {
	db := &gorm.DB{}
	now, _ := time.Parse(time.RFC3339, "2019-05-21T00:00:00+00:00") // HKT Tuesday 0000
	tmock := &mocks.TaskInterface{}
	histRepoMock := &repoMocks.ScheduleTaskHistoryRepo{}
	pName := "Example"
	forceName := ""
	forceInstance := 0
	noRecord := 0
	created := &models.ScheduleTaskHistory{
		Name:    pName,
		Success: true,
		Model: models.Model{
			CreatedAt: now,
		},
	}
	startFrom, _ := time.Parse(time.RFC3339, "0000-01-01T00:00:00+00:00")
	endAt, _ := time.Parse(time.RFC3339, "3000-12-31T23:59:59+00:00")
	pStart, _ := time.Parse(time.RFC3339, "2019-05-20T00:00:00+00:00")
	pEnd, _ := time.Parse(time.RFC3339, "2019-05-24T00:00:00+00:00")

	s := Schedule{
		StartFrom: startFrom,
		EndAt:     endAt,
		Period: &period.Static{
			Start: pStart,
			End:   pEnd,
		},
		Instances: []string{"0h", "1m", "49h"}, // 0hour from start of week, 1 minute from start of week, 49 hour from start of week
	}
	SetHistRepo(histRepoMock)
	SetDB(db)
	histRepoMock.On("LastRun", db, pName).Return(nil, nil).Once()     //.Return(created, nil)
	histRepoMock.On("LastRun", db, pName).Return(created, nil).Once() //.Return(created, nil)
	histRepoMock.On("Create", db, testify.AnythingOfType("*models.ScheduleTaskHistory")).Return(created, nil)
	tmock.On("Execute", s.Period).Return(true, nil)
	t.Run("schedule.Execute executes all past instances", func(t *testing.T) {
		// expected behavior:
		// first time, last run not found, execute
		// second time, last run found created just now, after schedule time, do not execute
		// third time, current time before schedule time, do not check last run, do not execute
		s.Execute(pName, tmock, forceName, forceInstance, noRecord, now)
		histRepoMock.AssertNumberOfCalls(t, "LastRun", 2)
		histRepoMock.AssertCalled(t, "LastRun", db, pName)
		tmock.AssertNumberOfCalls(t, "Execute", 1)
		tmock.AssertCalled(t, "Execute", s.Period)
	})

}
