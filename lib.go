// Package mygomod provides simple math operations.
package my_go_mod

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add returns the sum of two integers.
//
// For more information about addition, visit:
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
