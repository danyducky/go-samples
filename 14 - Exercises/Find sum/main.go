package main

import (
	"fmt"
	"os"
)

func main() {
	value1, value2, operation := readTask()

	var a float64
	if v, ok := value1.(float64); ok {
		a = v
	} else {
		printAndExit(value1)
	}

	var b float64
	if v, ok := value2.(float64); ok {
		b = v
	} else {
		printAndExit(value2)
	}

	switch v := operation.(type) {
	case string:
		switch v {
		case "+":
			fmt.Printf("%.4f", a+b)
		case "-":
			fmt.Printf("%.4f", a-b)
		case "*":
			fmt.Printf("%.4f", a*b)
		case "/":
			fmt.Printf("%.4f", a/b)
		default:
			fmt.Println("неизвестная операция")
		}
	default:
		fmt.Println("неизвестная операция")
	}
}

func readTask() (value1, value2, operation interface{}) {
	return 5.0, 5.6, "/" // can change any param...
}

func printAndExit(value interface{}) {
	fmt.Printf("value=%v: %T", value, value)

	os.Exit(0)
}
