package main

import (
	io "fmt"
    sys "os"
)
func main() {
	var input int

	io.Scanf("%d", &input)
    
    sys.Mkdir("Teste", sys.ModePerm)
    
	switch input {
		case 1:
			io.Println("January")
			break	
		case 2:
			io.Println("February")
			break
		case 3:
			io.Println("March")
			break
		case 4:
			io.Println("April")
			break
		case 5:
			io.Println("May")
			break
		case 6:
			io.Println("June")
			break
		case 7:
			io.Println("July")
			break
		case 8:
			io.Println("August")
			break
		case 9:
			io.Println("September")
			break
		case 10:
			io.Println("October")
			break
		case 11:
			io.Println("November")
			break
		case 12:
			io.Println("December")
			break
	}
}
