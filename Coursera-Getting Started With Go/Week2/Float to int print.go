/*Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.
*/

package main
import "fmt"

func main(){
	var fnum float32
	fmt.Printf("Enter floating number: ")
	fmt.Scan(&fnum) //takes input
	fmt.Print(fnum) //print float numb
	
	fmt.Print("\nThe truncated value is : ",int(fnum)) //print the ans


}
