// https://go.dev/tour/concurrency/8
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t != nil {
			walker(t.Left)
			ch <- t.Value
			walker(t.Right)
		}
	}
	walker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

func main() {
	// 1. Test the Walk function
	// Expected: 1, 2, 3, ... 10
	t, ch := tree.New(1), make(chan int)
	go Walk(t, ch)
	for {
		v, ok := <-ch

		if !ok {
			break
		}

		fmt.Println(v)
	}

	// 2. Test the Same function
	// Expected: true
	fmt.Println(Same(tree.New(1), tree.New(1)))
	// Expected: false
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
