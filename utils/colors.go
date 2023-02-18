package utils

// colors.go - Coloring and formatting configuration and helper functions.

var headerString string = "\033[33m"
var boldString string = "\033[32m"
var keyString string = "\033[95m"
var endFormatting string = "\033[0m"

func FormatHeaderString(s string) string {
	return headerString + s + endFormatting
}

func FormatBoldString(s string) string {
	return boldString + s + endFormatting
}

func FormatKeyString(s string) string {
	return keyString + s + endFormatting
}

// EOF
