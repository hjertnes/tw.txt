package output

import (
	"testing"

)

func TestBlue(t *testing.T) {
	assert.Equal(t, "\x1b[34mTest\x1b[0m", Blue("Test"))
}

func TestBoldGreen(t *testing.T) {
	assert.Equal(t, "\x1b[32;1mTest\x1b[0m", BoldGreen("Test"))
}

func TestGreen(t *testing.T) {
	assert.Equal(t, "\x1b[32mTest\x1b[0m", Green("Test"))
}

func TestRed(t *testing.T) {
	assert.Equal(t, "\x1b[31mTest\x1b[0m", Red("Test"))
}
