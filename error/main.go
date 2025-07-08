package main

import (
	"errors"
	"fmt"
	"strconv"
)

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("both values cannot be zero")
	}

	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("both values cannot be zero: a=%d, b=%d", a, b)
	}

	return nil
}

func main() {

	err := check(0, 10)
	if err != nil {
		fmt.Println("!Check() se ejecuto normalmente!!!")
	} else {
		fmt.Println(err)
	}

	err = check(0, 0)
	if err.Error() == "both values cannot be zero" {
		fmt.Println("!Check() se ejecuto normalmente!!!")
	} else {
		fmt.Println(err)
	}

	err = formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("-1")
	if err == nil {
		fmt.Println("Valor entero:", i)
	}

	i, err = strconv.Atoi("Y1")
	if err != nil {
		fmt.Println(err)
	}
}
