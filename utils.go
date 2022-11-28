package main

func addNumbers(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}
