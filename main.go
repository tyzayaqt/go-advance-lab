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

// Question: Will this modify the original variable?
// Answer: No, because Go passes by value, meaning the function
// receives a copy of the integer[cite: 254, 261].
func DoubleValue(x int) {
	x = x * 2
}

// Question: Will this modify the original variable?
// Answer: Yes, because the function uses the memory address to modify
// the original value directly[cite: 262, 263, 264].
func DoublePointer(x *int) {
	*x = *x * 2
}

// Comment: This variable stays on the stack[cite: 265, 268].
func CreateOnStack() int {
	x := 10
	return x
}

// Comment: This variable escapes to the heap[cite: 269, 271].
func CreateOnHeap() *int {
	x := 20
	return &x
}

func SwapValues(a, b int) (int, int) {
	return b, a
}

func SwapPointers(a, b *int) {
	if a == nil || b == nil {
		return
	}
	temp := *a
	*a = *b
	*b = temp
}

func AnalyzeEscape() {
	_ = CreateOnStack()
	_ = CreateOnHeap()
}

/*
ESCAPE ANALYSIS REPORT:
1. Which variables escaped to the heap?
   The variable 'x' in CreateOnHeap escaped.
2. Why did they escape?
   Because a pointer to a local variable was returned; the compiler
   must move it to the heap so it lives longer than the function call.
3. What does "escapes to heap" mean?
   It means the memory is allocated in the dynamic heap area rather
   than the function's local stack frame because its lifetime is
   unpredictable or exceeds the function scope.
*/

func main() {
	ExploreProcess()
}
