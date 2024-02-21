package main

import "fmt"

/*func recursive(nb int) int {
	if nb == 0 {
		return 1
	}
	/*if nb < 0 {
		return 0
	}
	result := 1
	//if !(nb > 0 || nb < 21) {
	//return result * recursive(nb-1)
	//eturn 0
	//}
	return (result * recursive(nb-1))
}*/
func countdown(nb int) {
	if nb > 0 {
		i := 0
		for i < nb {

			fmt.Println(nb)
			nb--

		}
		i++
	}
}
func recursiveer(nb int) int {
	if nb > 0 {
		return nb * recursiveer(nb-1)
	} else {
		return 1
	}
}

func main() {
	//number := 4
	//fmt.Println(recursive(number))
	countdown(4)
	fmt.Println("##########")
	fmt.Print(recursiveer(4))
}
