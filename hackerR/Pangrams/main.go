package main

import "fmt"

// Pangrams
func isPangram(phrase string) string {
	alphabet := [26]bool{}
	for _, char := range phrase {
		if char >= 'a' && char <= 'z' {
			alphabet[char-'a'] = true
		}
	}

	for _, value := range alphabet {
		if !value {
			return "not pangram"
		}
	}

	return "pangram"
}

func main() {

	phrase := "the quick brown fox jumps over the lazy dog"
	fmt.Println(isPangram(phrase))

	phrase = "the quick brown fox jumps over the lazy cat"
	fmt.Println(isPangram(phrase))

}
