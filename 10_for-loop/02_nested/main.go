package	main

import "fmt"

func main () {
	for i := 0 ; i < 3; i++ {
		for j := 0 ; j<5 ; j++  {
			for k:=0;k<2;k++ {
				fmt.Println(i, "-", j, "-", k)
			}

		}
	}
}