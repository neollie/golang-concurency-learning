package examples_test

import (
	"testing"
	"github.com/neollie/examples/patterns/Ex4_FunIn"
)

/**
=== RUN   TestRun4
2015/09/21 14:39:31 prod2 - Dia: 0
2015/09/21 14:39:31 prod2 - Dia: 1
2015/09/21 14:39:31 prod2 - Moni: 0
2015/09/21 14:39:31 prod2 - Moni: 1
2015/09/21 14:39:31 prod2 - Dia: 2
2015/09/21 14:39:31 prod2 - Moni: 2
2015/09/21 14:39:31 prod2 - Moni: 3
2015/09/21 14:39:31 prod2 - Dia: 3
2015/09/21 14:39:31 prod2 - Dia: 4
2015/09/21 14:39:31 prod2 - Moni: 4

 */
func TestRun4(t *testing.T) {
	examples.Run4()

}