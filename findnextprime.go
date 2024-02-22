package main

import "fmt"

func checkprime(nb int) bool {
	if nb < 1 {
		return false
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
func nextprime(nb int) int {
	if checkprime(nb) {
		return nb
	} else {
		return nextprime(nb + 1)
	}
}
func main() {
	fmt.Println(nextprime(4))
}
