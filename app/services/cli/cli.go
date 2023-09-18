package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/Bruno-10/calc/business/core/calc"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := r.ReadString('\n')
	result, err := calc.Execute(context.Background(), text)
	if err != nil {
		fmt.Println(err)
	}

	for i, op := range result.SumGroup {
		fmt.Printf("Group %d: %f", i, op)
		fmt.Println()
	}

	fmt.Printf("Total calculate: %f", result.Total)
	fmt.Println()
}
