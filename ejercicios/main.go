package main

import (
	"fmt"
	"practica/ejercicios/Ejer"
	ejer2 "practica/ejercicios/Ejer2"
)

func main() {

	// Ejercicio 1
	fmt.Println("Ejercicio 1\n")
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	k1 := 3
	fmt.Printf("Input: %v, k: %d -> Output: %v\n", nums1, k1, Ejer.SumSubArrayK(nums1, k1)) // Expected: [6 15 7]

	nums2 := []int{10, 20, 30, 40}
	k2 := 2
	fmt.Printf("Input: %v, k: %d -> Output: %v\n", nums2, k2, Ejer.SumSubArrayK(nums2, k2)) // Expected: [30 70]

	fmt.Println("Ejercicio 2\n")

	// Ejercicio 2
	mainList1 := []string{"apple", "banana", "orange", "grape", "kiwi", "apple"}
	itemsToRemove1 := []string{"banana", "grape", "mango"}
	fmt.Printf("Input mainList: %v, itemsToRemove: %v -> Output: %v\n", mainList1, itemsToRemove1, Ejer.FilterAndRemove1(mainList1, itemsToRemove1)) // Expected: [apple orange kiwi apple]

	mainList2 := []string{"cat", "dog", "bird"}
	itemsToRemove2 := []string{"dog"}
	fmt.Printf("Input mainList: %v, itemsToRemove: %v -> Output: %v\n", mainList2, itemsToRemove2, Ejer.FilterAndRemove2(mainList2, itemsToRemove2)) // Expected: [cat bird]

	nums := []int{1, 1, 3, 3, 5, 6, 7, 8, 8, 8, 99, 8, 345, 678}
	fmt.Println(ejer2.CountFrecuencyInts(nums))

	arrStrings := []string{"go", "java", "go", "python", "java", "go"}
	fmt.Println(ejer2.EliminarDuplicados(arrStrings))

	arrWords := []string{"go", "java", "rust", "python", "elixir"}
	result := ejer2.AgruparXLongitud(arrWords)
	for k, v := range result {
		fmt.Printf("Cantidad de letras: %d, palabras: %v\n", k, v)
	}

	palabras := []string{"go", "java", "go", "go", "rust", "java"}
	masFrecuente := ejer2.MaxCountWord(palabras)
	fmt.Println("Palabra más frecuente:", masFrecuente)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	numbers2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Printf("Contienen los mismos elementos: %v\n", Ejer.AreSlicesEqual(numbers, numbers2))

	idioma := map[string]string{
		"es": "español",
		"en": "inglés",
		"fr": "francés",
	}
	fmt.Println("Mapa original:", idioma)
	invertido := Ejer.InvertMap(idioma)
	fmt.Println("Mapa invertido:", invertido)

	// Ejercicio 3
	array3 := []int{1, 1, 3, 4, 4, 4, 4, 4, 6, 7, 8, 8, 8, 8, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Printf("El numero mas frecuente es(si hay empate el menor): %d\n", Ejer.MostFrequentNumber(array3))

	// Ejercicio 4
	nums4 := []int{1, 1, 1, 2, 2, 3, 3, 4, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10}
	fmt.Printf("Input: %v -> Output: %v\n", nums4, Ejer.SumSingleNumbers(nums4))

	entrada := []string{"roma", "amor", "mora", "mar", "ram", "arco", "roca"}
	fmt.Printf("Input: %v -> Output: %v\n", entrada, Ejer.GroupAnagrams(entrada))

	// Ejercicio 5
	arrNum := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("Output %v\n", Ejer.ReverseSlice(arrNum))

	// Ejercicio 6
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{2, 3, 7, 8, 9, 0}
	fmt.Printf("Los datos de Intersection son: %v\n", Ejer.Intersection(arr1, arr2))

	// Ejercicio 7
	users := []Ejer.User{
		{1, "César", true},
		{2, "Ana", false},
		{3, "Luis", true},
		{4, "Eduardo", false},
		{5, "Karli", true},
		{6, "Enrique", false},
	}

	fmt.Printf("Los Usuarios Activos son:\n ")
	result1 := Ejer.UserIsActive(users)
	for _, u := range result1 {
		fmt.Printf(" -. %s\n", u.Name)
	}

	// Ejercio 8
	result2 := Ejer.GroupUserByStatus(users)
	fmt.Printf("Los Usuarios segun su estatus\n ")
	for k, v := range result2 {
		fmt.Printf("%v --> %v\n", k, v)
	}

	ventas := []Ejer.Venta{
		{"Café", 10},
		{"Té", 15},
		{"Café", 20},
		{"Mate", 5},
	}

	// Ejercicio 9
	fmt.Println(Ejer.ProductMoreSeller(ventas))

}
