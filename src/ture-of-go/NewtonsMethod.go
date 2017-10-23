package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	min := 0.001
	for {
		next := z - (z * z - x) / 2 * z
		if math.Abs(z - next) < min {
			break
		}
		z = next
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))       //1.4137132974342435 (math.Sqrtの結果に近くなった)
	fmt.Println(math.Sqrt(2))  //1.4142135623730951
}
