package SearchingAlgorithms

import (
	"fmt"
)

//LinearSearch searches for an integer.
//Input: which takes in a slice and an integer as input.
//Output: index of the item or -1
func LinearSearch(input []int, toFind int) int {

	// fmt.Printf("Input slice is of length %d and capacity %d \n", len(input), cap(input))

	for each := 0; each < len(input); each++ {
		if input[each] == toFind {

			return each
		}
	}
	return -1
}

//PrintTheValue: Prints the value in the slice at an index
//Input: arrays
//Output: index
func PrintTheValue(arrayToRef []int, index int) {
	if index == -1 {
		fmt.Println("\n couldnt find the value")
		return
	}
	fmt.Printf("The value %d is at index %d\n", arrayToRef[index], index)
}

// func main() {
// 	fmt.Println("Calling linear function and sending an array with one item in array")
// 	// inputString := make([]int)
// 	inputString := []int{6, 5, 4, 3, 2, 1}
// 	// result := linearSearch(inputString, 3)
// 	// printTheValue(inputString, result)
// 	// result = linearSearch(make([]int, 0), 3)
// 	// printTheValue(make([]int, 0), result)
// }
