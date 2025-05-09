package utils

import (
	"testing"
)

// It is a helper function to run a test.
func It(ts *testing.T, name string, fn func(t *testing.T)) {
	ts.Run(name, fn)
}
