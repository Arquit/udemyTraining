package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		for i%1 == 0 {}
		//for i%1 == 0&&i%2 == 0 && i%3 == 0 && i%4 == 0 {
			fmt.Println(i)
		}
	}
//}