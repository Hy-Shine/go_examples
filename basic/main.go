package main

import "fmt"

func main() {
	arg := 4
	switch arg {
	case 1:
		multiArgs()
	case 2:
		pointerHandle()
	case 3:
		iotaX()
	case 4:
		mapRead()
	default:
		panic(fmt.Sprintln("arg mismatch"))
	}
	fmt.Printf("--- PROCESS OVER ---\n")
}
