package examples_test

import "testing"
import "github.com/neollie/examples/patterns/Ex5_Timeout"

/*
=== RUN   TestRunReadTimeout
-Wait for 0-th string
2015/09/21 15:42:16 prod2 - Moni: 0
-Wait for 1-th string
2015/09/21 15:42:16 prod2 - Moni: 1
-Wait for 2-th string
2015/09/21 15:42:17 prod2 - Moni: 2
-Wait for 3-th string
2015/09/21 15:42:19 prod2 - Moni: 3
-Wait for 4-th string
2015/09/21 15:42:22 prod2 - Moni: 4
-Wait for 5-th string
2015/09/21 15:42:26 prod2 - Moni: 5
-Wait for 6-th string
!!!Timeout for phase 6
-Wait for 7-th string
2015/09/21 15:42:31 prod2 - Moni: 6
-Wait for 8-th string
!!!Timeout for phase 8
-Wait for 9-th string
2015/09/21 15:42:37 prod2 - Moni: 7

 */
func TestRunReadTimeout(t *testing.T) {
	examples.RunReadTimeout()
}
