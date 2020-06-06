package main
import ( 
	"fmt" 
	"bufio" 
	"os" 
	"strings" 
	"io"
	)
	
type Name struct {	
		fname string	
		lname string 
		}
		
func main() { 
		fname:="" 
		fmt.Println("Input file name like name.txt :") 
		fmt.Scanln(&fname) 
		var name[] Name //array of object
			
		f, err := os.Open(fname) 
		if err != nil { 
				panic(err) 
				} 
		defer f.Close() 
				
		rd := bufio.NewReader(f) 
		for { 
				line, _, err:= rd.ReadLine() 
				if err != nil || io.EOF == err {
					 break }
				
   tname:=strings.Split(string(line)," ") 
   tpname:= Name { fixLongName(tname[0]), fixLongName(tname[1]), }
   name=append(name,tpname) 
    } 
    
   for i:=0;i<len(name);i++ { 
    		fmt.Println(i+1) 
    		fmt.Println("First Name:"+name[i].fname) 
    		fmt.Println("Last Name:"+name[i].lname) }
 }
 
    		
 // Cut the string if the name>20
 
func fixLongName(buffer string) string { 
 	if len(buffer)>20 { 
 		return string(buffer[0:20]) 
 		} else { return buffer }
}
