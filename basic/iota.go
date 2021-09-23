package main

import "fmt"

const (
	A int = iota
	B     // = iota
	C     // = iota
)

const (
	a int = iota
	b
	c = 10
	d // = 10
	e // = 10
)

const (
	_ int = iota
	f     // = iota
	g int = 12
	h     = iota + 20
	i     // = iota + 20
)

func iotaX() {
	fmt.Printf("A:%d B:%d C:%d\n", A, B, C)                      //A:0 B:1 C:2
	fmt.Printf("a: %d b: %d c: %d d: %d e: %d\n", a, b, c, d, e) // a:0 b:1 c:10 d:10 e:10
	fmt.Printf("f: %d g: %d h: %d i: %d\n", f, g, h, i)          // f:1 g:12 h:23 i:24

}
