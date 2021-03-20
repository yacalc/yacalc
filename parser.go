package main

import(
	"fmt"
	"bufio"
	"os"
)

func isOperator(op byte) (bool) {
	operators := [...]byte{'-', '+', '*', '/'}
	for i := 0; i < len(operators); i++ {
		if op == operators[i] {
			return true
		}
	}
	return false
}

// Separates a given string in a slice of strings 
// such that each element is wheter a number or an operator.
func sep(toParse string) ([]string) {
	elements := []string{}
	currentElement := ""
	for i := 0; i <= len(toParse); i++ {
		if i == len(toParse) {
			elements = append(elements, currentElement)
		} else if isOperator(toParse[i])  {
			elements = append(elements, currentElement, string(toParse[i]))
			currentElement = ""
		} else if toParse[i] != ' '{
			currentElement += string(toParse[i])
		}
	}
	return elements
}

//Receives a entire line from the user and convert in to a
//slice of strings. Try sending "12 -14 +6+9 " to it.
//doesn't check for errors.
func main() {
	buffer := bufio.NewReader(os.Stdin);
	for true {
		fmt.Print("> ")
		userInput, _ := buffer.ReadString('\n')
		fmt.Println(sep(userInput[:len(userInput)-1]))
	}
}
