package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, world!")
	InitDB()
	defer DB.Close()
	// SeedDB()
	tasks, err := GetTasks()
	if err != nil {
		fmt.Println(err)
	}
    jsonData, err := json.MarshalIndent(tasks, "", " ")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(jsonData))
}
