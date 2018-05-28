package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 200; i++ {
		//fmt.Println(42)
		fmt.Printf("%d \t %b \t %#X \t %q \n", i, i, i,i)

	}
}
