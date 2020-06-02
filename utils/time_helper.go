package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseDuration(a string) (time.Duration, error) {
	if len(a) == 0 {
		r, _ := time.ParseDuration("0h")
		return r, nil
	}
	if strings.Index(a, "d") > 0 {
		temp := strings.Split(a, "d")
		if len(temp) > 2 {
			return 0, errors.New("duration string has more than one d")
		}
		days := temp[0]
		hms := temp[1]

		i, err := strconv.ParseInt(days, 10, 64)
		if err != nil {
			return 0, err
		}
		r, err := time.ParseDuration(fmt.Sprintf("%vh", i*24))
		if err != nil {
			return 0, err
		}
		hmsDuration, err := ParseDuration(hms)
		if err != nil {
			return 0, err
		}
		result := r + hmsDuration
		return result, nil
	}
	return time.ParseDuration(a)
}
