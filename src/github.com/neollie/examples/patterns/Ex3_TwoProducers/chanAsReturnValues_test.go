package examples

import (
	"testing"
	"github.com/neollie/examples/patterns/Ex3_TwoProducers"
)

// Test it with #go test -v nopatter_test.go
func TestNoPattern(t *testing.T) {
	Run3();

	/**
	// Expected output
=== RUN   TestNoPattern
2015/09/21 13:29:28 prod2 - Moni: 0
2015/09/21 13:29:28 prod2 - Dia: 0
2015/09/21 13:29:28 prod2 - Moni: 1
2015/09/21 13:29:28 prod2 - Dia: 1
2015/09/21 13:29:28 prod2 - Moni: 2
2015/09/21 13:29:28 prod2 - Dia: 2
2015/09/21 13:29:28 prod2 - Moni: 3
2015/09/21 13:29:28 prod2 - Dia: 3
2015/09/21 13:29:28 prod2 - Moni: 4
2015/09/21 13:29:28 prod2 - Dia: 4

	 */
}