package SortingAlgo

import "fmt"

//BubleSort sorts the given slice
func BubleSort(sortThisArrayInput []int) []int {
	if len(sortThisArrayInput) == 0 {
		return nil
	}

	sortThisArray := sortThisArrayInput[:]

	n := len(sortThisArray)
	for iteration := 0; iteration < n; iteration++ {
		// fmt.Println("Iteration number", iteration, " and the array is ", sortThisArray)
		for compare := 0; compare < n-iteration-1; compare++ {
			if sortThisArray[compare] > sortThisArray[compare+1] {
				swap := sortThisArray[compare]
				sortThisArray[compare] = sortThisArray[compare+1]
				sortThisArray[compare+1] = swap
			}

		}
	}
	return sortThisArray
}

//PrintTheValue : Prints the sorted slice
//Input: Slice
func PrintTheValue(sliceToSort []int) {
	fmt.Println("\n======================================================")
	fmt.Println(" \n Final Sorted array is ----> ", sliceToSort, "TADA!!! ")
	fmt.Println("\n======================================================")
}

// func main() {
// 	fmt.Println("Sorting algorithms")
// 	inputString := []int{6, 5, 4, 3, 2, 1}
// 	bubleSort(inputString)

// }
