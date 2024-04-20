package main

import (
	"fmt"
	"unsafe"
	"strconv"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar)

	var myUintVar uint
	myUintVar = 12
	fmt.Println("My variable is: ", myUintVar)

	var myStringVar string
	myStringVar = "My string value"
	fmt.Println("My variable is: ", myStringVar)

	var myBoolVariable bool
	myBoolVariable = true
	fmt.Println("My variable is: ", myBoolVariable)
	fmt.Println("My variable address is: ", &myBoolVariable)

	myIntVar2 := 12
	fmt.Println("My variable is: ", myIntVar2)

	myStringVar2 := "My string value 2"
	fmt.Println("My variable is: ", myStringVar2)

	const myFirstConst = "a12"
	fmt.Println("My const is: ", myFirstConst)

	const myIntConst int = 12
	fmt.Println("My const is: ", myIntConst)

	const myStringConst string = "String const"
	fmt.Println("My const is: ", myStringConst)

	/*
		int8	8 bis		-128 to 127
		int16	16 bits		-2^15 to 2^15 - 1
		int32	32 bits		-2^31 to 2^31 - 1
		int64	64 bits		-2^63 to 2^63 - 1
		int Platform dependent Platform dependent
	*/

	fmt.Println()

	var my8BitsIntVar int8
	fmt.Printf("Int default value: %d\n", my8BitsIntVar)
	fmt.Printf("type: %T, bytes: %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar) * 8)

	var my16BitsIntVar int16
	fmt.Printf("Int default value: %d\n", my16BitsIntVar)
	fmt.Printf("type: %T, bytes: %d\n", my16BitsIntVar, unsafe.Sizeof(my16BitsIntVar) * 8)

	var my32BitsIntVar int32
	fmt.Printf("Int default value: %d\n", my32BitsIntVar)
	fmt.Printf("type: %T, bytes: %d\n", my32BitsIntVar, unsafe.Sizeof(my32BitsIntVar) * 8)


	var my64BitsIntVar int64
	fmt.Printf("Int default value: %d\n", my64BitsIntVar)
	fmt.Printf("type: %T, bytes: %d\n", my64BitsIntVar, unsafe.Sizeof(my64BitsIntVar) * 8)

	var myAnyBitsIntVar int
	fmt.Printf("Int default value: %d\n", myAnyBitsIntVar)
	fmt.Printf("type: %T, bytes: %d\n", myAnyBitsIntVar, unsafe.Sizeof(myAnyBitsIntVar) * 8)

	/*
		uint8	8 bis		0 to 255
		uint16	16 bits		0 to 2^16 - 1
		uint32	32 bits		0 to 2^32 - 1
		uint64	64 bits		0 to 2^64 - 1
		uint Platform dependent Platform dependent
	*/

	fmt.Println()

	var my8BitsUintVar uint8
	fmt.Printf("Int default value: %d\n", my8BitsUintVar)
	fmt.Printf("type: %T, bytes: %d\n", my8BitsUintVar, unsafe.Sizeof(my8BitsUintVar) * 8)

	var my16BitsUintVar uint16
	fmt.Printf("Int default value: %d\n", my16BitsUintVar)
	fmt.Printf("type: %T, bytes: %d\n", my16BitsUintVar, unsafe.Sizeof(my16BitsUintVar) * 8)

	var my32BitsUintVar uint32
	fmt.Printf("Int default value: %d\n", my32BitsUintVar)
	fmt.Printf("type: %T, bytes: %d\n", my32BitsUintVar, unsafe.Sizeof(my32BitsUintVar) * 8)


	var my64BitsUintVar uint64
	fmt.Printf("Int default value: %d\n", my64BitsUintVar)
	fmt.Printf("type: %T, bytes: %d\n", my64BitsUintVar, unsafe.Sizeof(my64BitsUintVar) * 8)

	var myAnyBitsUintVar uint
	fmt.Printf("Int default value: %d\n", myAnyBitsUintVar)
	fmt.Printf("type: %T, bytes: %d\n", myAnyBitsUintVar, unsafe.Sizeof(myAnyBitsUintVar) * 8)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	var myStringVar3 string
	fmt.Printf("Int default value: %s\n", myStringVar3)

	myStringVar5 := `My string value in golang
	with multiple
	line` 

	fmt.Printf("The variable value is: %s\n", myStringVar5)

	{
		/* LO QUE SE HACE EN ESTE SCOPE SOLO VIVE EN EL SCOPE Y SE USA SOLO AHÍ	*/

		fmt.Println()

		floatVar := 33.11
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)

		floatStrVar := fmt.Sprintf("%f", floatVar)
		fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

		intVal1, err := strconv.ParseInt("1333", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal1, intVal1)


		intVal2, err2 := strconv.ParseInt("AA1333", 0, 64)
		if err == nil {
			fmt.Println(err2)
		} else {
			fmt.Printf("type: %T, value: %d\n", intVal2, intVal2)
		}
		
		intVal3, _ := strconv.ParseInt("1333", 0, 64)
		fmt.Printf("type: %T, value: %d\n", intVal3, intVal3)
	}
}
