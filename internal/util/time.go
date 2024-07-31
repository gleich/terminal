package util

import (
	"fmt"
	"math"
	"time"
)

func RenderExactFromNow(date time.Time) string {
	currentTime := time.Now()

	yearsDiff := math.Abs(float64(date.Year() - currentTime.Year()))
	monthsDiff := math.Abs(float64(int(date.Month())-int(currentTime.Month()))) - yearsDiff*12
	daysDiff := math.Abs(float64(date.Day() - currentTime.Day()))
	hoursDiff := math.Abs(float64(date.Hour() - currentTime.Hour()))
	minutesDiff := math.Abs(float64(date.Minute() - currentTime.Minute()))
	secondsDiff := math.Abs(float64(date.Second() - currentTime.Second()))
	var fromNow string

	if yearsDiff > 0 {
		fromNow = fmt.Sprintf("%.0f %s & %.0f %s", yearsDiff, pluralize(int(yearsDiff), "year"), monthsDiff, pluralize(int(monthsDiff), "month"))
	} else if monthsDiff > 0 {
		fromNow = fmt.Sprintf("%.0f %s & %.0f %s", monthsDiff, pluralize(int(monthsDiff), "month"), daysDiff, pluralize(int(daysDiff), "day"))
	} else if daysDiff > 0 {
		fromNow = fmt.Sprintf("%.0f %s & %.0f%s", daysDiff, pluralize(int(daysDiff), "day"), hoursDiff, "h")
	} else if hoursDiff > 0 {
		fromNow = fmt.Sprintf("%.0f%s & %.0f%s", hoursDiff, "h", minutesDiff, "m")
	} else if minutesDiff > 0 {
		fromNow = fmt.Sprintf("%.0f%s & %.0f%s", minutesDiff, "m", secondsDiff, "s")
	} else {
		fromNow = fmt.Sprintf("%.0f%s", secondsDiff, "s")
	}

	return fromNow + " ago"
}

func pluralize(value int, unit string) string {
	if value == 1 {
		return unit
	}
	return unit + "s"
}
