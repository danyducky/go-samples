package main

import "fmt"

func main() {
	var seconds int

	fmt.Scan(&seconds)

	hours := seconds / 3600
	rest := seconds % 3600
	minutes := rest / 60

	fmt.Println("It is", hours, "hours", minutes, "minutes.")
}
