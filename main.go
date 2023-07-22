package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	InitDB()
	defer DB.Close()
	tasks, err := GetTasks()
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
