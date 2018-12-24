package github

import "time"

func ConvertTimeToTimestamp(timeString string) (int64, error) {
	dateTime, err := time.Parse(time.RFC3339, timeString)
	return dateTime.Unix(), err
}
