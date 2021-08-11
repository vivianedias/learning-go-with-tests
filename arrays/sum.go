package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// Stores the length of the array numbersToSum
	lengthOfNumbers := len(numbersToSum)
	// Creates an array with the size of lengthOfNumbers
	sums := make([]int, lengthOfNumbers)

	// Stores the array inside numbers
	// Allocates the result of Sum(numbers) into the respective sums index
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}