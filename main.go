package main

import "fmt"

func CountingSort1(arr []int32) []int32 {
	arrFrecuency := make([]int32, 100)

	for _, value := range arr {
		arrFrecuency[value]++
	}

	return arrFrecuency
}

func main() {

	arr := []int32{1, 1, 3, 2, 1}
	fmt.Println(CountingSort1(arr))

}
