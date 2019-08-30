package Controller

import (
	"fmt"
	"leetcode/InputUtility"
	"leetcode/SearchingAlgorithms"
)

//HandleSearch : Selecting the search algorithm
//Input: Integer that has to be searched and the slice where the integer has to be searched
func HandleSearch(integrarToSearch int, searchElement []int) {
	name := InputUtility.GetStringInputFromUser("\t Enter the searching algorithm to search: linear")
	switch name {
	case "linear":
		fmt.Println("\t Linear Search Implementation")
		result := SearchingAlgorithms.LinearSearch(searchElement, integrarToSearch)
		SearchingAlgorithms.PrintTheValue(searchElement, result)
		break
	default:
		break

	}
}
