package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Id int `json:"global_id"`
}

func main() {
	bytes, err := os.ReadFile("data-20190514T0100.json")
	if err != nil {
		panic(err)
	}

	var tasks []Task

	err = json.Unmarshal(bytes, &tasks)
	if err != nil {
		panic(err)
	}

	var sum int
	for _, task := range tasks {
		sum += task.Id
	}

	fmt.Println(sum)
}
