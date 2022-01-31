package stringOps

import (
	"fmt"

)

func EvenChecker(str string) bool {
	if len (str) % 2 == 1 {
		return false
	} else {
		return true
	}
}