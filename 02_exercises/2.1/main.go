package main

import (
	"fmt"

	"github.com/franlopezm/golang_first_steps/02_exercises/2.1/tempconv"
)

func main() {
	fmt.Printf("cold %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("%v -> %v\n", tempconv.Fahrenheit(100), tempconv.FToK(100))
	fmt.Printf("%v -> %v\n", tempconv.Kelvin(310), tempconv.KToF(310))
	fmt.Printf("%v -> %v\n", tempconv.Celsius(100), tempconv.CToF(100))
}
