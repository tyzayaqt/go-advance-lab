package main

import (
	"reflect"
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

		add(50)
		sub(30)

		if got := get(); got != 120 {
			t.Errorf("Accumulator final value = %d, want 120", got)
		}
	})
}

func TestApply(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		operation func(int) int
		want      []int
	}{
		{"Square", []int{1, 2, 3}, func(x int) int { return x * x }, []int{1, 4, 9}},
		{"Double", []int{10, 20}, func(x int) int { return x * 2 }, []int{20, 40}},
		{"Negate", []int{-1, 5}, func(x int) int { return -x }, []int{1, -5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Apply(tt.nums, tt.operation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "even numbers",
			nums:      []int{1, 2, 3, 4, 5, 6},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      []int{2, 4, 6},
		},
		{
			name:      "numbers greater than 5",
			nums:      []int{1, 5, 10, 15},
			predicate: func(x int) bool { return x > 5 },
			want:      []int{10, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.nums, tt.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		initial   int
		operation func(int, int) int
		want      int
	}{
		{"Sum", []int{1, 2, 3, 4}, 0, func(acc, curr int) int { return acc + curr }, 10},
		{"Product", []int{1, 2, 3, 4}, 1, func(acc, curr int) int { return acc * curr }, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.nums, tt.initial, tt.operation); got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	addTwo := func(x int) int { return x + 2 }
	double := func(x int) int { return x * 2 }

	t.Run("Double then add two", func(t *testing.T) {
		f := Compose(addTwo, double)
		if got := f(5); got != 12 {
			t.Errorf("Compose(5) = %d, want 12", got)
		}
	})
}

func TestSwapPointers(t *testing.T) {
	// Correct: Create actual variables on the stack
	x, y := 10, 20

	// Correct: Pass the addresses of those variables
	SwapPointers(&x, &y)

	if x != 20 || y != 10 {
		t.Errorf("Swap failed: x=%d, y=%d", x, y)
	}
}

func TestSwapValues(t *testing.T) {
	tests := []struct {
		name         string
		a, b         int
		wantA, wantB int
	}{
		{name: "swap 5 and 10", a: 5, b: 10, wantA: 10, wantB: 5},
		{name: "swap 0 and -1", a: 0, b: -1, wantA: -1, wantB: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := SwapValues(tt.a, tt.b)
			if gotA != tt.wantA || gotB != tt.wantB {
				t.Errorf("SwapValues() = %v, %v; want %v, %v", gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}
