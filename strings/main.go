package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// funcion para contar letras solamente en una oracion
// func countLetters(sentence string) int {
// 	count := 0
// 	for _, char := range sentence {
// 		if unicode.IsLetter(char) {
// 			count++
// 		}
// 	}
// 	return count
// }

// funcion que recibe un timestring y regresa la hora parseada en nuevayork, londres y tokio
func parseTimeString(timestring string) (string, string, string, error) {
	layout := "2006-01-02 15:04 MST"
	t, err := time.Parse(layout, timestring)
	if err != nil {
		return "", "", "", fmt.Errorf("error parsing time string: %v", err)
	}

	newYorkTime, _ := time.LoadLocation("America/New_York")
	londonTime, _ := time.LoadLocation("Europe/London")
	tokyoTime, _ := time.LoadLocation("Asia/Tokyo")

	return t.In(newYorkTime).Format(layout), t.In(londonTime).Format(layout), t.In(tokyoTime).Format(layout), nil
}

func main() {

	// s := "\x99\x00ab\x50\x00\x23\x50\x29\x9c"

	// cadena := "Hola, Mundo!"
	// word := "Hola"

	// for i := 0; i < len(s); i++ {
	// 	if unicode.IsPrint(rune(s[i])) {
	// 		fmt.Printf("Caracter: %c\n", s[i])
	// 	} else {
	// 		fmt.Printf("Caracter no imprimible: %d\n", s[i])
	// 	}
	// }

	// fmt.Println(strings.ToUpper(cadena))
	// fmt.Println(strings.Contains(cadena, word))
	// fmt.Println(strings.Count(cadena, word))
	// fmt.Println(strings.Split(cadena, " "))

	ahora := time.Now()
	// fmt.Println("Fecha y hora actual: ", ahora)
	// fmt.Println("Año:", ahora.Year())
	// fmt.Println("Mes:", ahora.Month())
	// fmt.Println("Día:", ahora.Day())
	// fmt.Println("Hora:", ahora.Hour())
	// fmt.Println("Minuto:", ahora.Minute())
	// fmt.Println("Segundo:", ahora.Second())
	// fmt.Println("Nanosegundo:", ahora.Nanosecond())
	// fmt.Println("Día de la semana:", ahora.Weekday())
	// fmt.Println("Formato de fecha:", ahora.Format("2006-01-02 15:04:05"))
	// fmt.Println("Formato de fecha corto:", ahora.Format("02/01/2006"))
	fmt.Println("Formato de fecha largo:", ahora.Format("Monday, 02-Jan-06 15:04:05 MST"))
	// fmt.Println("Timestamp:", ahora.Unix())

	// reader := bufio.NewReader(os.Stdin)
	// var sentence string
	// fmt.Println("Enter a sentence: ")
	// sentence, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading input:", err)
	// 	return
	// }

	// fmt.Printf("The number of letters in the sentence is: %d\n", countLetters(sentence))

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a timestring (e.g., '2023-10-01 15:30 MST'): ")

	timestring, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if timestring[len(timestring)-1] == '\n' {
		timestring = timestring[:len(timestring)-1]
	}

	newYorkTime, londonTime, tokyoTime, err := parseTimeString(timestring)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New York Time:", newYorkTime)
	fmt.Println("London Time:", londonTime)
	fmt.Println("Tokyo Time:", tokyoTime)

}
