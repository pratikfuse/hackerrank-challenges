package main

import (
	"fmt"
	"math"
)

func divide(dividend, divisor int) (int, int) {
	// check if divisor or dividend is negative

	isNegative := (divisor < 0 && dividend > 0) || (divisor > 0 && dividend < 0)
	quotient := 0

	divisor = int(math.Abs(float64(divisor)))
	dividend = int(math.Abs(float64(dividend)))

	for dividend >= divisor {
		dividend = dividend - divisor

		quotient++
	}

	if isNegative {
		quotient = -quotient
	}

	return quotient, dividend
}

func main() {
	quotient, remainder := divide(21, 2)
	fmt.Println(quotient, remainder)
}
