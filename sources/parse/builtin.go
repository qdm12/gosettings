package parse

import (
	"fmt"
	"strings"
)

// String returns the string found at the given key in the given
// keyValues map, which may be modified depending on the parse default
// settings and the parse options given. The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
//
// If the key is not set in the keyValues map, the empty string is
// returned.
func String(keyValues map[string]string, key string,
	options ...Option) (value string) {
	s := Get(keyValues, key, options...)
	if s == nil {
		return ""
	}
	return *s
}

// CSV returns a slice of strings from a comma separated
// value found at the given key in the given keyValues map.
// The entire CSV string value may be modified depending on the
// parse default settings and the parse options given.
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
//
// The slice is returned as `nil` if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func CSV(keyValues map[string]string, key string,
	options ...Option) (values []string) {
	csv := Get(keyValues, key, options...)
	if csv == nil {
		return nil
	}
	return strings.Split(*csv, ",")
}

// GetParse parses the value found at the given key in the given
// keyValues map using the given parse function, and returns the
// typed parsed value with an eventual error.
//
// The value is returned as the empty `T` value if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AcceptEmpty option, if the
//     key is set in the mapping and its corresponding value is empty.
//
// Note the value from the map may be modified depending on the parse
// default settings and the parse options given.
//
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func GetParse[T any](keyValues map[string]string, key string, //nolint:ireturn
	parse ParseFunc[T], options ...Option) (value T, err error) {
	s := Get(keyValues, key, options...)
	if s == nil {
		return value, nil
	}

	value, err = parse(*s)
	if err != nil {
		return value, fmt.Errorf("key %s: %w", key, err)
	}

	return value, nil
}

// GetParsePtr parses the value found at the given key in the given
// keyValues map using the given parse function, and returns the
// typed pointer to the parsed value with an eventual error.
//
// The value is returned as `nil` if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AcceptEmpty option, if the
//     key is set in the mapping and its corresponding value is empty.
//
// Note the value from the map may be modified depending on the parse
// default settings and the parse options given.
//
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func GetParsePtr[T any](keyValues map[string]string, key string,
	parse ParseFunc[T], options ...Option) (value *T, err error) {
	s := Get(keyValues, key, options...)
	if s == nil {
		return nil, nil //nolint:nilnil
	}

	value = new(T)
	*value, err = parse(*s)
	if err != nil {
		return nil, fmt.Errorf("key %s: %w", key, err)
	}

	return value, nil
}

// Int returns an `int` from the value found at the given key
// in the given keyValues map.
// If the value is not a valid integer string, an error is
// returned with the key in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Int(keyValues map[string]string, key string,
	options ...Option) (n int, err error) {
	return GetParse(keyValues, key, parseInt, options...)
}

// Float64 returns a `float64` from the value found at the given key
// in the given keyValues map.
// If the value is not a valid float64 string, an error is returned
// with the key in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Float64(keyValues map[string]string, key string,
	options ...Option) (f float64, err error) {
	return GetParse(keyValues, key, parseFloat64, options...)
}

// BoolPtr returns a pointer to a `bool` from the value found at the
// given key in the given keyValues map.
//   - 'true' string values are: "enabled", "yes", "on", "true".
//   - 'false' string values are: "disabled", "no", "off", "false".
//
// The value is returned as `nil` if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
//
// Otherwise, if the value is not one of the above, an error is returned
// with the key name in its message.
func BoolPtr(keyValues map[string]string, key string, options ...Option) (
	boolPtr *bool, err error) {
	return GetParse(keyValues, key, parseBool, options...)
}

// IntPtr returns a pointer to an `int` from the value found at the
// given key in the given keyValues map.
// If the value is not a valid integer string, an error is returned
// with the key name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func IntPtr(keyValues map[string]string, key string, options ...Option) (
	intPtr *int, err error) {
	return GetParsePtr(keyValues, key, parseInt, options...)
}

// Uint8Ptr returns a pointer to an `uint8` from the value found at the
// given key in the given keyValues map.
// If the value is not a valid integer string between 0 and 255,
// an error is returned with the key name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Uint8Ptr(keyValues map[string]string, key string, options ...Option) (
	uint8Ptr *uint8, err error) {
	return GetParsePtr(keyValues, key, parseUint8, options...)
}

// Uint16Ptr returns a pointer to an `uint16` from the value found at the
// given key in the given keyValues map.
// If the value is not a valid integer string between 0 and 65535,
// an error is returned with the key name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Uint16Ptr(keyValues map[string]string, key string, options ...Option) (
	uint16Ptr *uint16, err error) {
	return GetParsePtr(keyValues, key, parseUint16, options...)
}

// Uint32Ptr returns a pointer to an `uint32` from the value found at the
// given key in the given keyValues map.
// If the value is not a valid integer string between 0 and 4294967295,
// an error is returned with the key name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Uint32Ptr(keyValues map[string]string, key string, options ...Option) (
	uint32Ptr *uint32, err error) {
	return GetParsePtr(keyValues, key, parseUint32, options...)
}
