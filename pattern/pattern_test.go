package pattern

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// func TestBuildWeekPattern(t *testing.T) {
// 	i := BuildWeekPattern()
// 	n := time.Now()
// 	expectStart := now.BeginningOfWeek()
// 	expectEnd := expectStart.Add(7 * 24 * time.Hour)
// 	assert.False(t, i.IsPassed(n), "isPassed is not false")
// 	assert.True(t, i.ResolveCurrentPeriodStart(n).Equal(expectStart), "start not equal")
// 	assert.True(t, i.ResolveCurrentPeriodEnd(n).Equal(expectEnd), "end not equal")

// 	assert.True(t, i.ResolveCurrentPeriodStart(n).Equal(expectStart), "start not equal")
// 	assert.True(t, i.ResolveCurrentPeriodEnd(n).Equal(expectEnd), "end not equal")
// }

// func TestBuildEquilengthPeriod(t *testing.T) {
// 	pparam := map[string]interface{}{
// 		"Start":    "2019-05-19T00:00:00+00:00",
// 		"Duration": "15h",
// 	}
// 	n, _ := time.Parse(time.RFC3339, "2019-05-21T00:00:00+00:00")           // HKT Tuesday 0000
// 	expectStart, _ := time.Parse(time.RFC3339, "2019-05-20T21:00:00+00:00") // HKT Sunday 0000, start of week
// 	expectEnd, _ := time.Parse(time.RFC3339, "2019-05-21T12:00:00+00:00")   // HKT Sunday 0000, start of week
// 	pr, err := BuildEquilengthPeriod(pparam, n)
// 	assert.Nil(t, err, "Expect err to be nil, found not nil")
// 	assert.True(t, pr.ResolveStart().Equal(expectStart), "start not equal")
// 	assert.True(t, pr.ResolveEnd().Equal(expectEnd), "end not equal")
// }

// func TestBuildStaticPeriod(t *testing.T) {
// 	pparam := map[string]interface{}{
// 		"Start": "2019-05-20T21:00:00+00:00",
// 		"End":   "2019-05-21T12:00:00+00:00",
// 	}
// 	expectStart, _ := time.Parse(time.RFC3339, "2019-05-20T21:00:00+00:00") // HKT Sunday 0000, start of week
// 	expectEnd, _ := time.Parse(time.RFC3339, "2019-05-21T12:00:00+00:00")   // HKT Sunday 0000, start of week
// 	pr, err := BuildStaticPeriod(pparam)
// 	assert.Nil(t, err, "Expect err to be nil, found not nil")
// 	assert.True(t, pr.ResolveStart().Equal(expectStart), "start not equal")
// 	assert.True(t, pr.ResolveEnd().Equal(expectEnd), "end not equal")
// }
