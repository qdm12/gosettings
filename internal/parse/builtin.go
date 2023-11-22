package parse

import (
	"fmt"
	"strings"
)

// String returns the first string value found at the
// given key from the given sources in order.
// The value may be modified depending on the parse default settings
// and the parse options given. The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
//
// If the key is not set in any of the sources, the empty string is
// returned.
func String(sources []Source, key string,
	options ...Option) (value string) {
	s, _ := get(sources, key, options...)
	if s == nil {
		return ""
	}
	return *s
}

// CSV returns a slice of strings from the first comma separated
// value found from the given sources in order.
// The entire CSV string value may be modified depending on the
// parse default settings and the parse options given.
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
//
// The slice is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func CSV(sources []Source, key string,
	options ...Option) (values []string) {
	csv, _ := csv(sources, key, options...)
	return csv
}

func csv(sources []Source, key string,
	options ...Option) (values []string, sourceName string) {
	csv, sourceName := get(sources, key, options...)
	if csv == nil {
		return nil, sourceName
	}
	return strings.Split(*csv, ","), sourceName
}

// GetParse parses the first value found at the given key
// from the given sources in order, using the given parse
// function, and returns the typed parsed value with an
// eventual error containing the key name and the source name
// in its message.
//
// The value is returned as the empty `T` value if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AcceptEmpty option, if the
//     key is set in a source and its corresponding value is empty.
//
// Note the value from the map may be modified depending on the parse
// default settings and the parse options given.
//
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func GetParse[T any](sources []Source, key string, //nolint:ireturn
	parse ParseFunc[T], options ...Option) (value T, err error) {
	s, sourceKind := get(sources, key, options...)
	if s == nil {
		return value, nil
	}

	value, err = parse(*s)
	if err != nil {
		return value, fmt.Errorf("%s %s: %w", sourceKind, key, err)
	}

	return value, nil
}

// GetParsePtr parses the first value found at the given key from
// the given sources in order, using the given parse function, and
// returns the typed pointer to the parsed value with an eventual
// error which contains the key name and source name in its message.
//
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
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
func GetParsePtr[T any](sources []Source, key string,
	parse ParseFunc[T], options ...Option) (value *T, err error) {
	s, sourceKind := get(sources, key, options...)
	if s == nil {
		return nil, nil //nolint:nilnil
	}

	value = new(T)
	*value, err = parse(*s)
	if err != nil {
		return nil, fmt.Errorf("%s %s: %w", sourceKind, key, err)
	}

	return value, nil
}

// Int returns an `int` from the first value found at the given
// key in the given sources in order.
// If the value is not a valid integer string, an error is
// returned with the key name and the source name in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Int(sources []Source, key string,
	options ...Option) (n int, err error) {
	return GetParse(sources, key, parseInt, options...)
}

// Float64 returns a `float64` from the first value found at the given
// key in the given sources in order.
// If the value is not a valid float64 string, an error is
// returned with the key name and the source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option,
//     if the key is set and its corresponding value is empty.
func Float64(sources []Source, key string,
	options ...Option) (f float64, err error) {
	return GetParse(sources, key, parseFloat64, options...)
}

// BoolPtr returns a pointer to a `bool` from the first value found
// at the given key in the given sources in order.
//   - 'true' string values are: "enabled", "yes", "on", "true".
//   - 'false' string values are: "disabled", "no", "off", "false".
//
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
//
// Otherwise, if the value is not one of the above, an error is returned
// with the key name and source name in its message.
func BoolPtr(sources []Source, key string, options ...Option) (
	boolPtr *bool, err error) {
	return GetParse(sources, key, parseBool, options...)
}

// IntPtr returns a pointer to an `int` from the first value found
// at the given key in the given sources in order.
// If the value is not a valid integer string, an error is returned
// with the key name and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func IntPtr(sources []Source, key string, options ...Option) (
	intPtr *int, err error) {
	return GetParsePtr(sources, key, parseInt, options...)
}

// Uint8Ptr returns a pointer to an `uint8` from the first value found
// at the given key in the given sources in order.
// If the value is not a valid integer string between 0 and 255,
// an error is returned with the key name and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Uint8Ptr(sources []Source, key string, options ...Option) (
	uint8Ptr *uint8, err error) {
	return GetParsePtr(sources, key, parseUint8, options...)
}

// Uint16Ptr returns a pointer to an `uint16` from the first value found
// at the given key in the given sources in order.
// If the value is not a valid integer string between 0 and 65535,
// an error is returned with the key name and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Uint16Ptr(sources []Source, key string, options ...Option) (
	uint16Ptr *uint16, err error) {
	return GetParsePtr(sources, key, parseUint16, options...)
}

// Uint32Ptr returns a pointer to an `uint32` from the first value found
// at the given key in the given sources in order.
// If the value is not a valid integer string between 0 and 4294967295,
// an error is returned with the key name and source name in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in any of the sources.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Uint32Ptr(sources []Source, key string, options ...Option) (
	uint32Ptr *uint32, err error) {
	return GetParsePtr(sources, key, parseUint32, options...)
}
