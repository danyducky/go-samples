package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	for i := Sunday; i <= Saturday; i++ {
		fmt.Println(i)
	}
}
