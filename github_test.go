package github_test

import (
	. "github-api-driver"
	"testing"
)

func Test_ConvertTimeToTimestamp_Input_Time_RFC3339_Format_Should_Be_Error(t *testing.T) {
	inputTimeString := "1970-01-01 00:00:00"

	_, actualError := ConvertTimeToTimestamp(inputTimeString)

	if actualError == nil {
		t.Errorf("expected %v but it got %v", nil, actualError)
	}
}

func Test_ConvertTimeToTimestamp_Input_Time_RFC3339_Format_Should_Be_Unix_Time(t *testing.T) {
	expectedUnixTime := int64(0)
	inputTimeString := "1970-01-01T00:00:00Z"

	actualUnixTime, _ := ConvertTimeToTimestamp(inputTimeString)

	if expectedUnixTime != actualUnixTime {
		t.Errorf("expect '%d' but it got '%d'", expectedUnixTime, actualUnixTime)
	}
}
