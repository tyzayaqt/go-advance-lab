package main

import (
	"errors"
	"fmt"
	"math"
	"os"
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

func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = operation(v)
	}
	return result
}

func Filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range nums {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce(nums []int, initial int, operation func(accumulator, current int) int) int {
	acc := initial
	for _, v := range nums {
		acc = operation(acc, v)
	}
	return acc
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func ExploreProcess() {
	fmt.Println("=== Process Information ===")
	//A process ID (PID) is a unique numerical identifier assigned by the OD=S
	// to every active process to track its excution and resources.
	fmt.Printf("Current Process ID: %d\n", os.Getpid())

	fmt.Printf("PArent Process ID: %d\n", os.Getpid())

	data := []int{1, 2, 3, 4, 5}

	fmt.Printf("Memory address of slice: %p\n", &data)

	//The slice header contains metadata (pointer, length, capacity),
	// while the firtst element address is where the actual data resides in memory.
	fmt.Printf("Memoy address of first element: %p\n", &data[0])

	/* Why process validation is important? It prevents processes from accessing or modifyin the memory
	of other processes, ensuring system stability and security. Other processses
	cannot access these memory addresses.*/
	fmt.Println("Note: Other processes cannot access these memory addresses due to process isolation")
}

func main() {
	ExploreProcess()
}
