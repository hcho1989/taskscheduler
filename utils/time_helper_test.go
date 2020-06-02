package utils

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TimeHelperTestSuite struct {
	suite.Suite
}

func (suite *TimeHelperTestSuite) SetupSuite() {
}

func (suite *TimeHelperTestSuite) TestTimeHelper_ParseDuration() {
	suiteTest := suite.T()
	suiteTest.Run("string ends with d parsed as days", func(t *testing.T) {
		got, err := ParseDuration("2d")
		expected, _ := time.ParseDuration("48h")
		assert.Nil(t, err)
		assert.EqualValues(t, got, expected)
	})
	suiteTest.Run("Nd parsed as days, others coming after parsed as hms", func(t *testing.T) {
		got, err := ParseDuration("2d10h38m23s")
		expected, _ := time.ParseDuration("58h38m23s")
		assert.Nil(t, err)
		assert.EqualValues(t, got, expected)
	})
	suiteTest.Run("fail if more than one d", func(t *testing.T) {
		got, err := ParseDuration("2d2d10h38m23s")
		assert.EqualValues(t, err, errors.New("duration string has more than one d"))
		assert.EqualValues(t, got, 0)
	})
	suiteTest.Run("parse as hms", func(t *testing.T) {
		got, err := ParseDuration("10h38m23s")
		expected, _ := time.ParseDuration("10h38m23s")
		assert.Nil(t, err)
		assert.EqualValues(t, got, expected)
	})
}

func TestTimeHelperTestSuite(t *testing.T) {
	suite.Run(t, new(TimeHelperTestSuite))
}
