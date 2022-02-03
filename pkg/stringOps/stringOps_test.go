/*
Test (table-test format), Benchmark, Examle implementation for the following fuctions in stringOps package:
- EvenCheck
- Seperate
- SymmCheck
*/
package stringOps

import (
	"testing"
	"fmt"
)

////////////////////////////////////////////
////		FUNCTION -> EvenCheck		////
////////////////////////////////////////////

// ### EvenCheck - Test

type evenTest struct {
	arg string; expected bool
}

var evenTests = []evenTest{
	{"arezzo", true},
	{"si", true},
	{"lucca", false},
	{"n", false}, 
}

func TestEvenCheck(t *testing.T) {
	for i, test := range evenTests {
		t.Logf("Test no: %v", i+1)
		t.Logf("Test input is: '%v'. Expected output is '%v'", test.arg, test.expected)
		t.Logf("Test output is: %v", EvenCheck(test.arg))

		if output := EvenCheck(test.arg); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}

// ### EvenCheck - Benchmark

func BenchmarkEvenCheck(b *testing.B){
	for i := 0; i < b.N; i++ {
		EvenCheck("pecorino")
	}
}

// ### EvenCheck - Example

func ExampleEvenCheck()  {
	fmt.Println(EvenCheck("pisa"))
	// Output: true
}


////////////////////////////////////////////
////		FUNCTION -> Seperate		////
////////////////////////////////////////////

// ### Seperate - Test

type sepTest struct {
	arg string; expected []rune
}

var sepTests = []sepTest{
	{"Firenze", []rune{102, 105, 114, 101, 110, 122, 101}},
	{"firenze", []rune{102, 105, 114, 101, 110, 122, 101}},
	{"vinci", []rune{118, 105, 110, 99, 105}},
	{"Lorenzana", []rune{108, 111, 114, 101, 110, 122, 97, 110, 97}},
}

func TestSeperate(t *testing.T) {
	for i, test := range sepTests {
		t.Logf("Test no: %v", i+1)
		t.Logf("Test input is: '%v'. Expected output is '%v'", test.arg, test.expected)
		t.Logf("Test output is: %v", Seperate(test.arg))
		
		for i, v := range test.expected {
			res := Seperate(test.arg)
			if v != res[i] {
				t.Errorf("Output %v not equal to expected %v", Seperate(test.arg), test.expected)
			}
		} 	
	}	
}

// ### Seperate - Benchmark

func BenchmarkSeperate(b *testing.B){
	for i := 0; i < b.N; i++ {
		Seperate("pecorino")
	}
}

// ### Seperate - Example

func ExampleSeperate()  {
	fmt.Println(Seperate("pistoia"))
	// Output: [112 105 115 116 111 105 97]
}


////////////////////////////////////////////
////		FUNCTION -> SymmCheck		////
////////////////////////////////////////////

// ### SymmCheck - Test

type symmTest struct {
	arg1 rune ; arg2 []rune ; expected bool
}

var symmTests = []symmTest{
	{ 63, []rune{97, 98, 98, 97}, true },
	{ 63, []rune{109, 97, 115, 115, 97, 63}, true },
	{ 42, []rune{109, 97, 115, 115, 97, 42}, true },
	{ 63, []rune{102, 105, 114, 101, 110, 122, 101}, false },
	{ 63, []rune{102, 105, 99, 111, 63, 105, 116}, false },
}

func TestSymmCheck(t *testing.T) {
	for i, test := range symmTests {
		t.Logf("Test no: %v", i+1)
		t.Logf("Test input is: wcard->'%v', array->'%v'. Expected output is '%v'", test.arg1, test.arg2, test.expected)
		t.Logf("Test output is: %v", SymmCheck(test.arg1, test.arg2))
		
		if output := SymmCheck(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}	
	}	
}

// ### SymmCheck - Benchmark

func BenchmarkSymmCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SymmCheck(63, []rune{109, 97, 115, 115, 97, 63})
	}
}

// ### SymmCheck - Example

