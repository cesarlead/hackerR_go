package main

import "fmt"

// Flipping Bits
func flippingBits(n int64) int64 {
	result := n ^ 0xFFFFFFFF
	return result

}

func main() {

	n := int64(500)
	fmt.Println(flippingBits(n))

}
