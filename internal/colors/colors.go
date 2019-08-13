package colors

import (
	"fmt"
)

// parameter represents an SGR parameter
type parameter int

// Base parameters
const (
	reset parameter = iota
	bold
)

// Foreground color parameters
const (
	redFg parameter = 31 + iota
	greenFg
	yellowFg
	blueFg
	magentaFg
	cyanFg
)

const esc = "\x1b"

func sequence(p parameter) string {
	return fmt.Sprintf("%s[%dm", esc, p)
}

// Bold returns the input string bolded
func Bold(s string) string {
	return styleString(s, bold)
}

// Red returns the input string with a red foreground
func Red(s string) string {
	return styleString(s, redFg)
}

// Green returns the input string with a green foreground
func Green(s string) string {
	return styleString(s, greenFg)
}

// Yellow returns the input string with a yellow foreground
func Yellow(s string) string {
	return styleString(s, yellowFg)
}

// Blue returns the input string with a blue foreground
func Blue(s string) string {
	return styleString(s, blueFg)
}

// Magenta returns the input string with a magenta foreground
func Magenta(s string) string {
	return styleString(s, magentaFg)
}

// Cyan returns the input string with a cyan foreground
func Cyan(s string) string {
	return styleString(s, cyanFg)
}

// styleString returns the input string preceded by the SGR sequence
// containing the parameter and terminated with the reset sequence
func styleString(s string, p parameter) string {
	return sequence(p) + s + sequence(reset)
}
