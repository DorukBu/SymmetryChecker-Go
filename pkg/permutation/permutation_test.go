
/*
Test (table-test format), Benchmark, Examle implementation for the following fuctions in permutation package:
- Permutate
*/

package permutation

import (
	"fmt"
	"testing"
)

////////////////////////////////////////////
////        FUNCTION -> Permutate       ////
////////////////////////////////////////////

// ### Permutate - Test

type permTest struct {
	arg1 []rune ; arg2 []int; expected [][]rune
}

var permTests = []permTest{
	{ []rune{97, 98, 99, 100}, []int{2, 3},  [][]rune{ {97, 98}, {98, 97}, {97, 99}, {99, 97}, {97, 100}, {100, 97}, {98, 99}, {99, 98}, {98, 100}, {100, 98}, {99, 100}, {100, 99}, {97, 97}, {98, 98}, {99, 99}, {100, 100} } },
	{ []rune{97, 98, 99, 100}, []int{1},  [][]rune{ {97}, {98}, {99}, {100} } },
	{ []rune{32, 102, 101}, []int{3},  [][]rune{ {32}, {102}, {101} } },
}

func TestPermutate(t *testing.T) {
	for i, test := range permTests {
		t.Logf("Test no: %v", i+1)
		t.Logf("Test input: Char Set->'%v', dblWC indexes->'%v'. Expected output: Char Combins->'%v'", test.arg1, test.arg2, test.expected)
		x := Permutate(test.arg1, test.arg2)
		t.Logf("Test output: Combinations->'%v'", x)
		if len(test.expected) != len(x) {
			t.Errorf("Output: %v not equal to expected %v", x, test.expected)
		}

		for l, v := range test.expected {
			for j, k := range v {
				if k != x[l][j] {
					t.Errorf("Output: %v not equal to expected %v", x[l][j], k)
				}
			}
		} 			
	}	
}

// ## Permutate - Benchmark

func BenchmarkPermutate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Permutate([]rune{97, 98, 99, 100}, []int{2, 3})
	}
}

// ## Permutate - Example 

func ExamplePermutate()  {
	fmt.Println(Permutate([]rune{97, 98, 99, 100}, []int{2, 3}))
	// Output: [[97 98] [98 97] [97 99] [99 97] [97 100] [100 97] [98 99] [99 98] [98 100] [100 98] [99 100] [100 99] [97 97] [98 98] [99 99] [100 100]]
}