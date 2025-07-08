package main

import (
	"fmt"
	"strconv"
)

// funcion que retorna la hora en formato de 12 horas

func timeConversion(s string) string {

	hour, minute, second := s[:2], s[3:5], s[6:8]
	hourInt, err := strconv.Atoi(hour)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if s[8:] == "PM" && hour != "12" {
		hourInt += 12
	}
	if s[8:] == "AM" && hour == "12" {
		hourInt = 00
	}

	return fmt.Sprintf("%02d:%s:%s\n", hourInt, minute, second)

}

func main() {
	timeConversion("07:05:45PM")
	timeConversion("12:05:45AM")
}
