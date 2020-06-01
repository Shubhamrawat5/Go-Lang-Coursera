package main

import "fmt"

func main() {
	//syntax :-
	//var name type
	//int8=8bit,int16=16bit (signed)
	//unint=8bit (unsigned)
	//float32,float64
	//string
	//bool
	//Gets an error if an variable is declare but haven't used in the program
	//Default values are 0, false, ""
	
	x:=50 //short variable declaration (wallrus operator)
	
	var numb int = 10 //explicit
	var numb2 = 20 //implicit

	var check bool =true
	var name string="Shubham"
	
	var x1 int = 100
	var x2 int8 = 120
	//var x3 int8 = 260 - overflow
	var x3 uint8 = 254
	
	fmt.Println(x,numb,numb2,check,name,x1,x2,x3)
}
