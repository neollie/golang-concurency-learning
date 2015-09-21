
package examples

/*
import (
	"fmt"
	"testing"
	"reflect"
)


type FibonaciTest struct {
	i   int // index starting from 0
	num int // expected element at i-th position in Fibonaci sequence
}

func TestFibonaciNth(t *testing.T) {
	fmt.Println("Running TestFibonaciNth")
	data := []FibonaciTest{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, // 0,1,1,2,3,5,
	}
	for _, d := range data {
		//fmt.Printf("(%T)%v\n",d, d)
		if num := FibonaciNth(d.i); num != d.num {
			t.Errorf("FibonaciNth(%v) = %v, expected %v", d.i, num, d.num)
		}
	}
}

func TestFibonaciSequence(t *testing.T) {
	fmt.Println("Running TestFibonaciSequence")
	data := []struct {
		expectedSeq []int
	}{
		{[]int{}},
		{[]int{0}},
		{[]int{0, 1}},
		{[]int{0, 1, 1}},
		{[]int{0, 1, 1, 2}},
		{[]int{0,1,1,2,3,5,8,13,21,34,55,89,144}},
	}
	
	for _,d := range data {
		length:= len(d.expectedSeq)
		if seq:= FibonaciSequence(length); reflect.DeepEqual(d.expectedSeq, seq) == false {
			t.Errorf("FibonaciSequence(%v) = %v, expected %v", length, seq, d.expectedSeq)
		}
	}
}

func TestFibonaciNth2(t *testing.T) {
	fmt.Println("Running TestFibonaciNth2")
	data := []FibonaciTest{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, // 0,1,1,2,3,5,
	}
	for _, d := range data {
		//fmt.Printf("(%T)%v\n",d, d)
		if num := FibonaciNth2(d.i); num != d.num {
			t.Errorf("FibonaciNth2(%v) = %v, expected %v", d.i, num, d.num)
		} else {
			t.Logf("FibonaciNth2(%v) = %v, PASS",  d.i, num)
		}
	}
}

*/