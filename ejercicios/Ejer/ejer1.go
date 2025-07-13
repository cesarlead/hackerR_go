package Ejer

import (
	"fmt"
	"sort"
)

type User struct {
	Id       int
	Name     string
	IsActive bool
}

type Venta struct {
	Producto string
	Cantidad int
}

// Given an array of integers nums and an integer k, return the total number of subarrays whose count is equals to k.
func SumSubArrayK(nums []int, k int) []int {

	if k <= 0 {
		return []int{}
	}

	var result []int
	var currentSum, count int

	for _, num := range nums {
		currentSum += num
		count++

		if count == k {
			result = append(result, currentSum)
			currentSum = 0
			count = 0
		}
	}

	if count > 0 {
		result = append(result, currentSum)
	}

	return result

} // Big O(n)

// FilterAndRemove
func FilterAndRemove1(mainList, itemsToRemove []string) []string {
	var result []string
	for _, word := range mainList {
		var found bool
		for _, item := range itemsToRemove {
			if word == item {
				found = true
			}
		}

		if !found {
			result = append(result, word)
		}
	} // Big O(n)

	return result
}

func FilterAndRemove2(mainList, itemsToRemove []string) []string {
	iRemove := make(map[string]bool)

	for _, item := range itemsToRemove {
		iRemove[item] = true
	}

	var result []string
	for _, word := range mainList {
		if !iRemove[word] {
			result = append(result, word)
		}
	} // Big O(n)

	return result
}

// ProcessLogFile
//func ProcessLogFile() {}

// Verificar si dos slices tienen los mismos elementos (sin importar el orden)
func AreSlicesEqual(arr1, arr2 []int) bool {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	for _, num := range arr1 {
		set1[num] = true
	}

	for _, num := range arr2 {
		set2[num] = true
	}

	if len(set1) != len(set2) {
		return false
	}

	for num := range set1 {
		if !set2[num] {
			return false
		}
	}

	return true
}

// Dado un map[string]string, crea un nuevo map[string]string donde las claves y valores estén invertidos.
func InvertMap(m map[string]string) map[string]string {
	var mInvert = make(map[string]string)

	for key, value := range m {
		mInvert[value] = key
	}

	return mInvert
}

// Dado un slice de enteros, retorna el número que aparece más veces. En caso de empate, retorna el menor de ellos. arr := []int{1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 5, 6, 7, 8, 9, 10} esperado := 5
func MostFrequentNumber(arr []int) int {
	var numFrequency = make(map[int]int)
	var maxFrequency, result int

	for _, num := range arr {
		if numFrequency[num] > maxFrequency || numFrequency[num] == maxFrequency && num < result {
			maxFrequency = numFrequency[num]
			result = num
		}
		numFrequency[num]++
	}

	return result

}

// Dado un slice de enteros, suma solo los números que aparecen una sola vez.
func SumSingleNumbers(arr []int) int {
	var result int
	var numFrequency = make(map[int]int)

	for _, num := range arr {
		numFrequency[num]++
	}

	for num, f := range numFrequency {
		if f == 1 {
			result += num
		}
	}

	return result

}

// Dado un slice de palabras, agrúpalas por anagramas. Usa un map[string][]string donde la clave sea la palabra ordenada alfabéticamente.
func GroupAnagrams(arr []string) map[string][]string {
	var result = make(map[string][]string)

	for _, w := range arr {
		sortedWord := []rune(w)
		sort.Slice(sortedWord, func(i, j int) bool {
			return sortedWord[i] < sortedWord[j]
		})

		key := string(sortedWord)
		result[key] = append(result[key], w)
	}

	return result

}

// Revertir los elementos de un slice en el mismo slice.
func ReverseSlice(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

// Dado dos slices de enteros, encuentra los elementos que aparecen en ambos (sin repetir).
func Intersection(arr1, arr2 []int) []int {
	var result []int
	var setDatos = make(map[int]bool)

	for _, num := range arr1 {
		setDatos[num] = true
	}

	for _, num := range arr2 {
		if setDatos[num] {
			result = append(result, num)
		}
	}

	return result

}

// Dado un slice de User, retorna solo los que están activos (IsActive == true).
func UserIsActive(u []User) []User {
	var result []User

	for _, user := range u {
		if user.IsActive {
			result = append(result, user)
		}
	}

	return result
}

// Agrupa los usuarios en un map[bool][]User según su estado activo.
func GroupUserByStatus(u []User) map[bool][]User {

	var result = make(map[bool][]User)

	for _, user := range u {
		result[user.IsActive] = append(result[user.IsActive], user)
	}

	return result
}

// Dado un slice de Venta{Producto string, Cantidad int}, determina cuál producto tiene la mayor cantidad total vendida.
func ProductMoreSeller(v []Venta) string {
	var result string
	var maxS int
	var totales = make(map[string]int)

	for _, venta := range v {
		totales[venta.Producto] += venta.Cantidad
	}

	for product, total := range totales {
		if total > maxS {
			maxS = total
			result = product
		}
	}

	return fmt.Sprintf("El producto mas vendido es: %s, con la cantidad de unidades: %d", result, maxS)
}
