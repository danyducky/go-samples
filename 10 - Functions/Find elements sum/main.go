package main

func main() {

}

func sum(elements ...int) (int, int) {
	var counter = 0
	var sum = 0
	for _, el := range elements {
		sum += el
		counter++
	}
	return counter, sum
}
