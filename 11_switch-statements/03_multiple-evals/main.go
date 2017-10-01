package main

import "fmt"

func main()  {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("wassup Tim, or err, Jenny")
	case "Marcus", "Mehdi":
		fmt.Println("both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushant")



	}

}