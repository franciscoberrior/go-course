package main

import (
	"fmt"
)

func main(){
	var A byte = 'A'
	var a byte = 'a'
	vector := []byte{65,97}

	fmt.Println()
	fmt.Println("Value in ASCII Code: ")
	fmt.Println(A)
	fmt.Println(a)
	fmt.Println()
	fmt.Println(string(a))
	fmt.Println(string(vector))
}