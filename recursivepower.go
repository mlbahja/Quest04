package main

import "fmt"

func recursivepowert(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 1 {
		return nb
	} else if power == 0 {
		return 1
	} else {
		return nb * recursivepowert(nb, power-1)
	}
}
func main() {
	fmt.Println(recursivepowert(4, 3))
}
