package calc

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Bruno-10/calc/foundation/logger"
)

// TODO: Exported identifier without comment.
type Result struct {
	Total    float64   `json:"total"`
	SumGroup []float64 `json:"sumGroup"`
}

// Core manages the set of APIs for user access.
type Core struct {
	Execute func(ctx context.Context, text string) (Result, error)
	log     *logger.Logger
}

// NewCore constructs a core for user api access.
func NewCore(log *logger.Logger) *Core {
	return &Core{
		Execute: Execute,
		log:     log,
	}
}

// TODO: Unexported doesn't need a comment, but if you will add it, all comments
// must be proper sentences with proper sentence structure.
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

func Execute(ctx context.Context, text string) (Result, error) {
	var sg []float64
	var total float64

	// TODO: IT IS A SMELL WHEN USING THE BLANK IDENTIFIER. CHECK THAT BOOLEAN.
	text, _ = strings.CutSuffix(text, ",")
	sections := strings.Split(text, ",")

	for _, v := range sections {
		if v != "\n" {
			sums := strings.Split(v, "+")

			var st float64
			for _, sv := range sums {
				ts, err := calculate(sv + "\n")
				if err != nil {
					return Result{}, fmt.Errorf("Calculate error: %w", err)
				}
				st += ts
			}

			total += st
			sg = append(sg, st)
		}
	}

	return Result{total, sg}, nil
}
