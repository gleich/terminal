package util

import (
	"fmt"
	"time"
)

func RenderDuration(seconds int) string {
	duration := time.Duration(seconds) * time.Second
	totalHours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	var formattedDuration string

	if totalHours > 0 {
		formattedDuration = fmt.Sprintf("%dh", totalHours)
		if minutes > 0 {
			formattedDuration += fmt.Sprintf(" & %dm", minutes)
		}
	} else if seconds < 3660 && seconds > 3540 {
		formattedDuration = "1h"
	} else {
		remainingSeconds := seconds % 60
		formattedDuration = fmt.Sprintf("%dm", minutes)
		if remainingSeconds > 0 {
			formattedDuration += fmt.Sprintf(" & %ds", remainingSeconds)
		}
	}

	return formattedDuration
}
