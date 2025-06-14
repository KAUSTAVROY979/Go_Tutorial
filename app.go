package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Hello World\n")
	fmt.Print("Your return would be: ", investment_calculator())
}

func investment_calculator() float64 {
	var investedAmount = 1000
	var returnRate = 5.5
	var years = 10

	return float64(investedAmount) + math.Pow(1+returnRate/100, float64(years))

}
