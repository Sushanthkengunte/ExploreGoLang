package SearchingAlgorithms

import (
	"fmt"
	"math"
)

//LinearSearch searches for an integer.
//Input: which takes in a slice and an integer as input.
//Output: index of the item or -1. Time Complexity: O(n) Space Complexity: O(n)
func LinearSearch(input []int, toFind int) int {

	// fmt.Printf("Input slice is of length %d and capacity %d \n", len(input), cap(input))

	for each := 0; each < len(input); each++ {
		if input[each] == toFind {

			return each
		}
	}
	return -1
}

//BinarySearch searches for an integer in a sorted array.
//Input: which takes in a slice and an integer as input.
//Output: index of the item or -1. Time complexity: O(nlogn) Space complexity: O(n)
func BinarySearch(input []int, toFind int) int {

	return binaryHelper(0, len(input), toFind, input)
}

func binaryHelper(low int, high int, toFind int, input []int) int {
	if len(input) < 0 {
		return -1
	}
	for low <= high {
		mid := low + (high-low)/2
		if input[mid] == toFind {
			return mid
		} else if input[mid] <= toFind {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

//JumpSearch searches for an integer in a sorted array.
//Input: which takes in a slice and an integer as input.
//Output: index of the item or -1. Time complexity:  Space complexity:
func JumpSearch(input []int, toFind int) int {

	lengthOfTheElement := float64(len(input))

	jump := int(math.Sqrt(lengthOfTheElement))

	var leftPointer int
	for input[int(math.Min(float64(jump), lengthOfTheElement))-1] < toFind {
		//update the left pointer so it points to the next step
		leftPointer = jump

		//update the step to go to the next window
		jump = jump * 2

		if leftPointer > int(lengthOfTheElement) {
			return -1
		}
	}

	for index := leftPointer; index <= int(math.Min(float64((leftPointer+1)*jump), lengthOfTheElement))-1; index++ {
		if input[index] == toFind {
			return index
		}
	}

	return -1
}

//PrintTheValue : Prints the value in the slice at an index
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
