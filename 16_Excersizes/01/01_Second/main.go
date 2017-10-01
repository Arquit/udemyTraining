package main

import "fmt"

func calculate(variable int) (int, bool) {
	return variable / 2, variable%2 == 0 }

func main() {
	v, potato := calculate(42)
	fmt.Println(v, potato)
}