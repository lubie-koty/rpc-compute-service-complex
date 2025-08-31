package base

import (
	"testing"
)

type UnaryTestCase struct {
	Name     string
	Input    float64
	Expected float64
}

type BinaryTestCase struct {
	Name     string
	InputA   float64
	InputB   float64
	Expected float64
}

var s ComplexMathService

func TestBaseSqrt(t *testing.T) {
	tests := []UnaryTestCase{
		{
			Name:     "x^1/2 should be x",
			Input:    4,
			Expected: 2,
		},
		{
			Name:     "x^1/2 should be x",
			Input:    9,
			Expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := s.Sqrt(test.Input)
			if result != test.Expected {
				t.Errorf("Expected %v, got %v", test.Expected, result)
			}
		})
	}
}

func TestBaseAbs(t *testing.T) {
	tests := []UnaryTestCase{
		{
			Name:     "0 should be 0",
			Input:    0,
			Expected: 0,
		},
		{
			Name:     "-x should be x",
			Input:    -1,
			Expected: 1,
		},
		{
			Name:     "x should be x",
			Input:    1,
			Expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := s.Abs(test.Input)
			if result != test.Expected {
				t.Errorf("Expected %v, got %v", test.Expected, result)
			}
		})
	}
}

func TestBasePower(t *testing.T) {
	tests := []BinaryTestCase{
		{
			Name:     "x^0 should be 1",
			InputA:   5,
			InputB:   0,
			Expected: 1,
		},
		{
			Name:     "x^y should be x^y",
			InputA:   2,
			InputB:   5,
			Expected: 32,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := s.Power(test.InputA, test.InputB)
			if result != test.Expected {
				t.Errorf("Expected %v, got %v", test.Expected, result)
			}
		})
	}
}

func TestBaseLog(t *testing.T) {
	tests := []BinaryTestCase{
		{
			Name:     "case 1: log value should match",
			InputA:   10,
			InputB:   10,
			Expected: 1,
		},
		{
			Name:     "case 2: log value should match",
			InputA:   32,
			InputB:   2,
			Expected: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := s.Log(test.InputA, test.InputB)
			if result != test.Expected {
				t.Errorf("Expected %v, got %v", test.Expected, result)
			}
		})
	}
}

func TestBaseRound(t *testing.T) {
	tests := []BinaryTestCase{
		{
			Name:     "1.11111 should be 1",
			InputA:   1.11111,
			InputB:   0,
			Expected: 1,
		},
		{
			Name:     "1.11111 should be 1.1",
			InputA:   1.11111,
			InputB:   1,
			Expected: 1.1,
		},
		{
			Name:     "1.11111 should be 1.11",
			InputA:   1.11111,
			InputB:   2,
			Expected: 1.11,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := s.Round(test.InputA, int64(test.InputB))
			if result != test.Expected {
				t.Errorf("Expected %v, got %v", test.Expected, result)
			}
		})
	}
}
