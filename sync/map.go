package main

import (
	"fmt"
	"sync"
)

var myMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
}

func printLn(wg *sync.WaitGroup, c int) {
	defer wg.Done()
	fmt.Println(c)
}

func main() {
	var str string
	N := len(myMap)
	var wg sync.WaitGroup
	wg.Add(N)
	for k, v := range myMap {
		str = fmt.Sprintf("key: %s, value: %d", k, v)
		fmt.Println(str)
		go printLn(&wg, v)
	}
	wg.Wait()
	fmt.Printf("This's processing is over!")
}
