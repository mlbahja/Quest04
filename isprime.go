package main

import "fmt"

func isprime(nb int) bool {
	
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
			//break
		}

	}
	return true

}
func main() {
	fmt.Println(isprime(5))
	fmt.Println(isprime(4))
}
