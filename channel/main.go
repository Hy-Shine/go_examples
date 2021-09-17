package main

import "fmt"

func gen(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range nums {
			ch <- v * v
		}
		close(ch)
	}()
	return ch
}

func sq(ch <-chan int) <-chan int {
	ch2 := make(chan int)
	go func() {
		for v := range ch {
			ch2 <- v*v - v
		}
		close(ch2)
	}()
	return ch2
}

func main() {
	sq := sq(gen(1, 2, 3, 4, 5, 6, 7, 8))
	// sq := sq(gen)
	for x := range sq {
		fmt.Println(x)
	}
}
