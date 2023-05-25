//go:build go1.18
// +build go1.18

package gosettings

// MergeWithNumber returns the existing argument if it is not zero,
// otherwise it returns the other argument.
func MergeWithNumber[T Number](existing, other T) (result T) { //nolint:ireturn
	if existing != 0 {
		return existing
	}
	return other
}

// MergeWithPointer returns the existing argument if it is not nil.
// Otherwise it returns a new pointer to the copied value of the
// other argument value, for added mutation safety.
// For interfaces where the underlying type is not known,
// use MergeWithInterface instead.
func MergeWithPointer[T any](existing, other *T) (result *T) {
	if existing != nil {
		return existing
	}
	result = new(T)
	*result = *other
	return result
}

// MergeWithInterface returns the existing argument if it is not nil,
// otherwise it returns the other argument.
// Note you should NOT use this function with concrete pointers
// such as *int, and only use this for interfaces.
// This function is not type safe nor mutation safe, be careful.
func MergeWithInterface(existing, other any) (result any) {
	if existing != nil {
		return existing
	}
	return other
}

// MergeWithString returns the existing string argument if it is not empty,
// otherwise it returns the other string argument.
func MergeWithString(existing, other string) (result string) {
	if existing != "" {
		return existing
	}
	return other
}

// MergeWithSlice returns the existing slice argument if is not nil.
// Otherwise it returns a new slice with the copied values of the
// other slice argument.
// Note it is preferrable to use this function for added mutation safety
// on the result, but one can use MergeWithSliceRaw if performance matters.
func MergeWithSlice[T any](existing, other []T) (result []T) {
	if existing != nil {
		return existing
	}
	result = make([]T, len(other))
	copy(result, other)
	return result
}

// MergeWithSliceRaw returns the existing slice argument if it is not nil,
// otherwise it returns the other slice argument.
func MergeWithSliceRaw[T any](existing, other []T) (result []T) {
	if existing != nil {
		return existing
	}
	return other
}

// MergeWithValidator returns the existing argument if it is valid,
// otherwise it returns the defaultValue argument.
func MergeWithValidator[T SelfValidator](existing, defaultValue T) ( //nolint:ireturn
	result T) {
	if existing.IsValid() {
		return existing
	}
	return defaultValue
}
