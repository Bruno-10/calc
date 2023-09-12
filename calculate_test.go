package main

import "testing"

func TestSum(t *testing.T) {
	input1 := "1+1+1;2+2+2\n"
	total1 := float64(9)

	sg, ts, err := calculate(input1)
	if err != nil {
		t.Fatal(err)
	}

	if sg == nil {
		t.Fatal("Groups slice shouldn't be empty")
	}

	if len(sg) != 2 {
		t.Log(sg, ts)
		t.Fatalf("sg length for input: %s should equal 2, instead was: %d", input1, len(sg))
	}

	if ts != total1 {
		t.Fatalf("ts should be equal to total1. ts: %.4f | total: %.4f", ts, total1)
	}

	// Should behave well when input finished with ;
	input2 := "1+1+1;2+2+2;\n"
	total2 := float64(9)

	sg2, ts2, err := calculate(input2)
	if err != nil {
		t.Fatal(err)
	}

	if sg2 == nil {
		t.Fatal("Groups slice shouldn't be empty")
	}

	if len(sg2) != 2 {
		t.Fatalf("sg2 length for input: %s should equal 2, instead was: %d", input2, len(sg2))
	}

	if ts2 != total2 {
		t.Fatalf("ts2 should be equal to total2. ts: %.4f | total: %.4f", ts2, total2)
	}

	// Should behave well when input finishes with +
	input3 := "1+1+\n"
	total3 := float64(2)

	sg3, ts3, err := calculate(input3)
	if err != nil {
		t.Fatal(err)
	}

	if sg3 == nil {
		t.Fatal("Groups slice shouldn't be empty")
	}

	if len(sg3) != 1 {
		t.Fatalf("sg3 length for input: %s should equal 1, instead was %d", input3, len(sg3))
	}

	if ts3 != total3 {
		t.Fatalf("Expected ts3 to be equal to total3. ts: %.4f | total: %.4f", ts3, total3)
	}
}
