package intutil

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random generates a random integer between min and max (inclusive)
func Random(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min+1) + min
}

// RandomN generates a random integer between 0 and n (exclusive)
func RandomN(n int) int {
	return rand.Intn(n)
}

// RandomSlice generates a slice of random integers
func RandomSlice(length, min, max int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = Random(min, max)
	}
	return result
}

// Abs returns the absolute value of an integer
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Clamp restricts a value to be within a specified range
func Clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// IsEven checks if a number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// IsOdd checks if a number is odd
func IsOdd(n int) bool {
	return n%2 != 0
}

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Factorial calculates the factorial of a number
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// GCD calculates the greatest common divisor of two numbers
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return Abs(a)
}

// LCM calculates the least common multiple of two numbers
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return Abs(a*b) / GCD(a, b)
}
