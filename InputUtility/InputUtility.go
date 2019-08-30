package InputUtility

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//GetStringInputFromUser : gets an integer from user
// Input: what message has to be displayed
//output: a String
func GetStringInputFromUser(whatTheUserHasToEnter string) string {
	fmt.Println(whatTheUserHasToEnter)
	var choiceOfTheUser string
	reader := bufio.NewReader(os.Stdin)
	for choiceOfTheUser == "" {
		choiceOfTheUser, _ = reader.ReadString('\n')
	}

	choiceOfTheUser = strings.Replace(strings.ToLower(choiceOfTheUser), " ", "", -1)
	choiceOfTheUser = choiceOfTheUser[:len(choiceOfTheUser)-1]
	return choiceOfTheUser
}

//GetIntegerInputFromUser : gets an integer from user
// Input: what message has to be displayed
//output: an integer
func GetIntegerInputFromUser(whatTheUserHasToEnter string) int {

	stringInput := GetStringInputFromUser(whatTheUserHasToEnter)
	choiceOfTheUser, err := strconv.Atoi(stringInput)

	if err != nil {
		fmt.Println("Error is there and its message is", err)
	}
	return choiceOfTheUser
}
