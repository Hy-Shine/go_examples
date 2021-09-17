package main

import (
	"encoding/json"
	"fmt"
)

type jsonObj2 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func serialize() {
	// go object to json
	goStruct := jsonObj2{
		ID:   100,
		Name: "Tom",
		Age:  10,
	}
	structToJson, err := json.MarshalIndent(goStruct, "", "") //json.Marshal()
	if err != nil {
		panic(fmt.Sprintf("err: %v\n", err))
	}
	fmt.Println(string(structToJson))

	// json data to go object
	jsonData := `{"id":10,"name":"John","age":5}`
	x := []byte(jsonData)
	r := goStruct
	if err = json.Unmarshal(x, &r); err != nil {
		panic(fmt.Sprintf("err: %v\n", err))
	}
	fmt.Printf("id=%d, name=%s, age=%d\n", r.ID, r.Name, r.Age)
}
