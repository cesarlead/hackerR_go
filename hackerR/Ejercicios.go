package hackerr

import (
	"fmt"
	"strconv"
)

// Plus Minus
func PlusMinus(arr []int) string {
	fmt.Println("Ejercicio 1: Plus Minus")

	var positive, negative, zero, length int
	var result string

	if len(arr) > 0 {
		length = len(arr)
	} else {
		return "Error: The array is empty."
	}

	for _, num := range arr {
		if num > 0 {
			positive++
		} else if num < 0 {
			negative++
		} else {
			zero++
		}
	}

	result = fmt.Sprintf("Positive: %.6f\nNegative: %.6f\nZero: %.6f\n", float64(positive)/float64(length), float64(negative)/float64(length), float64(zero)/float64(length))

	return result

}

// Mini-Max Sum
func MiniMaxSum(arr []int) string {

	fmt.Println("Ejercicio 2: Mini-Max Sum")

	var max, totalSum int
	min := arr[0]

	if len(arr) <= 0 {
		return "Error: The array is empty."
	}

	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}

		totalSum += num

	}

	return fmt.Sprintf("%d %d", totalSum-max, totalSum-min)

}

// Time Conversion
func TimeConversion(time string) string {

	fmt.Println("Ejercicio 3: Time Conversion")

	if len(time) != 10 {
		return "Error: Invalid time format."
	}

	h, _ := strconv.Atoi(time[:2])
	m, s, ind := time[3:5], time[6:8], time[8:]

	if ind == "PM" && h != 12 {
		h += 12
	}

	if ind == "AM" && h == 12 {
		h = 0
	}

	return fmt.Sprintf("%02d:%s:%s", h, m, s)
}

// Sparse Arrays
func SparseArrays(arr, queries []string) []int {

	fmt.Println("Ejercicio 4: Sparse Arrays")

	var result []int

	for _, query := range queries {
		count := 0
		for _, item := range arr {
			if item == query {
				count++
			}
		}

		result = append(result, count)
	}

	return result

}

// Lonely Integer
func LonelyInteger(arr []int) int {

	fmt.Println("Ejercicio 5: Lonely Integer")

	var result int
	var seen = make(map[int]bool)

	for _, num := range arr {
		if !seen[num] {
			seen[num] = true
		} else {
			seen[num] = false
		}
	}

	for num, s := range seen {
		if s {
			result = num
		}
	}

	return result
}
