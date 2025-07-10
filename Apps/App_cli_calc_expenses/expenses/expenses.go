package expenses

import (
	"fmt"
	"practica/Apps/App_cli_calc_expenses/commands"
	"strconv"
	"strings"
)

func CapExpense() []float64 {
	var expensesList []float64

	fmt.Println("Enter expenses: ")
	fmt.Println("Enter 'cls' return to menu")

	for {
		input, err := commands.GetInputUser()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if input == "cls" {
			fmt.Println("to menu...")
			break
		}

		inputFloat, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}

		expensesList = append(expensesList, inputFloat)
	}

	return expensesList
}

func GenerateAverageMinMaxTotalExpenses(expenses []float64) (float64, float64, float64, float64) {
	total, average, minExpense, maxExpense := float64(0), float64(0), float64(0), float64(0)

	if len(expenses) == 0 {
		return total, average, minExpense, maxExpense
	}

	for _, expense := range expenses {
		total += expense

		if expense > maxExpense {
			maxExpense = expense
		}

		if expense < minExpense || minExpense == 0 {
			minExpense = expense
		}
	}
	average = total / float64(len(expenses))

	return total, average, minExpense, maxExpense
}

func PrintDetailsExpenses(expenses []float64) string {

	total, average, minExpense, maxExpense := GenerateAverageMinMaxTotalExpenses(expenses)

	buffer := strings.Builder{}

	buffer.WriteString(fmt.Sprintf("#########################################\n"))
	buffer.WriteString("Details Expenses: \n")

	for _, expense := range expenses {
		buffer.WriteString(fmt.Sprintf(" -> %.2f\n", expense))
	}

	buffer.WriteString(fmt.Sprintf("Total: %.2f\n", total))
	buffer.WriteString(fmt.Sprintf("Average: %.2f\n", average))
	buffer.WriteString(fmt.Sprintf("Min: %.2f\n", minExpense))
	buffer.WriteString(fmt.Sprintf("Max: %.2f\n", maxExpense))
	buffer.WriteString(fmt.Sprintf("#########################################\n"))

	return buffer.String()
}
