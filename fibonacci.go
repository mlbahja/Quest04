package main

import "fmt"

func fiboni(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {

		return fiboni(i-1) + fiboni(i-2)

	}
}
func main() {
	fmt.Println(fiboni(15))
	fmt.Println(fiboni(1))
	fmt.Println(fiboni(10))
}
