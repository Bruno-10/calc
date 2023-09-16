package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := r.ReadString('\n')
	ts, sg, err := execute(text)
	if err != nil {
		fmt.Println(err)
	}

	for i, op := range sg {
		fmt.Printf("Group %d: %f", i, op)
		fmt.Println()
	}

	fmt.Printf("Total calculate: %f", ts)
	fmt.Println()
}

// sum adds two float64
func sum(cs float64, cf float64) float64 {
	return cs + cf
}

// multiply multiplies two float64
func multiply(cs float64, cf float64) float64 {
	return cs * cf
}

// divide divides two float64
func divide(cs float64, cf float64) float64 {
	return cs / cf
}

// calculate takes a string and while it iterates through it, it does two things
// executes operations, and separates into groups when it encounters a ,
func calculate(text string) (float64, error) {
	text = "+" + text

	var lop func(float64, float64) float64
	var cs float64
	var lns string
	var cns string

	ops := map[string]func(float64, float64) float64{
		"+": sum,
		"*": multiply,
		"/": divide,
	}

	for i, v := range text {
		vString := string(v)

		cns += vString

		op, ok := ops[vString]
		if ok {
			cns, _ = strings.CutSuffix(cns, vString)
			lop = op
			if cns != "" {
				lns = cns
				cns = ""
			}

			continue
		}

		if i+1 < len(text) {
			_, nopok := ops[string(text[i+1])]
			if cns != "" && lns != "" && lop != nil && (nopok || string(text[i+1]) == "\n") {
				cf, err := strconv.ParseFloat(cns, 64)
				if err != nil {
					return 0, fmt.Errorf("Parsing cf: %w", err)
				}

				lf, err := strconv.ParseFloat(lns, 64)
				if err != nil {
					return 0, fmt.Errorf("Parsing lf: %w", err)
				}

				cs = lop(lf, cf)
				cns = fmt.Sprintf("%f", cs)
				lns = ""
				lop = nil
			}
			continue
		}

		if cns != "" && vString == "\n" && lop != nil {
			cns = strings.ReplaceAll(cns, vString, "")
			cf, err := strconv.ParseFloat(cns, 64)
			if err == nil {
				cs = cf
			}
		}
	}

	return cs, nil
}

func execute(text string) (float64, []float64, error) {
	var sg []float64
	var total float64
	text, _ = strings.CutSuffix(text, ",")
	sections := strings.Split(text, ",")
	for _, v := range sections {
		if v != "\n" {
			sums := strings.Split(v, "+")

			var st float64
			for _, sv := range sums {
				ts, err := calculate(sv + "\n")
				if err != nil {
					return 0, nil, fmt.Errorf("Calculate error: %w", err)
				}
				st += ts
			}
			total += st
			sg = append(sg, st)
		}
	}

	return total, sg, nil
}
