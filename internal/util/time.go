package util

import (
	"fmt"
	"time"
)

func RenderExactFromNow(date time.Time) string {
	currentTime := time.Now()
	duration := currentTime.Sub(date)

	totalSeconds := int(duration.Seconds())
	totalMinutes := totalSeconds / 60
	totalHours := totalMinutes / 60
	totalDays := totalHours / 24
	totalMonths := totalDays / 30
	totalYears := totalMonths / 12

	yearsDiff := totalYears
	monthsDiff := totalMonths % 12
	daysDiff := totalDays % 30
	hoursDiff := totalHours % 24
	minutesDiff := totalMinutes % 60
	secondsDiff := totalSeconds % 60
	var fromNow string

	if yearsDiff > 0 {
		fromNow = fmt.Sprintf(
			"%d %s & %d %s",
			yearsDiff,
			pluralize(yearsDiff, "year"),
			monthsDiff,
			pluralize(monthsDiff, "month"),
		)
	} else if monthsDiff > 0 {
		fromNow = fmt.Sprintf("%d %s & %d %s", monthsDiff, pluralize(monthsDiff, "month"), daysDiff, pluralize(daysDiff, "day"))
	} else if daysDiff > 0 {
		fromNow = fmt.Sprintf("%d %s & %dh", daysDiff, pluralize(daysDiff, "day"), hoursDiff)
	} else if hoursDiff > 0 {
		fromNow = fmt.Sprintf("%dh & %dm", hoursDiff, minutesDiff)
	} else if minutesDiff > 0 {
		fromNow = fmt.Sprintf("%dm & %ds", minutesDiff, secondsDiff)
	} else {
		fromNow = fmt.Sprintf("%ds", secondsDiff)
	}

	return fromNow + " ago"
}

func pluralize(value int, unit string) string {
	if value == 1 {
		return unit
	}
	return unit + "s"
}
