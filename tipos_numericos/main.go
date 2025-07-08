package main

import (
	"fmt"
)

// func divide(a, b int) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("Dont divide by zero")
// 	}
// 	return float64(a) / float64(b), nil
// }

// func sumAndCheckOverflow(a, b int32) (int16, bool) {
// 	sum := a + b
// 	if sum > math.MaxInt16 || sum < math.MinInt16 {
// 		return int16(sum), true
// 	}

// 	return int16(sum), false

// }

func calculatePorcentageVendor(venda, totalVendasMeta int) float64 {
	if totalVendasMeta == 0 {
		return 0.0
	}

	return (float64(venda) / float64(totalVendasMeta)) * 100.0
}

func main() {

	var a, b int32

	fmt.Println("Enter one number integer: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Error reading first number:", err)
		return
	}

	fmt.Println("Enter another number integer: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Error reading second number: ", err)
		return
	}

	// result, err := divide(a, b)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Printf("The result of dividing %d by %d is: %.2f\n", a, b, result)

	// result, overflow := sumAndCheckOverflow(a, b)
	// if overflow {
	// 	fmt.Printf("The sum of %d and %d is %d, which causes an overflow for int16.\n", a, b, result)
	// } else {
	// 	fmt.Printf("The sum of %d and %d is %d, which is within the range of int16.\n", a, b, result)
	// }

	result := calculatePorcentageVendor(int(a), int(b))
	fmt.Printf("The percentage of sales is: %.2f%%\n", result)
	fmt.Println("End of program")

}
