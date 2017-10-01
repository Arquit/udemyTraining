package main

import "fmt"

func main()  {
	switch "Marcus" {
	case "Daniel" :
		fmt.Println("Wassup Daniel")
	case "Medhi" :
		fmt.Println("Wassup Mehdi")
	case "Marcus" :
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Mehdi" :
		fmt.Println("Wassup Mehdi")
	fallthrough

	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
		default:
		fmt.Println("have you no friends?")

	}

}