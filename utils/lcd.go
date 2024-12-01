package utils

// greatest common divisor
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// least common multiple
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func LeastCommonDenominator(numbers []int) int {
	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}
