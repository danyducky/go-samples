package main

import "fmt"

func main() {
	var number int

	fmt.Scan(&number)

	switch {
	case number >= 5 && number <= 20:
		fmt.Println(number, "korov")
	case number%10 == 1:
		fmt.Println(number, "korova")
	case number%10 == 2 || number%10 == 3 || number%10 == 4:
		fmt.Println(number, "korovy")
	default:
		fmt.Println(number, "korov")
	}
}
