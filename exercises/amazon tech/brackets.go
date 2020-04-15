package main

import (
	"fmt"
)

var m map[string]string

	
	
func main() {

	m = make(map[string]string)
	m["("] = ")"
	m["["] = "]"
	m["{"] = "}"
	
	str := "{}()[]"
	fmt.Println(isValidString(str))
	
}

func isValidString (str string) bool {
	var auxArr []string
	
	
	for i := 0; i < len(str); i++ {
		if str[i] == 40 || str[i] == 91 || str[i] == 123 { //40 == ( | 91 == [ | 123 == {
			auxArr = append(auxArr , string(str[i]))
		} else {
			if m[string(auxArr[len(auxArr)-1])] == string(str[i]) {
				auxArr = remove(auxArr)
			} else {
				fmt.Println("Invalid string")
				return false
			}
			
		}
	}
	
	if len(auxArr) == 0 {
		fmt.Println("Valid string")
		return true
	} else {
		fmt.Println("Invalid string")
		return false
	}
}

func remove (a []string) []string {
	i := len(a)-1 //remove last element
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = ""     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	return a
}