package intutil

import "testing"

func TestRandom(t *testing.T) {
	// Test that Random generates values within range
	for i := 0; i < 100; i++ {
		result := Random(1, 10)
		if result < 1 || result > 10 {
			t.Errorf("Random(1, 10) = %d; want value between 1 and 10", result)
		}
	}

	// Test reversed range
	for i := 0; i < 100; i++ {
		result := Random(10, 1)
		if result < 1 || result > 10 {
			t.Errorf("Random(10, 1) = %d; want value between 1 and 10", result)
		}
	}
}

func TestRandomN(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := RandomN(10)
		if result < 0 || result >= 10 {
			t.Errorf("RandomN(10) = %d; want value between 0 and 9", result)
		}
	}
}

func TestRandomSlice(t *testing.T) {
	result := RandomSlice(5, 1, 10)
	if len(result) != 5 {
		t.Errorf("RandomSlice(5, 1, 10) length = %d; want 5", len(result))
	}

	for _, val := range result {
		if val < 1 || val > 10 {
			t.Errorf("RandomSlice value %d out of range [1, 10]", val)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"positive", 5, 5},
		{"negative", -5, 5},
		{"zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Abs(tt.input)
			if result != tt.expected {
				t.Errorf("Abs(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"a greater", 5, 3, 5},
		{"b greater", 3, 5, 5},
		{"equal", 5, 5, 5},
		{"negative", -5, -3, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Max(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"a smaller", 3, 5, 3},
		{"b smaller", 5, 3, 3},
		{"equal", 5, 5, 5},
		{"negative", -5, -3, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Min(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Min(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		name           string
		value, min, max int
		expected       int
	}{
		{"within range", 5, 1, 10, 5},
		{"below min", -5, 1, 10, 1},
		{"above max", 15, 1, 10, 10},
		{"at min", 1, 1, 10, 1},
		{"at max", 10, 1, 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Clamp(tt.value, tt.min, tt.max)
			if result != tt.expected {
				t.Errorf("Clamp(%d, %d, %d) = %d; want %d", tt.value, tt.min, tt.max, result, tt.expected)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"even positive", 4, true},
		{"odd positive", 5, false},
		{"zero", 0, true},
		{"even negative", -4, true},
		{"odd negative", -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.input)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsOdd(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"even positive", 4, false},
		{"odd positive", 5, true},
		{"zero", 0, false},
		{"even negative", -4, false},
		{"odd negative", -5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsOdd(tt.input)
			if result != tt.expected {
				t.Errorf("IsOdd(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"2 is prime", 2, true},
		{"3 is prime", 3, true},
		{"4 not prime", 4, false},
		{"5 is prime", 5, true},
		{"17 is prime", 17, true},
		{"20 not prime", 20, false},
		{"1 not prime", 1, false},
		{"0 not prime", 0, false},
		{"negative not prime", -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPrime(tt.input)
			if result != tt.expected {
				t.Errorf("IsPrime(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"0!", 0, 1},
		{"1!", 1, 1},
		{"5!", 5, 120},
		{"7!", 7, 5040},
		{"negative", -5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Factorial(tt.input)
			if result != tt.expected {
				t.Errorf("Factorial(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"48 and 18", 48, 18, 6},
		{"100 and 50", 100, 50, 50},
		{"17 and 13", 17, 13, 1},
		{"with negative", -48, 18, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GCD(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("GCD(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"12 and 18", 12, 18, 36},
		{"4 and 6", 4, 6, 12},
		{"7 and 5", 7, 5, 35},
		{"with zero", 0, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LCM(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("LCM(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
