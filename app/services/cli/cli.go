package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/Bruno-10/calc/business/core/calc"
	"github.com/Bruno-10/calc/foundation/logger"
)

func main() {
	core := calc.NewCore(logger.New(os.Stdout, logger.LevelInfo, "CALC", func(ctx context.Context) string { return "" }))

	fmt.Print("Enter text: ")

	r := bufio.NewReader(os.Stdin)
	text, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result, err := core.Execute(context.Background(), text)
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