func ExampleSymmCheck()  {
	fmt.Println(SymmCheck(42, []rune{109, 97, 115, 115, 97, 42}))
	//Output: true
}


////////////////////////////////////////////
////		FUNCTION -> CharSet			////
////////////////////////////////////////////

// ### CharSet - Test

type chSetTest struct {
	arg1 rune ; arg2 []rune ; expected []rune
}

var chSetTests = []chSetTest{
	{ 63, []rune{97, 98, 98, 97}, []rune{97, 98} },
	{ 63, []rune{109, 97, 115, 115, 97, 63}, []rune{109, 97, 115} },
	{ 42, []rune{109, 97, 115, 115, 97, 42}, []rune{109, 97, 115} },
	{ 63, []rune{102, 105, 114, 101, 110, 122, 101}, []rune{102, 105, 114, 101, 110, 122} },
	{ 63, []rune{102, 105, 99, 111, 63, 105, 116}, []rune{102, 105, 99, 111, 116} },
}

func TestCharSet(t *testing.T) {
	for i, test := range chSetTests {
		t.Logf("Test no: %v", i+1)
		t.Logf("Test input is: wcard->'%v', rune array->'%v'. Expected output is '%v'", test.arg1, test.arg2, test.expected)
		t.Logf("Test output is: %v", CharSet(test.arg1, test.arg2))
		
		for i, v := range test.expected {
			res := CharSet(test.arg1, test.arg2)
			if v != res[i] {
				t.Errorf("Output %v not equal to expected %v", CharSet(test.arg1, test.arg2), test.expected)
			}
		} 	
	}	
}

// ### CharSet - Benchmark

func BenchmarkCharSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CharSet(63, []rune{109, 97, 115, 115, 97, 63})
	}
}

// ### CharSet - Example

func ExampleCharSet()  {
	fmt.Println(CharSet(63, []rune{102, 105, 114, 101, 110, 122, 101}))
	// Output: [102 105 114 101 110 122]
}


////////////////////////////////////////////
////		FUNCTION -> WCardIndex		////
////////////////////////////////////////////

// ## WCardIndex - Test

type wciTest struct {
	arg1 rune ; arg2 []rune ; expected1 []int; expected2 []int
}

var wciTests = []wciTest{
	{ 63, []rune{97, 98, 98, 97}, []int{}, []int{} },
	{ 63, []rune{97, 63, 98, 97}, []int{1}, []int{} },
	{ 35, []rune{35, 97, 115, 115, 97, 35}, []int{}, []int{0} },
	{ 42, []rune{42, 97, 115, 42, 97, 42}, []int{3}, []int{0} },
	{ 36, []rune{36, 105, 36, 36, 110, 122, 36}, []int{2}, []int{0, 3} },
	{ 36, []rune{36, 36, 36, 36, 110, 122, 36}, []int{1, 2}, []int{0, 3} },
}

func TestWCardIndex(t *testing.T)  {
	for i, test := range wciTests {
		t.Logf("Test no: %v", i+1)
		t.Logf("Test input: wcard->'%v', rune array->'%v'. Expected output: sngWcards(index)->'%v', dblWcards(index)->'%v'", test.arg1, test.arg2, test.expected1, test.expected2)
		x, y := WCardIndex(test.arg1, test.arg2)
		t.Logf("Test output: sngWcards(index)->'%v', dblWcards(index)->'%v'", x, y)
		
		for i, v := range test.expected1 {
			if v != x[i] {
				t.Errorf("Output in sngWCIndex: %v not equal to expected %v", x, v)
			}
		} 	
		for i, v := range test.expected2 {
			if v != y[i] {
				t.Errorf("Output in dblWCIndex: %v not equal to expected %v", y, v)
			}
		}			
	}	
}

// ### WCardIndex - Benchmark

func BenchmarkWCardIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WCardIndex(63, []rune{109, 97, 115, 115, 97, 63})
	}
}

func ExampleWCardIndex()  {
	fmt.Println(WCardIndex(36, []rune{36, 36, 36, 36, 110, 122, 36}))
	// Output: [1 2] [0 3]
}