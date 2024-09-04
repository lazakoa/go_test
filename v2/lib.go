// A sample package to test imports.
package go_test

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add two numbers together.
// [Addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a T, b T) T {
	return a + b
}
