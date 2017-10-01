package main

import "fmt"

func main()   {

	var a int
	var b int
	fmt.Print("Enter the number you want to check ")
	fmt.Scan(&a)

	b = a / 2
	fmt.Println("You variable devided by 2 =" , b)

	if true {
		fmt.Println("The caculated variable is uneven")
	}
	if false {
		fmt.Println("The Calculater variable is even")

	}
}

