//go:build go1.18
// +build go1.18

package defaults

import (
	"github.com/qdm12/gosettings/merge"
	"golang.org/x/exp/constraints"
)

// NumberContraint represents a number type, it can notably be
// an integer, a float, and any type aliases of them such
// as `time.Duration`.
type NumberContraint interface {
	constraints.Integer | constraints.Float
}

// Number returns the existing argument if it is not zero,
// otherwise it returns the other argument.
func Number[T NumberContraint](existing, other T) (result T) { //nolint:ireturn
	return merge.WithNumber(existing, other)
}

// Pointer returns the existing argument if it is not nil.
// Otherwise it returns a new pointer to the defaultValue argument.
// For interfaces where the underlying type is not known,
// use Interface instead.
func Pointer[T any](existing *T, defaultValue T) (result *T) {
	if existing != nil {
		return existing
	}
	result = new(T)
	*result = defaultValue
	return result
}

// Interface returns the existing argument if it is not nil,
// otherwise it returns the defaultValue argument.
// Note you should NOT use this function with concrete pointers
// such as *int, and only use this for interfaces.
// This function is not type safe nor mutation safe, be careful.
func Interface(existing, defaultValue any) (result any) {
	return merge.WithInterface(existing, defaultValue)
}

// String returns the existing string argument if it is not empty,
// otherwise it returns the defaultValue string argument.
func String(existing, defaultValue string) (result string) {
	return merge.WithString(existing, defaultValue)
}

// CopiedSlice returns the existing slice argument if is not nil.
// Otherwise it returns a new slice with the copied values of the
// defaultValue slice argument.
// Note it is preferrable to use this function for added mutation safety
// on the result, but one can use Slice if performance matters.
func CopiedSlice[T any](existing, other []T) (result []T) {
	return merge.WithCopiedSlice(existing, other)
}

// Slice returns the existing slice argument if it is not nil,
// otherwise it returns the defaultValue slice argument.
func Slice[T any](existing, defaultValue []T) (result []T) {
	return merge.WithSlice(existing, defaultValue)
}

// Validator returns the existing argument if it is valid,
// otherwise it returns the defaultValue argument.
func Validator(existing, defaultValue merge.SelfValidator) ( //nolint:ireturn
	result merge.SelfValidator) {
	return merge.WithValidator(existing, defaultValue)
}
