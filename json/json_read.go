package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type jsonObj struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func ReadFromFile() {
	var jsonParse jsonObj
	data, err := ioutil.ReadFile("temp.json")
	if err != nil {
		panic(fmt.Sprintf("it's error occored: %v", err))
	}
	if err = json.Unmarshal(data, &jsonParse); err != nil {
		panic(fmt.Sprintln(err))
	}
	fmt.Printf("json file: id is %d, name is %s, gender is %s\n", jsonParse.ID, jsonParse.Name, jsonParse.Gender)
}
