/*Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case*/


package main

import(
	"fmt"
	"strings"
	)

func main() {
	var str string
	
	fmt.Printf("Input a string: ")
	fmt.Scan(&str)
	
	str=strings.ToLower(str)
	
	switch{
		
		case !strings.Contains(str,"a"):
		fmt.Print("Not found!")
		
 	case !strings.HasPrefix(str,"i"):
		fmt.Print("Not Found!")
	 
	 case !strings.HasSuffix(str,"n"):
	  fmt.Print("Not Found!")
	 
	 default:
 	fmt.Print("Found!")
 	
	}
	
}

