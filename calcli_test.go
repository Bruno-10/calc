package main

import (
	"testing"
)

// CalcliTests holds methods for each calcli subtest. This type allows passing
// dependencies for test while still providing a convenient syntax when
// subtests are registered.
type CalcliTests struct {
}

func TestCalcli(t *testing.T) {
	tests := CalcliTests{}

	t.Run("testSum", tests.testSum())
	t.Run("testMultiply", tests.testMultiply())
	t.Run("testDivition", tests.testDivition())
	t.Run("testMix", tests.testMix())
}

// testSum tests different addition scenarios.
func (ct *CalcliTests) testSum() func(t *testing.T) {
	return func(t *testing.T) {
		input1 := "1+1+1,2+2+2"
		total1 := float64(9)

		tc, tg, err := execute(input1)
		if err != nil {
			t.Fatal(err)
		}

		if tg == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(tg) != 2 {
			t.Log(tg, tg)
			t.Fatalf("tg length for input: %s should equal 2, instead was: %d", input1, len(tg))
		}

		if tc != total1 {
			t.Fatalf("tc should be equal to total1. tc: %.4f | total: %.4f", tc, total1)
		}

		// Should behave well when input finished with ;
		input2 := "1+1+1,2+2+2,"
		total2 := float64(9)

		tc2, tg2, err := execute(input2)
		if err != nil {
			t.Fatal(err)
		}

		if tg2 == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(tg2) != 2 {
			t.Fatalf("tg2 length for input: %s should equal 2, instead was: %d", input2, len(tg2))
		}

		if tc2 != total2 {
			t.Fatalf("tc2 should be equal to total2. tc3: %.4f | total: %.4f", tc2, total2)
		}

		// Should behave well when input finishes with +
		input3 := "1+1+"
		total3 := float64(2)

		tc3, tg3, err := execute(input3)
		if err != nil {
			t.Fatal(err)
		}

		if tg3 == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(tg3) != 1 {
			t.Fatalf("tc3 length for input: %s should equal 1, instead was %d", input3, len(tg3))
		}

		if tc3 != total3 {
			t.Fatalf("Expected tc3 to be equal to total3. tc3: %.4f | total: %.4f", tc3, total3)
		}
	}
}

// testMultiply tests different multiplication scenarios.
func (ct *CalcliTests) testMultiply() func(t *testing.T) {
	return func(t *testing.T) {
		input1 := "2*2*2,3*3*3"
		total1 := float64(35)

		tc, tg, err := execute(input1)
		if err != nil {
			t.Fatal(err)
		}

		if tg == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(tg) != 2 {
			t.Log(tg, tg)
			t.Fatalf("tg length for input: %s should equal 2, instead was: %d", input1, len(tg))
		}

		if tc != total1 {
			t.Fatalf("tc should be equal to total1. tc: %.4f | total: %.4f", tc, total1)
		}

		// Should behave well when input finished with ;
		input2 := "2*2*2,3*3*3,"
		total2 := float64(35)

		tc2, tg2, err := execute(input2)
		if err != nil {
			t.Fatal(err)
		}

		if tg2 == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(tg2) != 2 {
			t.Fatalf("tg2 length for input: %s should equal 2, instead was: %d", input2, len(tg2))
		}

		if tc2 != total2 {
			t.Fatalf("tc2 should be equal to total2. tc3: %.4f | total: %.4f", tc2, total2)
		}

		// Should behave well when input finishes with +
		input3 := "2*2*"
		total3 := float64(4)

		tc3, tg3, err := execute(input3)
		if err != nil {
			t.Fatal(err)
		}

		if tg3 == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(tg3) != 1 {
			t.Fatalf("tc3 length for input: %s should equal 1, instead was %d", input3, len(tg3))
		}

		if tc3 != total3 {
			t.Fatalf("Expected tc3 to be equal to total3. tc3: %.4f | total: %.4f", tc3, total3)
		}
	}
}

