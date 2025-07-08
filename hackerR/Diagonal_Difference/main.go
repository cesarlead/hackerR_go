package main

import "fmt"

// Diagonal_Difference
func diagonalDifference(arr [][]int32) int32 {
	var leftToRight, rightToLeft int32

	for i := 0; i < len(arr); i++ {
		leftToRight += arr[i][i]
		rightToLeft += arr[i][len(arr)-i-1]
	}

	return int32(uint32(int(leftToRight - rightToLeft)))
}

func main() {

	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	fmt.Println(diagonalDifference(arr))

}
