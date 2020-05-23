package problems1to50

/*Multiples of 3 and 5
Description: https://projecteuler.net/problem=1
*/
import (
	"fmt"
)

func Problem1() {
	i := 1
	result := 0

	for i <= 1000 {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
		i++
	}
	fmt.Println(result)
}
