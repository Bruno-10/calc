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

	for i, o := range sg {
		fmt.Printf("Group %d: %f", i, o)
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

	var ts float64
	var lo func(float64, float64) float64
	var cs float64

	cns := "0"

	ops := map[string]func(float64, float64) float64{
		"+": sum,
		"*": multiply,
		"/": divide,
	}

	for _, v := range text {
		vString := string(v)

		o, ok := ops[vString]
		if ok && cns != "" {
			cf, err := strconv.ParseFloat(cns, 64)
			if err != nil {
				return 0, err
			}

			if cs == 0 && vString != "+" {
				cs = 1
			}
			cs = o(cs, cf)
			lo = o
			cns = ""
			continue
		}

		cns += vString

		if vString == "," || vString == "\n" {
			cns, _ = strings.CutSuffix(cns, vString)
			if cns != "" {
				cf, err := strconv.ParseFloat(cns, 64)
				if err != nil {
					return 0, err
				}
				cs = lo(cs, cf)
				cns = ""
			}
			if cs != 0 {
				ts += cs
				cs = 0
			}
		}
	}

	return ts, nil
}

func execute(text string) (float64, []float64, error) {
	var sg []float64
	var total float64
	sections := strings.Split(text, ",")

	for _, v := range sections {
		sums := strings.Split(v, "+")

		var st float64
		for _, sv := range sums {
			ts, err := calculate(sv + "\n")
			if err != nil {
				return 0, nil, err
			}
			st += ts
		}
		sg = append(sg, st)
	}

	for _, v := range sg {
		total += v
	}

	return total, sg, nil
}
