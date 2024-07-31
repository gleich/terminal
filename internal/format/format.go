package format

import "github.com/fatih/color"

var (
	UnderlinedBold = color.New(color.Bold, color.Underline).SprintfFunc()
)
