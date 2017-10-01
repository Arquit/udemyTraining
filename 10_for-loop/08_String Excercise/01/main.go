package main

import "fmt"

func main() {
	for i := 50; i <= 155; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}

	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))

	food := "b"
	fmt.Println(food)
	fmt.Printf("%T \n", food)


	}

