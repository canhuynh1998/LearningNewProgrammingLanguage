package main

import (
	"fmt"
	"math"
)

func NewtonSquareRoot(x float64) float64 {
	result := x
	temp := result

	for {
		temp = result

		result -= (result * result - x) / (2 * result)

		if math.Abs(temp - result) < 0.0000000001{
			break
		}
	}
	return result
}

func main() {
	fmt.Println(NewtonSquareRoot(4))
}