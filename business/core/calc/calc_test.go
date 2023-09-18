package calc

import (
	"context"
	"testing"
)

// CalcliTests holds methods for each calcli subtest. This type allows passing
// dependencies for test while still providing a convenient syntax when
// subtests are registered.
type CalcliTests struct {
}

func TestCalcli(t *testing.T) {
	tests := CalcliTests{}

	ctx := context.Background()
	t.Run("testSum", tests.testSum(ctx))
	t.Run("testMultiply", tests.testMultiply(ctx))
	t.Run("testDivition", tests.testDivition(ctx))
	t.Run("testMix", tests.testMix(ctx))
}

// testSum tests different addition scenarios.
func (ct *CalcliTests) testSum(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		input1 := "1+1+1,2+2+2"
		total1 := float64(9)

		result, err := Execute(ctx, input1)
		if err != nil {
			t.Fatal(err)
		}

		if result.SumGroup == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(result.SumGroup) != 2 {
			t.Log(result.SumGroup, result.SumGroup)
			t.Fatalf("result.SumGroup length for input: %s should equal 2, instead was: %d", input1, len(result.SumGroup))
		}

		if result.Total != total1 {
			t.Fatalf("result.Total should be equal to total1. result.Total: %.4f | total: %.4f", result.Total, total1)
		}

		// Should behave well when input finished with ;
		input2 := "1+1+1,2+2+2,"
		total2 := float64(9)

		result2, err := Execute(ctx, input2)
		if err != nil {
			t.Fatal(err)
		}

		if result2.SumGroup == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(result2.SumGroup) != 2 {
			t.Fatalf("result2.SumGroup length for input: %s should equal 2, instead was: %d", input2, len(result2.SumGroup))
		}

		if result2.Total != total2 {
			t.Fatalf("result2.Total should be equal to total2. result3.Total: %.4f | total: %.4f", result2.Total, total2)
		}

		// Should behave well when input finishes with +
		input3 := "1+1+"
		total3 := float64(2)

		result3, err := Execute(ctx, input3)
		if err != nil {
			t.Fatal(err)
		}

		if result3.SumGroup == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(result3.SumGroup) != 1 {
			t.Fatalf("result3.Total length for input: %s should equal 1, instead was %d", input3, len(result3.SumGroup))
		}

		if result3.Total != total3 {
			t.Fatalf("Expected result3.Total to be equal to total3. result3.Total: %.4f | total: %.4f", result3.Total, total3)
		}
	}
}

// testMultiply tests different multiplication scenarios.
func (ct *CalcliTests) testMultiply(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		input1 := "2*2*2,3*3*3"
		total1 := float64(35)

		result, err := Execute(ctx, input1)
		if err != nil {
			t.Fatal(err)
		}

		if result.SumGroup == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(result.SumGroup) != 2 {
			t.Log(result.SumGroup, result.SumGroup)
			t.Fatalf("result.SumGroup length for input: %s should equal 2, instead was: %d", input1, len(result.SumGroup))
		}

		if result.Total != total1 {
			t.Fatalf("result.Total should be equal to total1. result.Total: %.4f | total: %.4f", result.Total, total1)
		}

		// Should behave well when input finished with ;
		input2 := "2*2*2,3*3*3,"
		total2 := float64(35)

		result2, err := Execute(ctx, input2)
		if err != nil {
			t.Fatal(err)
		}

		if result2.SumGroup == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(result2.SumGroup) != 2 {
			t.Fatalf("result2.SumGroup length for input: %s should equal 2, instead was: %d", input2, len(result2.SumGroup))
		}

		if result2.Total != total2 {
			t.Fatalf("result2.Total should be equal to total2. result3.Total: %.4f | total: %.4f", result2.Total, total2)
		}

		// Should behave well when input finishes with +
		input3 := "2*2*"
		total3 := float64(4)

		result3, err := Execute(ctx, input3)
		if err != nil {
			t.Fatal(err)
		}

		if result3.SumGroup == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(result3.SumGroup) != 1 {
			t.Fatalf("result3.Total length for input: %s should equal 1, instead was %d", input3, len(result3.SumGroup))
		}

		if result3.Total != total3 {
			t.Fatalf("Expected result3.Total to be equal to total3. result3.Total: %.4f | total: %.4f", result3.Total, total3)
		}
	}
}

