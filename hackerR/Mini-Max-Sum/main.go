package main

import (
	"fmt"
	"sort"
)

// Mini-Max Sum
// si todos los elementos del arreglo son iguales, el resultado de la suma minima y maxima es el mismo
func minMaxSum(arr []int) {
	sort.Ints(arr)
	minValue := arr[0]
	maxValue := arr[len(arr)-1]

	var totalSum int
	for _, value := range arr {
		totalSum += value
	}
	fmt.Println(totalSum-maxValue, totalSum-minValue)
}

func main() {
	//arr := []int{1, 3, 5, 7, 9}
	arr := []int{5, 5, 5, 5, 5}
	minMaxSum(arr)

}