// testDivition tests different dividing scenarios.
func (ct *CalcliTests) testDivition() func(t *testing.T) {
	return func(t *testing.T) {
		input1 := "10/5/2,8/2/2"
		total1 := float64(3)

		tc, tg, err := execute(input1)
		if err != nil {
			t.Fatal(err)
		}

		if tg == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(tg) != 2 {
			t.Log(tg, tg)
			t.Fatalf("tg length for input: %s should equal 2, instead was: %d", input1, len(tg))
		}

		if tc != total1 {
			t.Fatalf("tc should be equal to total1. tc: %.4f | total: %.4f", tc, total1)
		}

		// Should behave well when input finished with ;
		input2 := "10/5/2,8/2/2,"
		total2 := float64(3)

		tc2, tg2, err := execute(input2)
		if err != nil {
			t.Fatal(err)
		}

		if tg2 == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(tg2) != 2 {
			t.Fatalf("tg2 length for input: %s should equal 2, instead was: %d", input2, len(tg2))
		}

		if tc2 != total2 {
			t.Fatalf("tc2 should be equal to total2. tc3: %.4f | total: %.4f", tc2, total2)
		}

		// Should behave well when input finishes with +
		input3 := "10/5/"
		total3 := float64(2)

		tc3, tg3, err := execute(input3)
		if err != nil {
			t.Fatal(err)
		}

		if tg3 == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(tg3) != 1 {
			t.Fatalf("tc3 length for input: %s should equal 1, instead was %d", input3, len(tg3))
		}

		if tc3 != total3 {
			t.Fatalf("Expected tc3 to be equal to total3. tc3: %.4f | total: %.4f", tc3, total3)
		}
	}
}

// testMix test different operations inside one string
func (ct *CalcliTests) testMix() func(t *testing.T) {
	return func(t *testing.T) {
		input1 := "10/5/2*4+2"
		total1 := float64(6)

		tc, tg, err := execute(input1)
		if err != nil {
			t.Fatal(err)
		}

		if tg == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(tg) != 1 {
			t.Fatalf("tg length for input: %s should equal 1, instead was: %d", input1, len(tg))
		}

		if tc != total1 {
			t.Fatalf("tc should be equal to total1. tc: %.4f | total1: %.4f", tc, total1)
		}

		input2 := "20/10/2*4+2"
		total2 := float64(6)

		tc2, tg2, err := execute(input2)
		if err != nil {
			t.Fatal(err)
		}

		if tg2 == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(tg2) != 1 {
			t.Fatalf("tg2 length for input: %s should equal 1, instead was: %d", input2, len(tg2))
		}

		if tc2 != total2 {
			t.Fatalf("tc2 should be equal to total2. tc2: %.4f | total2: %.4f", tc2, total2)
		}

		input3 := "20/10/2*4+2"
		total3 := float64(6)

		tc3, tg3, err := execute(input3)
		if err != nil {
			t.Fatal(err)
		}

		if tg3 == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(tg3) != 1 {
			t.Fatalf("tg3 length for input: %s should equal 1, instead was: %d", input3, len(tg3))
		}

		if tc3 != total3 {
			t.Fatalf("tc3 should be equal to total3. tc3: %.4f | total3: %.4f", tc3, total3)
		}

		input4 := "20/10/2*4+2,3*3*3"
		total4 := float64(33)

		tc4, tg4, err := execute(input4)
		if err != nil {
			t.Fatal(err)
		}

		if tg4 == nil {
			t.Fatal("Groups slice shouldn't be empty")
		}

		if len(tg4) != 2 {
			t.Fatalf("tg4 length for input: %s should equal 1, instead was: %d", input4, len(tg4))
		}

		if tc4 != total4 {
			t.Fatalf("tc4 should be equal to total4. tc4: %.4f | total4: %.4f", tc4, total4)
		}
	}
}
