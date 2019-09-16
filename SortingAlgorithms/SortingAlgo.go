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

//SelectionSort : sorts the slice
func SelectionSort(inputSlice []int) []int {
	if len(inputSlice) <= 1 {
		return inputSlice
	}
	var indexOfTheMinElement int
	sortThisArray := inputSlice[:]
	for iterationNumber := 0; iterationNumber < len(sortThisArray)-1; iterationNumber++ {
		indexOfTheMinElement = iterationNumber
		for minElementIndexSearch := iterationNumber + 1; minElementIndexSearch < len(sortThisArray); minElementIndexSearch++ {
			if sortThisArray[minElementIndexSearch] < sortThisArray[indexOfTheMinElement] {
				indexOfTheMinElement = minElementIndexSearch
			}
		}
		swapTheMinElement := sortThisArray[indexOfTheMinElement]
		sortThisArray[indexOfTheMinElement] = sortThisArray[iterationNumber]
		sortThisArray[iterationNumber] = swapTheMinElement
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
