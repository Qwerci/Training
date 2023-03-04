package main

import "fmt"

func main() {
	var numbers []int = []int {21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}
	fmt.Println("Unsorted",numbers)
	// bubblesort(numbers)
	bs(numbers)

	fmt.Println("Sorted",numbers)
}

func bubblesort(numbers []int) {
	for i := 0; i < len(numbers); i++{
		if !swap(numbers,i){
			return
		}
	}
}

func swap(numbers []int, prev_passed int) bool{
	var N int = len(numbers)
	var firstindex int = 0
	var secondindex int = 1
	is_swapped := false

	// pair them
	for secondindex < (N - prev_passed){
		firstnumber := numbers[firstindex]
		secondnumber := numbers[secondindex]

		// swap the numbers
		if secondnumber > firstnumber {
			numbers[firstindex] = secondnumber
			numbers[secondindex] = firstnumber
			is_swapped = true
		}

		firstindex ++
		secondindex ++
	


	}

	return is_swapped

	// 
}






// Another implementation
func bs(numbers []int) {
	for i := len(numbers); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if numbers[j] < numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}