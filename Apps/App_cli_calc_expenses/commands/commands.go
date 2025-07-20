package commands

// /commands/commands.go

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Menu() string {
	fmt.Println("--- Expense Calculator Menu ---")
	fmt.Println("1. Add expenses")
	fmt.Println("2. Show expenses details")
	fmt.Println("3. Create expense file")
	fmt.Println("4. Read expense file")
	fmt.Println("0. Exit")
	fmt.Println("-----------------------------")

	for {
		input, err := GetInputUser()
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch input {
		case "0", "1", "2", "3", "4":
			return input
		default:
			fmt.Println("Invalid option. Please enter 0, 1, 2, 3, or 4.")
		}
	}
}

func GetInputUser() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	input = strings.Replace(input, "\n", "", -1)
	return input, nil
}

func ReadDirectory(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func ReadFile(nameFile string) string {

	nameFile = nameFile + ".txt"

	if !ReadDirectory(nameFile) {
		fmt.Printf("Error: File '%s' does not exist in the current directory.\n", nameFile)
		return ""
	}

	content, err := os.ReadFile(nameFile)
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", nameFile, err)
		return ""
	}

	return string(content)

}

func CreateFile(content string, nameFile string) {

	if nameFile == "" {
		nameFile = "expenses.txt"
	} else {
		nameFile = nameFile + ".txt"
	}

	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Printf("Error creating file '%s': %v\n", nameFile, err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to file '%s': %v\n", nameFile, err)
		return
	}

	fmt.Printf("File '%s' created successfully!\n", nameFile)

}
