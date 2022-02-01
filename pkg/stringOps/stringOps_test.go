package stringOps

import (
	"testing"
	"fmt"
)


// ### EvenChecker - Test

type evenTest struct {
	arg string; expected bool
}

var evenTests = []evenTest{
	{"arezzo", true},
	{"si", true},
	{"lucca", false},
	{"n", false}, 
}

func TestEvenChecker(t *testing.T) {
	for i, test := range evenTests {
		t.Logf("Test no: %v", i+1)
		t.Logf("Test input is: '%v'. Expected output is '%v'", test.arg, test.expected)
		t.Logf("Test output is: %v", EvenChecker(test.arg))

		if output := EvenChecker(test.arg); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}

// ### EvenChecker - Benchmark

func BenchmarkEvenChecker(b *testing.B){
	for i := 0; i < b.N; i++ {
		EvenChecker("pecorino")
	}
}

// ### EvenChecker - Example
func ExampleEvenChecker()  {
	fmt.Println(EvenChecker("pisa"))
	fmt.Println(EvenChecker("pistoia"))
}