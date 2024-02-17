package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("task.data")
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		for i, str := range record {
			if str == "0" {
				fmt.Println(i)
			}
		}
	}
}
