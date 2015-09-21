package examples

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	_walk(t, ch) // We have to separate traversing to be able walk and close correctly channel
}

func _walk(t *tree.Tree, ch chan int) {
	// defer close(ch) // Warning: we can't call defer here, because channel would be closed after first recursion call
	if t == nil {
		return
	}
	_walk(t.Left, ch)
	ch <- t.Value
	_walk(t.Right, ch)
}

// Enhancement
func Walk2(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil { return }
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}


// Same determines whether the trees
// t1 and t2 contain the same values.
func CheckEquivalent(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2:=  make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	x1, x2, ok1, ok2 := 0, 0, true, true

	for ok1 && ok2 {
		x1, ok1 = <- c1
		x2, ok2 = <- c2

		if ok1 != ok2 || x1 != x2 {
			return false
		}
	}

	return true
}

