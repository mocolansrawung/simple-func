package utils

import (
	"fmt"
)

func scanInputs(n int, inputType string) {
	switch inputType {
	case "int":
		inputs := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Printf("Enter input %d (integer): ", i+1)
			_, err := fmt.Scan(&inputs[i])
			if err != nil {
				fmt.Println("Error reading input: ", err)
				return
			}
		}
	case "string":
		inputs := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Printf("Enter input %d (string): ", i+1)
			_, err := fmt.Scan(&inputs[i])
			if err != nil {
				fmt.Println("Error reading input: ", err)
				return
			}
		}
	default:
		fmt.Println("inputType not supported right now.")
	}
}
