// This package is built on Gonum Project (https://github.com/gonum/gonum)
// Gonum Project's copyright notice is included below as requested. 

/*
Copyright ©2013 The Gonum Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:
    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in the
      documentation and/or other materials provided with the distribution.
    * Neither the name of the Gonum project nor the names of its authors and
      contributors may be used to endorse or promote products derived from this
      software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/


package permutation

import (
	"gonum.org/v1/gonum/stat/combin" // Example function: https://github.com/gonum/gonum/blob/master/stat/combin/combinations_example_test.go#L198
)

/*
func ExamplePermutations_index() {
	// The integer slices returned from Permutations can be used to index
	// into a data structure.
	data := []string{"a", "b", "c", "d"}
	cs := combin.Permutations(len(data), 2)
	for _, c := range cs {
		fmt.Printf("%s%s\n", data[c[0]], data[c[1]])
	}

	// Output:
	// ab
	// ba
	// ac
	// ca
	// ad
	// da
	// bc
	// cb
	// bd
	// db
	// cd
	// dc
}
*/

func Permutate(set []rune, dblIndex []int) [][]rune {
	cs := combin.Permutations(len(set), len(dblIndex))

	perms := make([][]rune, len(cs))
	for i := range perms {
		perms[i] = make([]rune, len(dblIndex))
		for j, k := range cs[i]{
			perms[i][j] = set[k]
		}
	}
	for _,  i := range set {
		t := make([]rune, len(dblIndex))
		for j := range dblIndex {
			t[j] = i
		}
		perms = append(perms, t)	
	}
	return perms
}