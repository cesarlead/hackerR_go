package main

import "fmt"

func plusMinus(arr []int) {

	var positiveCount, negativeCount, zeroCount int
	sizeArr := len(arr)

	for _, value := range arr {
		if value > 0 {
			positiveCount++
		}
		if value < 0 {
			negativeCount++
		}
		if value == 0 {
			zeroCount++
		}
	}

	positiveRatio := float64(positiveCount) / float64(sizeArr)
	negativeRatio := float64(negativeCount) / float64(sizeArr)
	zeroRatio := float64(zeroCount) / float64(sizeArr)

	fmt.Printf("%.6f\n", positiveRatio)
	fmt.Printf("%.6f\n", negativeRatio)
	fmt.Printf("%.6f\n", zeroRatio)

}

func main() {

	arr := []int{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}
