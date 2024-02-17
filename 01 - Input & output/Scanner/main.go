package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var sum int
	for scanner.Scan() {
		str := scanner.Text()
		num, _ := strconv.Atoi(str)

		sum += num
	}

	_, err := io.WriteString(os.Stdout, strconv.Itoa(sum))
	if err != nil {
		panic(err)
	}
}
