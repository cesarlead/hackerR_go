package main

// Apps/App_cli_calc_expenses/main.go

import (
	"fmt"
	"practica/Apps/App_cli_calc_expenses/commands"
	"practica/Apps/App_cli_calc_expenses/expenses"
)

// Caracteristicas de la App
// 1. Obtener todos los gastos infinitamente hasta ingresar un comando como cls o exit *
// 2. Poder obtener el promedio, valores maximos, minimos, promedio, y total *
// 3. Mostrar un resuemn en la consola de lo ingresado *
// 4. Exportar la informacion a un .txt

func main() {

	var data []float64
	var content string

	for {
		init := commands.Menu()

		switch init {
		case "1":
			fmt.Println("Entering expense capture mode...")
			data = expenses.CapExpense()
			if len(data) > 0 {
				content = expenses.PrintDetailsExpenses(data)
				fmt.Println("Expenses captured. Use option 2 to view details or option 3 to save.")
			} else {
				fmt.Println("No expenses were entered.")
				content = ""
			}

		case "2":
			if len(data) == 0 {
				fmt.Println("No expenses have been entered yet. Please use option 1 first.")
				break
			}
			if content == "" {
				content = expenses.PrintDetailsExpenses(data)
			}
			fmt.Println(content)
		case "3":
			if len(data) == 0 || content == "" {
				fmt.Println("No expenses captured or details generated yet. Please use option 1 to add expenses first.")
				break
			}
			fmt.Print("Enter filename to save (e.g., my_expenses): ")
			nameFile, err := commands.GetInputUser()
			if err != nil {
				fmt.Println("Error getting filename:", err)
				continue
			}
			if nameFile == "" {
				fmt.Println("Filename cannot be empty. Using default 'expenses'.")
				nameFile = "expenses"
			}
			content = expenses.PrintDetailsExpenses(data)
			commands.CreateFile(content, nameFile)
		case "4":
			fmt.Print("Enter filename to read (e.g., my_expenses): ")
			nameFile, err := commands.GetInputUser()
			if err != nil {
				fmt.Println("Error getting filename:", err)
				continue
			}
			if nameFile == "" {
				fmt.Println("No filename entered. Please provide a name.")
				continue
			}
			fileContent := commands.ReadFile(nameFile)
			if fileContent != "" {
				fmt.Println("\n--- Content of", nameFile+".txt ---")
				fmt.Println(fileContent)
				fmt.Println("----------------------------------\n")
			} else {
				fmt.Println("Could not read file or file is empty.")
			}
		case "0":
			fmt.Println("Closing App...")
			break

		}
	}

}
