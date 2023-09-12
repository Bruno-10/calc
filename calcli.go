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
	sg, ts, err := calculate(text)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range sg {
		fmt.Printf("Group %d: %f", i, v)
		fmt.Println()
	}

	fmt.Printf("Total calculate: %f", ts)
	fmt.Println()
}
func calculate(text string) ([]float64, float64, error) {

	var sg []float64
	var cs float64
	var cns string
	var ts float64

	for _, v := range text {
		vString := string(v)

		if vString == "+" && cns != "" {
			cf, err := strconv.ParseFloat(cns, 64)
			if err != nil {
				return nil, 0, err
			}
			cs += cf
			cns = ""
			continue
		}

		cns += vString

		if vString == ";" || vString == "\n" {
			cns, _ = strings.CutSuffix(cns, vString)
			if cns != "" {
				cf, err := strconv.ParseFloat(cns, 64)
				if err != nil {
					return nil, 0, err
				}
				cs += cf
				cns = ""
			}
			if cs != 0 {
				sg = append(sg, cs)
				ts += cs
				cs = 0
			}
		}
	}

	return sg, ts, nil
}
