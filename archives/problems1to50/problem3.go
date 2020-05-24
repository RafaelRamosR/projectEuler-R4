package problems1to50

import (
	"fmt"
)

// Problem3 : Largest prime factor
// Description: https://projecteuler.net/problem=3
func Problem3() {
	i := 2
	factor := 600851475143

	// The prime factor of a number cannot be greater than this same number
	for i < factor {
		if factor%i == 0 {
			// The quotient obtained now is the new limit
			factor = factor / i
		} else {
			// The divisor increases if the division is not exact
			i = i + 1
		}
	}
	// The last limit obtained will be the largest prime factor.
	fmt.Println("Largest prime factor:", factor)
}
