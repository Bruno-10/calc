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
	secondImpl(text)
}
func firstImpl(text string) {

	var sg []float64
	var cs float64
	var cns string
	var ts float64

	for _, v := range text {
		vString := string(v)

		if vString == "+" {
			cf, err := strconv.ParseFloat(cns, 64)
			if err != nil {
				fmt.Println(err)
			}
			cs += cf
			cns = ""
			continue
		}

		cns += vString

		if vString == ";" || vString == "\n" {
			cns, _ = strings.CutSuffix(cns, vString)
			cf, err := strconv.ParseFloat(cns, 64)
			if err != nil {
				fmt.Println(err)
			}
			cs += cf
			sg = append(sg, cs)
			ts += cs
			cs = 0
			cns = ""
		}
	}

	// for i, v := range sg {
	// 	fmt.Printf("Group %d: %f", i, v)
	// 	fmt.Println()
	// }

	// fmt.Printf("Total sum: %f", ts)
	// fmt.Println()
}
func secondImpl(text string) {
	var sg []float64
	var cs float64
	var cns string
	var ts float64

	for _, v := range text {
		vString := string(v)

		_, err := strconv.ParseFloat(vString, 64)
		if err != nil {
			if vString == "." {
				cns += vString
				continue
			}
			cf, err := strconv.ParseFloat(cns, 64)
			if err != nil {
				fmt.Println(err)
			}
			cs += cf
			cns = ""

			if vString == ";" || vString == "\n" {
				sg = append(sg, cs)
				// fmt.Printf("Group %d: %f", len(sg), cs)
				// fmt.Println()
				ts += cs
				cs = 0
			}
			continue
		}

		cns += vString
	}

	// fmt.Printf("Total sum: %f", ts)
	// fmt.Println()
}
