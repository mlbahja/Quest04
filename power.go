package main

import "fmt"

func power(nb int, power int) int {
// >> nb ^(p) == nb * nb ....(p foix)  
	if power < 0 {
		return 0
	}

	f := 1
	for i := 0; i < power; i++ {
		f *= nb

	}
	return f
}
func main() {
	fmt.Println(power(4, 3))
}
