package env

import "github.com/fatih/color"

func GetSuccess() *color.Color {
	return getConsole(color.FgHiGreen)
}

func GetWarning() *color.Color {
	return getConsole(color.FgHiYellow)
}

func GetError() *color.Color {
	return getConsole(color.FgHiRed)
}

func GetDefault() *color.Color {
	return getConsole(color.FgHiWhite)
}

func GetDescriptor() *color.Color {
	return getConsole(color.FgHiCyan)
}

func getConsole(attribute color.Attribute) *color.Color {
	return color.New(attribute)
}
