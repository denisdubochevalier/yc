package yc_test

import (
	"fmt"

	"yc"
)

func ExampleY() {
	fact := yc.Y[int](func(f yc.Function[int]) yc.Function[int] {
		return func(x int) int {
			if x <= 0 {
				return 1
			}
			return x * f(x-1)
		}
	})

	for i := 0; i < 20; i++ {
		fmt.Printf("Factorial of %d -> %d\n", i, fact(i))
	}

	// Output:
	// Factorial of 0 -> 1
	// Factorial of 1 -> 1
	// Factorial of 2 -> 2
	// Factorial of 3 -> 6
	// Factorial of 4 -> 24
	// Factorial of 5 -> 120
	// Factorial of 6 -> 720
	// Factorial of 7 -> 5040
	// Factorial of 8 -> 40320
	// Factorial of 9 -> 362880
	// Factorial of 10 -> 3628800
	// Factorial of 11 -> 39916800
	// Factorial of 12 -> 479001600
	// Factorial of 13 -> 6227020800
	// Factorial of 14 -> 87178291200
	// Factorial of 15 -> 1307674368000
	// Factorial of 16 -> 20922789888000
	// Factorial of 17 -> 355687428096000
	// Factorial of 18 -> 6402373705728000
	// Factorial of 19 -> 121645100408832000
}
