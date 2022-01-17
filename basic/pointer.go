package main

import "fmt"

type intX int

func pointerSum(x *intX, y intX) {
	*x = *x + y
}

func (p *intX) pointerSum2(x int) {
	*p = *p + intX(x)
}

func pointerHandle() {
	var x, y intX = 10, 20
	pointerSum(&x, y)
	fmt.Printf("the value of x is %d.\n", x)
	// reference a non-pointer
	x.pointerSum2(int(y))
	fmt.Printf("the value of x is %d.\n", x)

	p := new(intX)
	*p = 5
	p.pointerSum2(int(y))
	fmt.Printf("the value of x is %d.\n", *p)
}
