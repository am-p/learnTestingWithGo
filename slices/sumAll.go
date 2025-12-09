package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	fmt.Println(numbersToSum)
 	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] =Sum(numbers)
	}
	return sums
}

func SumAllAppend(numbersToSum ...[]int) []int{
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	
	return sums
}

func main() {
	fmt.Println(SumAll([]int{1,2},[]int{0,9}))
}


