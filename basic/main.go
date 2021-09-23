package main

import "fmt"

func main() {
	arg := 3
	switch arg {
	case 1:
		multiArgs()
	case 2:
		pointerHandle()
	case 3:
		iotaX()
	default:
		panic(fmt.Sprintln("arg mismatch"))
	}
	fmt.Printf("--- PROCESS OVER ---\n")
}
