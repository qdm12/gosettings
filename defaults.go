//go:build go1.18
// +build go1.18

package gosettings

// DefaultNumber returns the existing argument if it is not zero,
// otherwise it returns the other argument.
func DefaultNumber[T Number](existing, other T) (result T) { //nolint:ireturn
	return MergeWithNumber(existing, other)
}

// DefaultPointer returns the existing argument if it is not nil.
// Otherwise it returns a new pointer to the defaultValue argument.
// To default an interface to an implementation, use DefaultPointerRaw.
func DefaultPointer[T any](existing *T, defaultValue T) (result *T) {
	if existing != nil {
		return existing
	}
	result = new(T)
	*result = defaultValue
	return result
}

// DefaultPointerRaw returns the existing argument if it is not nil.
// Otherwise it returns the defaultValue pointer argument.
// For interfaces where the underlying type is not known,
// use DefaultInterface instead.
func DefaultPointerRaw[T any](existing, defaultValue *T) (result *T) {
	if existing != nil {
		return existing
	}
	return defaultValue
}

// DefaultInterface returns the existing argument if it is not nil,
// otherwise it returns the defaultValue argument.
// If used with an interface and an implementation of the interface,
// it must be instantiated with the interface type, for example:
// variable := DefaultInterface[Interface](variable, &implementation{})
// Avoid using this function for non-interface types.
func DefaultInterface[T comparable](existing, defaultValue T) ( //nolint:ireturn
	result T) {
	var zero T
	if existing != zero {
		return existing
	}
	return defaultValue
}

// DefaultString returns the existing string argument if it is not empty,
// otherwise it returns the defaultValue string argument.
func DefaultString(existing, defaultValue string) (result string) {
	return MergeWithString(existing, defaultValue)
}

// DefaultSlice returns the existing slice argument if is not nil.
// Otherwise it returns a new slice with the copied values of the
// defaultValue slice argument.
// Note it is preferrable to use this function for added mutation safety
// on the result, but one can use DefaultSliceRaw if performance matters.
func DefaultSlice[T any](existing, other []T) (result []T) {
	return MergeWithSlice(existing, other)
}

// DefaultSliceRaw returns the existing slice argument if it is not nil,
// otherwise it returns the defaultValue slice argument.
func DefaultSliceRaw[T any](existing, defaultValue []T) (result []T) {
	return MergeWithSliceRaw(existing, defaultValue)
}

// DefaultValidator returns the existing argument if it is valid,
// otherwise it returns the defaultValue argument.
func DefaultValidator[T SelfValidator](existing, defaultValue T) ( //nolint:ireturn
	result T) {
	return MergeWithValidator(existing, defaultValue)
}
