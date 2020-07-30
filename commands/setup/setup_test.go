package setup

import (
	"os"
	"testing"
)

func TestTest(t *testing.T) {
	_ = os.Setenv("TEST", "true")

	New().Execute()

	_ = os.Setenv("TEST", "")
}