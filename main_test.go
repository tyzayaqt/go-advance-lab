package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{name: "factorial of 0", input: 0, want: 1, wantErr: false},
		{name: "factorial of 5", input: 5, want: 120, wantErr: false},
		{name: "factorial of 1", input: 1, want: 1, wantErr: false},
		{name: "negative input", input: -5, want: 0, wantErr: true},
		{name: "factorial of 3", input: 3, want: 6, wantErr: false},
		{name: "factorial of 10", input: 10, want: 3628800, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{name: "prime 2", input: 2, want: true, wantErr: false},
		{name: "prime 17", input: 17, want: true, wantErr: false},
		{name: "composite 4", input: 4, want: false, wantErr: false},
		{name: "composite 25", input: 25, want: false, wantErr: false},
		{name: "error case 1", input: 1, want: false, wantErr: true},
		{name: "error case negative", input: -5, want: false, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		want     int
		wantErr  bool
	}{
		{"positive power", 2, 3, 8, false},
		{"zero power", 5, 0, 1, false},
		{"zero base", 0, 5, 0, false},
		{"base one", 1, 100, 1, false},
		{"negative base", -2, 2, 4, false},
		{"negative exponent", 2, -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Power(%d, %d) error = %v, wantErr %v", tt.base, tt.exponent, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Power(%d, %d) = %d, want %d", tt.base, tt.exponent, got, tt.want)
			}
		})
	}
}

func TestMakeCounter(t *testing.T) {
	t.Run("Counter independence", func(t *testing.T) {
		c1 := MakeCounter(0)
		c2 := MakeCounter(10)

		if got := c1(); got != 1 {
			t.Errorf("Counter1 first call = %d, want 1", got)
		}
		if got := c2(); got != 11 {
			t.Errorf("Counter2 first call = %d, want 11", got)
		}
		if got := c1(); got != 2 {
			t.Errorf("Counter1 second call = %d, want 2", got)
		}
	})
}

func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name   string
		factor int
		input  int
		want   int
	}{
		{"Double", 2, 5, 10},
		{"Triple", 3, 5, 15},
		{"Multiply by zero", 0, 10, 0},
		{"Negative factor", -2, 4, -8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mult := MakeMultiplier(tt.factor)
			if got := mult(tt.input); got != tt.want {
				t.Errorf("MakeMultiplier(%d)(%d) = %d, want %d", tt.factor, tt.input, got, tt.want)
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	t.Run("Shared state interaction", func(t *testing.T) {
		add, sub, get := MakeAccumulator(100)

		add(50) // 150
		sub(30) // 120

		if got := get(); got != 120 {
			t.Errorf("Accumulator final value = %d, want 120", got)
		}
	})
}
