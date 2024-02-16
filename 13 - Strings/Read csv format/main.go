package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	str := scanner.Text()
	reader := strings.NewReader(str)
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'

	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	replacer := strings.NewReplacer(",", ".", " ", "")

	var nums []float64
	for _, el := range records[0] {
		str := replacer.Replace(el)

		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}

	res := nums[0] / nums[1]

	fmt.Printf("%.4f", res)
}
