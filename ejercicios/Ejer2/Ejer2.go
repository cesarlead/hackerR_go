package ejer2

import "fmt"

// Dado un slice de enteros, crea un map[int]int que represente cuántas veces aparece cada número.
func CountFrecuencyInts(arr []int) map[int]int {

	var result = make(map[int]int)

	for _, num := range arr {
		result[num]++
	}

	fmt.Println(result)

	return result

}

// Dado un slice de strings, retorna otro slice sin duplicados, manteniendo el orden original.
func EliminarDuplicados(arr []string) []string {
	var stringsMain = make(map[string]bool)
	var result = []string{}

	for _, item := range arr {
		if !stringsMain[item] {
			result = append(result, item)
			stringsMain[item] = true
		}
	}

	return result

}

func AgruparXLongitud(arr []string) map[int][]string {

	var result = make(map[int][]string)
	var long int

	for _, item := range arr {
		long = len(item)
		result[long] = append(result[long], item)
	}

	return result
}

// Dado un slice de palabras, encuentra la que más veces se repite.
func MaxCountWord(arr []string) string {
	var frecuenciaDePalabras = make(map[string]int)
	var palabraQueMasSeRepite string
	var contador int

	for _, palabra := range arr {
		frecuenciaDePalabras[palabra]++
		if frecuenciaDePalabras[palabra] > contador {
			contador = frecuenciaDePalabras[palabra]
			palabraQueMasSeRepite = palabra
		}
	}

	return palabraQueMasSeRepite
}
