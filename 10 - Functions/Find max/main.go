package main

import "fmt"

func main() {
	var counter = vote(0, 0, 1)

	fmt.Println(counter)
}

func vote(x int, y int, z int) int {
	var arr = [3]int{x, y, z}

	var zeroCounter = 0
	var oneCounter = 0
	for _, el := range arr {
		switch el {
		case 0:
			zeroCounter++
		case 1:
			oneCounter++
		}
	}

	if zeroCounter > oneCounter {
		return 0
	} else {
		return 1
	}
}
