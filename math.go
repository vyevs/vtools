package vtools

import "golang.org/x/exp/constraints"

// GCD returns the greatest common divisor of a and b.
func GCD[T constraints.Integer](a, b T) T {
	if a < b {
		a, b = b, a
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// GCDAll returns the greatest common divisor of all the integers provided.
// It panics if len(nums) == 0.
func GCDAll[T constraints.Integer](nums ...T) T {
	if len(nums) == 0 {
		panic("0 items provided to GCDAll")
	}

	gcd := nums[0]
	for _, v := range nums[1:] {
		gcd = GCD(v, gcd)
	}
	return gcd
}

// LCM returns the least common multiple of a and b.
// LCM(0, 0) = 0.
func LCM[T constraints.Integer](a, b T) T {
	if a == 0 && b == 0 {
		return 0
	}

	// Order of multiplication and division is important so that the multiplication doesn't overflow the integer type.
	return a * (b / GCD(a, b))
}

// LCMAll returns the least common multiple of the provided integers.
// It panics if len(nums) == 0
func LCMAll[T constraints.Integer](nums ...T) T {
	if len(nums) == 0 {
		panic("0 items provided to LCMAll")
	}

	lcm := nums[0]
	for _, v := range nums[1:] {
		lcm = LCM(v, lcm)
	}
	return lcm
}
