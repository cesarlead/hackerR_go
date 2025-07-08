package main

import "fmt"

// matchingStrings funcion para contar las cadenas que se repiten en un array teniedo otro array como referencia
func matchingStrings(strings []string, queries []string) []int {

	sizeQueries, sizeStrings := len(queries), len(strings)
	var counts []int
	for i := 0; i < sizeQueries; i++ {
		count := 0
		for j := 0; j < sizeStrings; j++ {
			if queries[i] == strings[j] {
				count++
			}
		}
		counts = append(counts, count)
	}
	return counts
}

func main() {

	strings := []string{"ab", "ab", "abc"}
	queries := []string{"ab", "abc", "bc"}

	counts := matchingStrings(strings, queries)
	fmt.Println(counts)

}
