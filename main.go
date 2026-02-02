package main

import (
	"errors"
	"math"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers.")
	}
	if n == 0 {
		return 1, nil
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negatives exponents not supported")
	}
	res := 1
	for i := 0; i < exponent; i++ {
		res *= base

	}
	return res, nil
}

func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(input int) int {
		return input * factor
	}
}

func MakeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	acc := initial
	add = func(n int) {
		acc += n
	}
	subtract = func(n int) {
		acc -= n
	}
	get = func() int {
		return acc
	}
	return add, subtract, get
}
