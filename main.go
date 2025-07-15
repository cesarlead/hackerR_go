package main

import (
	"fmt"
	hackerr "practica/hackerR"
)

func main() {

	fmt.Printf("****** Hello CesarLead ******\n Ejercicios de HackerRank:\n")

	fmt.Println(hackerr.PlusMinus([]int{-4, 3, -9, 0, 4, 1}))

	fmt.Println(hackerr.MiniMaxSum([]int{1, 2, 3, 4, 5}))

	fmt.Println(hackerr.TimeConversion("07:05:45PM"))

	fmt.Println(hackerr.SparseArrays([]string{"abc", "ba", "cd", "def", "ghi", "jkl", "mno", "pqr", "cd", "vwx", "cd"}, []string{"ba", "b", "cd", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yz"}))

	fmt.Println(hackerr.LonelyInteger([]int{1, 2, 3, 4, 3, 2, 1}))

}
