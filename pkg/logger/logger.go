package logger

import (
	"github.com/fatih/color"
)

var colorLog *color.Color = color.New()

func LogError(message string) {
	colorLog.Add(color.FgRed)
	colorLog.Println("| ERROR |", message)
}

func LogWarning(message string) {
	colorLog.Add(color.FgYellow)
	colorLog.Println("| WARNING |", message)
}

func LogInfo(message string) {
	colorLog.Add(color.FgBlue)
	colorLog.Println("| INFO |", message)
}

func LogSuccess(message string) {
	colorLog.Add(color.FgGreen)
	colorLog.Println("| SUCCESS |", message)
}
