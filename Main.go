package main

import (
	"ExploreGolang/Controller"
	"ExploreGolang/InputUtility"
	"fmt"
	"strconv"
)

func getArrayLikeElementFromUser() []int {
	var numberOfElements int

	//check if there is 0 is entered
	for numberOfElements <= 0 {
		numberOfElements = InputUtility.GetIntegerInputFromUser("Enter a positive number of elements other than 0 ")
	}
	sliceOfElement := make([]int, numberOfElements)
	for index := 0; index < numberOfElements; index++ {

		message := "Enter the " + strconv.Itoa(index+1) + " element"
		value := InputUtility.GetIntegerInputFromUser(message)

		sliceOfElement[index] = value
	}

	return sliceOfElement
}

func main() {
	fmt.Println("Entry of the program ")

	inputString := getArrayLikeElementFromUser()

	// inputString := []int{6, 5, 4, 3, 2, 1}
	// result := SearchingAlgorithms.LinearSearch(inputString, 1)
	// SearchingAlgorithms.PrintTheValue(inputString, result)

	choiceOfTheUser := InputUtility.GetStringInputFromUser("Enter what kind of program to run. \"searching\" or \"sorting\" ")

	switch choiceOfTheUser {
	case "searching":
		fmt.Println("Searching Implementation")
		integerToSearch := InputUtility.GetIntegerInputFromUser("Enter integer to search")
		// SearchingController.HandleSearch(integerToSearch, inputString)
		Controller.HandleSearch(integerToSearch, inputString)
		break
	case "sorting":
		fmt.Println("Sorting Implementation")
		Controller.HandleSort(inputString)
		break
	default:
		fmt.Println("Select correct choice")
	}

}
