package examples_test

import (
	"testing"
	"github.com/neollie/examples/patterns/Ex4_FunIn"
)

/**
=== RUN   TestRunAsync2
2015/09/21 14:55:24 prod2 - Moni: 0
2015/09/21 14:55:24 prod2 - Moni: 1
2015/09/21 14:55:24 prod2 - Dia: 0
2015/09/21 14:55:24 prod2 - Moni: 2
2015/09/21 14:55:24 prod2 - Moni: 3
2015/09/21 14:55:24 prod2 - Dia: 1
2015/09/21 14:55:24 prod2 - Dia: 2
2015/09/21 14:55:24 prod2 - Moni: 4
2015/09/21 14:55:24 prod2 - Moni: 5
2015/09/21 14:55:24 prod2 - Moni: 6
--- PASS: TestRunAsync2 (0.00s)

 */
func TestRunAsync2(t *testing.T) {
	examples.RunAsync2()
}