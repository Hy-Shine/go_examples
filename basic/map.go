package basic

import "fmt"

// map
// map[key_type]values_type

// init map
var listMap = map[int]string{1: "1", 2: "2", 3: "3"}

// equal
// var listMap map[int]string
// listMap = make(map[int]string)
// or if the map in function, 
// listMap := make(map[int]string)

for k, v := range listMap {
	fmt.Println(k, v)
}

if v,ok := listMap[key];ok{
	fmt.Println(v)
}