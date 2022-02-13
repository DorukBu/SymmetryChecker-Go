package stringOps

import (
	"strings"
	"math"
)


// check if given string has even number of characters
func EvenCheck(str string) bool {
	if len (str) % 2 == 1 {
		return false
	} else {
		return true
	}
}


// Seperate string into an array with lowercase letters
func Seperate(str string) []rune  {
	arr := []rune(strings.ToLower(str))
	return arr
}


// Check symmetry of the given array of runes
func SymmCheck(wC rune, arr []rune) (sym bool) {
	sym = true
	for i := 0; i <= len(arr)/2-1; i++ {
		if arr[i] == wC || arr[len(arr)-1-i] == wC {
			continue
		} else {
			if arr[i] != arr[len(arr)-1-i] {
				sym = false
				return 
			} else {
				continue
			}
		}
	}
	return
}


// Create a slice of runes containing unique elements
func CharSet(wC rune, arr []rune) (set []rune) {
	for _, i := range arr {
		if i == wC {
			continue
		}
		ap := true
		for _, j := range set {
			if i == j {
				ap = false
			} else {
				continue
			}
		}
		if ap == true {
			set = append(set, i)
		} 
	}
	return
}


// Wildcard index
func WCardIndex(wC rune, arr []rune) (sngIndex []int, dblIndex []int) {
	for i := 0 ; i <= int(math.Ceil(float64(len(arr))/2-1)) ; i++ {
		a := arr[i]
		b := arr[len(arr)-1-i]
		if a == b && a == wC {
			dblIndex = append(dblIndex, i)
		} else if a == wC {
			sngIndex = append(sngIndex, i)
		} else if b == wC {
			sngIndex = append(sngIndex, len(arr)-1-i)
		} else {
			continue
		}
	}
	return sngIndex, dblIndex
}

// Unify -> unify permutations of runes ([][]rune) into array of strings ([]string)
func Unify(perms [][]rune) []string {
	uni := make([]string, len(perms))
	for i := range perms {
		str := string(perms[i][:])
		uni[i] = str
	}
	return uni
} 