package main

import (
	"fmt"

	"github.com/vyevs/vtools"
)

func main() {
	exampleCycle()
}

func exampleCycle() {
	fmt.Println("Here are some example usages of vtools.Cycle:")
	{
		toCycle := []int{1, 2, 3}
		const nIters = 10
		fmt.Printf("we will cycle over %v, printing each value, and stopping after %d iterations\n", toCycle, nIters)

		var i int
		for v := range vtools.Cycle(toCycle) {
			fmt.Println(v)
			i++
			if i == nIters {
				break
			}
		}
		fmt.Println()
	}
	{
		toCycle := []func(){
			func() { fmt.Println("foo") },
			func() { fmt.Println("bar") },
			func() { fmt.Println("baz") },
		}
		const nIters = 10
		fmt.Printf("we will cycle over a slice of functions, calling each, and stopping after %d iterations\n", nIters)
		var i int
		for f := range vtools.Cycle(toCycle) {
			f()
			i++
			if i == nIters {
				break
			}
		}

		fmt.Println()
	}
}
