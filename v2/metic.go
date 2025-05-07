package metic

import (
	"golang.org/x/exp/constraints"
)

// An Add is a genric function.
// it accepts below types:
//   - int
//   - float
//
// and perform addition math operation
// Read more about the math operation here: [mathsisfun]
//
// [mathsisfun]: https://mathsisfun.com/numbers/addition.html
func Add[T constraints.Integer | constraints.Float](a T, b T) T {
	return a + b
}
