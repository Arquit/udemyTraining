package main

import (
	"fmt"

)

func main()  {
	n := max(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}
func max(sf ...float64) float64 {
		var highest float64

		for _, v:= range sf {
			if v > highest {
				highest = v
			}
		}
		return highest
}
