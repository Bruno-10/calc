package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/Bruno-10/calc/business/core/calc"
)

func main() {
	fmt.Print("Enter text: ")

	// TODO: YOU CAN'T IGNORE THE ERROR!!!!
	r := bufio.NewReader(os.Stdin)
	text, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result, err := calc.Execute(context.Background(), text)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, op := range result.SumGroup {
		fmt.Printf("Group %d: %f", i, op)
		fmt.Println()
	}

	fmt.Printf("Total calculate: %f", result.Total)
	fmt.Println()
}
