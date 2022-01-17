package main

import "fmt"

// map
// map[key_type]values_type

// init map
// var listMap map[int]string
var listMap map[int]string

// equal
// var listMap map[int]string
// listMap = make(map[int]string)
// or if the map in function,
// listMap := make(map[int]string)

func mapRead() {
	listMap = make(map[int]string)
	listMap[1] = "One"
	listMap[2] = "Two"
	listMap[3] = "Three"
	for k, v := range listMap {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}
	//init a map with ":="
	listMap2 := make(map[int]string)
	listMap2[5] = "Five"
	listMap2[10] = "Ten"
	for key := range listMap {
		fmt.Printf("key: %d, value: %s\n", key, listMap[key])
	}
	//
	key := 5
	if v, ok := listMap2[key]; ok {
		fmt.Println(v)
	}
}
