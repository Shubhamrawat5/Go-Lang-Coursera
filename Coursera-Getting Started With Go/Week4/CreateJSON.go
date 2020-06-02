/*Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.*/
	
package main
import (
	"fmt"
	"encoding/json"
	)

func main() {
	var name string
	var addr string
	fmt.Println("Enter name: ")
	fmt.Scan(&name)
	fmt.Println("Enter address: ")
	fmt.Scan(&addr)
	
	//creating map
	var mp map[string]string
	mp=make(map[string]string)
	
	//adding data to map
	mp["Name"]=name
	mp["Address"]=addr
	
	//creating json
	jsn,_:=json.Marshal(mp)
	
	fmt.Println("JSON data in byte format:- ", jsn)
	
	fmt.Println("JSON data in string format:- ",string(jsn))
	
}
