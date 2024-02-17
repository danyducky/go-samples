package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Rating struct {
	Average float64
}

func NewRating(average float64) *Rating {
	return &Rating{
		Average: average,
	}
}

func main() {
	bytes, err := os.ReadFile("text.json")
	if err != nil {
		panic(err)
	}

	var group Group

	err = json.Unmarshal(bytes, &group)
	if err != nil {
		panic(err)
	}

	var ratingCounter int
	var studentsCounter int
	for _, student := range group.Students {
		ratingCounter += len(student.Rating)
		studentsCounter++
	}

	average := float64(ratingCounter) / float64(studentsCounter)
	rating := NewRating(average)

	output, err := json.MarshalIndent(rating, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}
