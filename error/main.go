package main

import "fmt"

func main() {
	error_handle()
	if err := error_handle2(); err != nil {
		fmt.Printf("err: %v", err)
	}
}
