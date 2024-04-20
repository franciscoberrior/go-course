package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var myIntVar int
	fmt.Println(myIntVar)
	fmt.Printf("Type: %T, bytes: %d, bits:%d \n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar) * 8)


}