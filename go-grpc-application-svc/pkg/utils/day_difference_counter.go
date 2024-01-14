package utils

import (
	"time"
)

// make GetDaysDiff function
func GetDaysDiff(startDate, endDate string) int {
	layout := "2006-01-02T15:04:05Z"
	startDateObj, _ := time.Parse(layout, startDate)
	endDateObj, _ := time.Parse(layout, endDate)
	difference := endDateObj.Sub(startDateObj)
	return int(difference.Hours() / 24)
}
