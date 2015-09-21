package examples_test

import (
	"testing"
	"github.com/neollie/examples/patterns/Ex1"
)

// Test it with #go test -v nopatter_test.go
func TestNoPattern(t *testing.T) {
	// Run
	examples.Cons1();
	// Test log
	t.Log("Say hi!") // Logging during testing also visible only with test -v option
}