// testDivition tests different dividing scenarios.
func (ct *CalcliTests) testDivition(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		input1 := "10/5/2,8/2/2"
		total1 := float64(3)

		result, err := Execute(ctx, input1)
		if err != nil {
			t.Fatal(err)
		}

		if result.SumGroup == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(result.SumGroup) != 2 {
			t.Log(result.SumGroup, result.SumGroup)
			t.Fatalf("result.SumGroup length for input: %s should equal 2, instead was: %d", input1, len(result.SumGroup))
		}

		if result.Total != total1 {
			t.Fatalf("result.Total should be equal to total1. result.Total: %.4f | total: %.4f", result.Total, total1)
		}

		// Should behave well when input finished with ;
		input2 := "10/5/2,8/2/2,"
		total2 := float64(3)

		result2, err := Execute(ctx, input2)
		if err != nil {
			t.Fatal(err)
		}

		if result2.SumGroup == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(result2.SumGroup) != 2 {
			t.Fatalf("result2.SumGroup length for input: %s should equal 2, instead was: %d", input2, len(result2.SumGroup))
		}

		if result2.Total != total2 {
			t.Fatalf("result2.Total should be equal to total2. result3.Total: %.4f | total: %.4f", result2.Total, total2)
		}

		// Should behave well when input finishes with +
		input3 := "10/5/"
		total3 := float64(2)

		result3, err := Execute(ctx, input3)
		if err != nil {
			t.Fatal(err)
		}

		if result3.SumGroup == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(result3.SumGroup) != 1 {
			t.Fatalf("result3.Total length for input: %s should equal 1, instead was %d", input3, len(result3.SumGroup))
		}

		if result3.Total != total3 {
			t.Fatalf("Expected result3.Total to be equal to total3. result3.Total: %.4f | total: %.4f", result3.Total, total3)
		}
	}
}

// testMix test different operations inside one string
func (ct *CalcliTests) testMix(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		input1 := "10/5/2*4+2"
		total1 := float64(6)

		result, err := Execute(ctx, input1)
		if err != nil {
			t.Fatal(err)
		}

		if result.SumGroup == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(result.SumGroup) != 1 {
			t.Fatalf("result.SumGroup length for input: %s should equal 1, instead was: %d", input1, len(result.SumGroup))
		}

		if result.Total != total1 {
			t.Fatalf("result.Total should be equal to total1. result.Total: %.4f | total1: %.4f", result.Total, total1)
		}

		input2 := "20/10/2*4+2"
		total2 := float64(6)

		result2, err := Execute(ctx, input2)
		if err != nil {
			t.Fatal(err)
		}

		if result2.SumGroup == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(result2.SumGroup) != 1 {
			t.Fatalf("result2.SumGroup length for input: %s should equal 1, instead was: %d", input2, len(result2.SumGroup))
		}

		if result2.Total != total2 {
			t.Fatalf("result2.Total should be equal to total2. result2.Total: %.4f | total2: %.4f", result2.Total, total2)
		}

		input3 := "20/10/2*4+2"
		total3 := float64(6)

		result3, err := Execute(ctx, input3)
		if err != nil {
			t.Fatal(err)
		}

		if result3.SumGroup == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(result3.SumGroup) != 1 {
			t.Fatalf("result3.SumGroup length for input: %s should equal 1, instead was: %d", input3, len(result3.SumGroup))
		}

		if result3.Total != total3 {
			t.Fatalf("result3.Total should be equal to total3. result3.Total: %.4f | total3: %.4f", result3.Total, total3)
		}

		input4 := "20/10/2*4+2,3*3*3"
		total4 := float64(33)

		result4, err := Execute(ctx, input4)
		if err != nil {
			t.Fatal(err)
		}

		if result4.SumGroup == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(result4.SumGroup) != 2 {
			t.Fatalf("result4.SumGroup length for input: %s should equal 1, instead was: %d", input4, len(result4.SumGroup))
		}

		if result4.Total != total4 {
			t.Fatalf("result4.Total should be equal to total4. result4.Total: %.4f | total4: %.4f", result4.Total, total4)
		}
	}
}
