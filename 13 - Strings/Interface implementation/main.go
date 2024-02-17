package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	Sequence string
}

func NewBattery(sequence string) *Battery {
	return &Battery{
		Sequence: sequence,
	}
}

func (battery *Battery) String() string {
	var zeros, ones int
	for _, el := range battery.Sequence {
		switch el {
		case '0':
			zeros++
		case '1':
			ones++
		}
	}

	return "[" +
		strings.Repeat(" ", zeros) +
		strings.Repeat("X", ones) +
		"]"
}

func main() {
	var str string

	fmt.Scan(&str)

	batteryForTest := NewBattery(str)

	fmt.Println(batteryForTest)
}
