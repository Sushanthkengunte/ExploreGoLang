package Controller

import (
	"ExploreGolang/InputUtility"
	SortingAlgo "ExploreGolang/SortingAlgorithms"
	"fmt"
)

//HandleSort : Selecting the search algorithm
//Input: the slice which has to be sorted
func HandleSort(sliceToBeSorted []int) {
	name := InputUtility.GetStringInputFromUser("Enter the sorting algorithm to search: buble, selection, insertion")
	switch name {
	case "buble":
		fmt.Println("Buble Sort Implementation")
		result := SortingAlgo.BubleSort(sliceToBeSorted)
		SortingAlgo.PrintTheValue(result)
		break
	case "selection":
		fmt.Println("Selection sort Implementation")
		result := SortingAlgo.SelectionSort(sliceToBeSorted)
		SortingAlgo.PrintTheValue(result)
		break
	case "insertion":
		fmt.Println("Insertion sort Implementation")
		result := SortingAlgo.InsertionSort(sliceToBeSorted)
		SortingAlgo.PrintTheValue(result)
		break

	default:
		break

	}
}
