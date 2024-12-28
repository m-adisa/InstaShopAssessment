package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("TESTING", "true")

	// Run tests
	code := m.Run()

	// Exit
	os.Exit(code)
}
