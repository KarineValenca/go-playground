package main

import (
	"fmt"
)

func main () {
	
	
	str := "{}"
	
	isValid := isValidString(str)
	fmt.Println(isValid)
	
}
func isValidString(str string) bool {
	switch {
		case str[0] == 40 && str[1] == 41:
			fmt.Println("Valid string")
			return true
		case str[0] == 91 && str[1] == 93:
			fmt.Println("Valid string")
			return true
		case str[0] == 123 && str[1] == 125:
			fmt.Println("Valid string")
			return true
		default:
			fmt.Println("Invalid string")
			return false
	}

}