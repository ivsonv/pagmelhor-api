package utils

import "time"

func IsValidDate(date *time.Time) bool {
	if date == nil {
		return false
	}

	if date.Format("2006-01-02") == "0001-01-01" {
		return false
	}

	return true
}
