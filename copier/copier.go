package copier

import (
	"golang.org/x/exp/slices"
)

// Pointer returns a new pointer to the copied value of the
// original argument value.
func Pointer[T any](original *T) (copied *T) {
	if original == nil {
		return nil
	}
	copied = new(T)
	*copied = *original
	return copied
}

// Slice returns a new slice with each element of the
// original slice copied.
func Slice[T any](original []T) (copied []T) {
	return slices.Clone(original)
}
