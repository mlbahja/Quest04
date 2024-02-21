package main

import "fmt"

func iterative(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb < 0 {
		return 0
	}
	result := 1
	if nb > 0 || nb <= 20 {
		for i := 1; i <= nb; i++ {
			result *= i
		}
		//return result
	}
	return result
}

func main() {
	arg := 21
	fmt.Println(iterative(arg))
}
