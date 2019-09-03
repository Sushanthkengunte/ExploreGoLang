package Controller

import (
	"fmt"
	"ExploreGolang/InputUtility"
	"ExploreGolang/SearchingAlgorithms"
)

//HandleSearch : Selecting the search algorithm
//Input: Integer that has to be searched and the slice where the integer has to be searched
func HandleSearch(integrarToSearch int, searchElement []int) {
	name := InputUtility.GetStringInputFromUser("\t Enter the searching algorithm to search: linear, binary, jump")
	switch name {
	case "linear":
		fmt.Println("Linear Search Implementation")
		result := SearchingAlgorithms.LinearSearch(searchElement, integrarToSearch)
		SearchingAlgorithms.PrintTheValue(searchElement, result)
		break
	case "binary":
		fmt.Println("Binary Search Implementation")
		result := SearchingAlgorithms.BinarySearch(searchElement, integrarToSearch)
		SearchingAlgorithms.PrintTheValue(searchElement, result)
		break
	case "jump":
		fmt.Println("Jump Search Implementation")
		result := SearchingAlgorithms.JumpSearch(searchElement, integrarToSearch)
		SearchingAlgorithms.PrintTheValue(searchElement, result)
		break
	default:
		break

	}
}
