package main

import (
	"fmt"
	"strconv"
)

func main() {
	yearsOld := 32


	fmt.Println("Operators")

	fmt.Println(yearsOld > 30)
	fmt.Println(yearsOld < 30)
	fmt.Println(yearsOld <= 30)
	fmt.Println(yearsOld >= 30)
	fmt.Println(yearsOld == 30)

	fmt.Println()


	fmt.Println(yearsOld < 30 || yearsOld == 32 )
	fmt.Println(yearsOld < 32 || yearsOld == 33 )
	fmt.Println(yearsOld < 40 || yearsOld == 33 )

	fmt.Println()

	fmt.Println(yearsOld < 30 && yearsOld == 32 )
	fmt.Println(yearsOld < 32 && yearsOld == 33 )
	fmt.Println(yearsOld < 40 && yearsOld == 33 )

	fmt.Println()

	fmt.Println(true)
	fmt.Println(!true)

	fmt.Println()

	fmt.Println(yearsOld < 40 && yearsOld == 33 || (yearsOld < 20 && yearsOld == 23))

	yearsOld = 18

	if yearsOld > 18 {
		fmt.Printf("%d is higer than 18 \n", yearsOld)
	}

	boolVal := true

	if boolVal {
		fmt.Println("is true")
	} else {
		fmt.Println("is false")
	}

	intVal2, err2 := strconv.ParseInt("AA1333", 0, 64)
	if err2 == nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("type: %T, value: %d\n", intVal2, intVal2)
	}

	number := 1

	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("undefined number")
	}

	switch  {
	case number < 0:
		fmt.Println("Negativo")
	case number > 0:
		fmt.Println("Positivo")
	default:
		fmt.Println("Zero")
	}
}
