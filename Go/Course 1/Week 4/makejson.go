package main

import (
	"encoding/json"
	"fmt"
)

/* Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys "name" and "address", respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
Submit your source code for the program, "makejson.go". */

func main() {
	var name string
	var address string
	fmt.Print("Enter name:")
	fmt.Scan(&name)
	fmt.Print("Enter address:")
	fmt.Scanln(&address)
	customMap := make(map[string]string)
	customMap["name"] = name
	customMap["address"] = address
	json, err := json.Marshal(customMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Json: %s", string(json))
}
