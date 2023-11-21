//go:build go1.18
// +build go1.18

package gosettings

// MergeWithComparable returns the existing argument if it is
// not the zero value, otherwise it returns the other argument.
// If used with an interface and an implementation of the interface,
// it must be instantiated with the interface type, for example:
// variable := MergeWithComparable[Interface](variable, &implementation{})
// Avoid using this function for non-interface pointers, use MergeWithPointer
// instead to create a new pointer.
func MergeWithComparable[T comparable](existing, other T) (result T) { //nolint:ireturn
	var zero T
	if existing != zero {
		return existing
	}
	return other
}

// MergeWithPointer returns the existing argument if it is not nil.
// Otherwise it returns a new pointer to the copied value of the
// other argument value, for added mutation safety.
// To merge an interface and an implementation, use MergeWithComparable.
func MergeWithPointer[T any](existing, other *T) (result *T) {
	if existing != nil || other == nil {
		return existing
	}
	result = new(T)
	*result = *other
	return result
}

// MergeWithSlice returns the existing slice argument if is not nil.
// Otherwise it returns a new slice with the copied values of the
// other slice argument.
// Note it is preferrable to use this function for added mutation safety
// on the result, but one can use MergeWithSliceRaw if performance matters.
func MergeWithSlice[T any](existing, other []T) (result []T) {
	if existing != nil || other == nil {
		return existing
	}
	result = make([]T, len(other))
	copy(result, other)
	return result
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
