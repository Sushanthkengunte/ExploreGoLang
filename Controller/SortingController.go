package Controller

import (
	"fmt"
	"leetcode/InputUtility"
	SortingAlgo "leetcode/SortingAlgorithms"
)

//HandleSort : Selecting the search algorithm
//Input: the slice which has to be sorted
func HandleSort(sliceToBeSorted []int) {
	name := InputUtility.GetStringInputFromUser("Enter the sorting algorithm to search: buble")
	switch name {
	case "buble":
		fmt.Println("Buble Sort Implementation")
		result := SortingAlgo.BubleSort(sliceToBeSorted)
		SortingAlgo.PrintTheValue(result)
		break

	default:
		break

	}
}
