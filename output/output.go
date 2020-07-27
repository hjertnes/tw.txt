// Package output contains various terminal output helper functions
package output

import "fmt"

// Red prints in text in ...
func Red(s string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", s)
}

// Green prints in text in ...
func Green(s string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", s)
}

// BoldGreen prints in text in ...
func BoldGreen(s string) string {
	return fmt.Sprintf("\033[32;1m%s\033[0m", s)
}

// Blue prints in text in ...
func Blue(s string) string {
	return fmt.Sprintf("\033[34m%s\033[0m", s)
}
