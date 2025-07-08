package main

import "fmt"

// Lonely Integer
func lonelyInteger(a []int) int {
	var result int
	for _, value := range a {
		result ^= value
	}
	return result
}

func main() {

	a := []int{1, 2, 3, 4, 3, 2, 1}
	fmt.Println(lonelyInteger(a))

}
