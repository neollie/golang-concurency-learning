package examples

import (
	"testing"
	"golang.org/x/tour/tree"
)

func TestEquivalentTree(t *testing.T) {
	t1, t2:=tree.New(1), tree.New(1)
	if !(CheckEquivalent(t1,t2)) {
		t.Errorf("%v not equivalent to %v, expected to be equivalent", t1, t2)
	}
}

func TestNonEquivalentTree(t *testing.T) {
	t1, t2:=tree.New(1), tree.New(2)
	if (CheckEquivalent(t1,t2)) {
		t.Errorf("%v is equivalent to %v, expected not to be equivalent", t1, t2)
	}
}

func TestNonEquivalentTree2(t *testing.T) {
	t1, t2:=tree.New(4), tree.New(2)
	if (CheckEquivalent(t1,t2)) {
		t.Errorf("%v is equivalent to %v, expected not to be equivalent", t1, t2)
	}
}