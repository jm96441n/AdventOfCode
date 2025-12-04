package utils

import "math"

func AbsValue(i int) int {
	if i > 0 {
		return i
	}
	return i * -1
}

// DivMod returns the quotient and remainder of integer division a / b
func DivMod(a, b int) (int, int) {
	return DivFloor(a, b), Mod(a, b)
}

// Mod returns the mathematical modulus of a and b
func Mod(a, b int) int {
	d := DivFloor(a, b)
	return a - (b * d)
}

// DivFloor returns the mathematical floor of a / b
func DivFloor(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}

// IntPow computes base^exp for integers using exponentiation by squaring.
// Assumes exp >= 0
func IntPow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return base
	}
	v := IntPow(base, exp/2)
	if exp%2 == 0 {
		return v * v
	}

	return base * v * v
}
