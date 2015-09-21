package examples_test

import "testing"
import "github.com/neollie/examples/patterns/Ex5_Timeout"

/*
=== RUN   TestRunTimeout
2015/09/21 15:26:46 prod2 - Moni: 0
!!!Timeout for phase 1
--- PASS: TestRunTimeout (3.00s)

 */
func TestRunTimeout(t *testing.T) {
	examples.RunTimeout()
}
