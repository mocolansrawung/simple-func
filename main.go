package main

import (
	"fmt"
	"simplefunc/utils/utils"
)

func main() {
	inputStr := utils.ReadAndWriteInputs()
	if inputStr == nil {
		fmt.Println("Error when read and write string inputs")
	}

	appendedString := utils.AppendAndPrintInputs(inputStr)
	if appendedString == nil {
		fmt.Println("Error when appending the strings")
	}

	utils.ReadIntAndWrite(appendedString)

	fmt.Printf("\n")
	fmt.Println("Program finished.")
}
