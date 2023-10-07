package main

import (
	"fmt"
)

func main() {
	input_number := 6

	for i := 1; i <= input_number; i++ {
		cube := i * i * i
		fmt.Printf("Current Number is : %d and the cube is %d\n", i, cube)
	}
}
