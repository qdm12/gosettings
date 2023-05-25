//go:build go1.18
// +build go1.18

package gosettings

// WithNumber returns the other argument if it is not zero,
// otherwise it returns the existing argument.
func WithNumber[T Number](existing, other T) (result T) { //nolint:ireturn
	if other == 0 {
		return existing
	}
	return other
}

// OverrideWithPointer returns the existing argument if the other argument
// is nil. Otherwise it returns a new pointer to the copied value
// of the other argument value, for added mutation safety.
// For interfaces where the underlying type is not known,
// use OverrideWithInterface instead.
func OverrideWithPointer[T any](existing, other *T) (result *T) {
	if other == nil {
		return existing
	}
	result = new(T)
	*result = *other
	return result
}

// OverrideWithInterface returns the other argument if it is not nil,
// otherwise it returns the existing argument.
// Note you should NOT use this function with concrete pointers
// such as *int, and only use this for interfaces.
// This function is not type safe nor mutation safe, be careful.
func OverrideWithInterface(existing, other any) (result any) {
	if other == nil {
		return existing
	}
	return other
}

// OverrideWithString returns the other string argument if it is not empty,
// otherwise it returns the existing string argument.
func OverrideWithString(existing, other string) (result string) {
	if other == "" {
		return existing
	}
	return other
}

// OverrideWithSlice returns the existing slice argument if the other
// slice argument is nil. Otherwise it returns a new slice with the
// copied values of the other slice argument.
// Note it is preferrable to use this function for added mutation safety
// on the result, but one can use OverrideWithSliceRaw if performance matters.
func OverrideWithSlice[T any](existing, other []T) (result []T) {
	if other == nil {
		return existing
	}
	result = make([]T, len(other))
	copy(result, other)
	return result
}

// OverrideWithSliceRaw returns the other slice argument if it is not nil,
// otherwise it returns the existing slice argument.
func OverrideWithSliceRaw[T any](existing, other []T) (result []T) {
	if other == nil {
		return existing
	}
	return other
}

// OverrideWithValidator returns the existing argument if other is not valid,
// otherwise it returns the other argument.
func OverrideWithValidator[T SelfValidator](existing, other T) ( //nolint:ireturn
	result T) {
	if !other.IsValid() {
		return existing
	}
	return other
}
