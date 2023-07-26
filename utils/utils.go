package utils

import (
	"fmt"
)

func scanInputs(n int, inputType string) (inputInt []int, inputStr []string) {
	switch inputType {
	case "int":
		var inputInt []int
		for i := 0; i < n; i++ {
			fmt.Printf("Enter input %d (integer): ", i+1)
			_, err := fmt.Scan(&inputInt[i])
			if err != nil {
				fmt.Println("Error reading input: ", err)
				return inputInt, nil
			}
		}
	case "string":
		var inputStr []string
		for i := 0; i < n; i++ {
			fmt.Printf("Enter input %d (string): ", i+1)
			_, err := fmt.Scan(&inputStr[i])
			if err != nil {
				fmt.Println("Error reading input: ", err)
				return nil, inputStr
			}
		}
	default:
		fmt.Println("inputType not supported right now.")
	}

	return nil, nil
}

func readAndWriteInputs() string {
	_, inputStr := scanInputs(2, "string")
	if inputStr != nil {
		result := inputStr[0] + inputStr[1]
		fmt.Println(result)
		return result
	}
	return "Input error."
}

// func appendAndPrintInputs(a, b string) {
// 	fmt.Printf("%s %s", a, b)
// 	return
// }

// func readIntAndTimes() {
// }

// // fmt.Println("You entered the following inputs:")
// // for i, val := range inputs {
// // 	fmt.Printf("Input %d: %d\n", i+1, val)
// // }
