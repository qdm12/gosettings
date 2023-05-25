//go:build go1.18
// +build go1.18

package override

import (
	"golang.org/x/exp/constraints"
)

// Number represents a number type, it can notably be
// an integer, a float, and any type aliases of them such
// as `time.Duration`.
type Number interface {
	constraints.Integer | constraints.Float
}

// WithNumber returns the other argument if it is not zero,
// otherwise it returns the existing argument.
func WithNumber[T Number](existing, other T) (result T) { //nolint:ireturn
	if other == 0 {
		return existing
	}
	return other
}

// WithPointer returns the existing argument if the other argument
// is nil. Otherwise it returns a new pointer to the copied value
// of the other argument value, for added mutation safety.
// For interfaces where the underlying type is not known,
// use WithInterface instead.
func WithPointer[T any](existing, other *T) (result *T) {
	if other == nil {
		return existing
	}
	result = new(T)
	*result = *other
	return result
}

// WithInterface returns the other argument if it is not nil,
// otherwise it returns the existing argument.
// Note you should NOT use this function with concrete pointers
// such as *int, and only use this for interfaces.
// This function is not type safe nor mutation safe, be careful.
func WithInterface(existing, other any) (result any) {
	if other == nil {
		return existing
	}
	return other
}

// WithString returns the other string argument if it is not empty,
// otherwise it returns the existing string argument.
func WithString(existing, other string) (result string) {
	if other == "" {
		return existing
	}
	return other
}

// WithCopiedSlice returns the existing slice argument if the other
// slice argument is nil. Otherwise it returns a new slice with the
// copied values of the other slice argument.
// Note it is preferrable to use this function for added mutation safety
// on the result, but one can use WithSlice if performance matters.
func WithCopiedSlice[T any](existing, other []T) (result []T) {
	if other == nil {
		return existing
	}
	result = make([]T, len(other))
	copy(result, other)
	return result
}

// WithSlice returns the other slice argument if it is not nil,
// otherwise it returns the existing slice argument.
func WithSlice[T any](existing, other []T) (result []T) {
	if other == nil {
		return existing
	}
	return other
}

// SelfValidator is an interface for a type that can validate itself.
// This is notably the case of netip.IP and netip.Prefix, and can be
// implemented by the user of this library as well.
type SelfValidator interface {
	// IsValid returns true if the value is valid, false otherwise.
	IsValid() bool
}

// WithValidator returns the existing argument if other is not valid,
// otherwise it returns the other argument.
func WithValidator(existing, other SelfValidator) ( //nolint:ireturn
	result SelfValidator) {
	if !other.IsValid() {
		return existing
	}
	return other
}
