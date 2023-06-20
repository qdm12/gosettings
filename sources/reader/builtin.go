package reader

import (
	"github.com/qdm12/gosettings/sources/parse"
)

// Get returns a value found at the given key as a string pointer.
//
// The value is returned as `nil` if:
//   - the key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option, if the
//     key is set and its corresponding value is empty.
//
// Otherwise, the value may be modified depending on the parse
// default settings and the parse options given.
//
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func (r *Reader) Get(key string, options ...Option) (value *string) {
	parseOptions := r.makeParseOptions(options)
	return parse.Get(r.keyToValue, key, parseOptions...)
}

// String returns a string from the value found at the given key,
// which may be modified depending on the parse default settings
// and the parse options given. The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
//
// If the key is not set, the empty string is returned.
func (r *Reader) String(key string, options ...Option) (value string) {
	parseOptions := r.makeParseOptions(options)
	return parse.String(r.keyToValue, key, parseOptions...)
}

// CSV returns a slice of strings from a comma separated value
// found at the given key.
// The entire CSV string value may be modified depending on the
// parse default settings and the parse options given.
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
//
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
func (r *Reader) CSV(key string, options ...Option) (values []string) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSV(r.keyToValue, key, parseOptions...)
}

// Int returns an `int` from the value found at the given key.
// If the value is not a valid integer string, an error is
// returned with the key in its message.
// The value is returned as `0` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Int(key string, options ...Option) (n int, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Int(r.keyToValue, key, parseOptions...)
}

// Float64 returns a `float64` from the value found at the given key.
// If the value is not a valid float64 string, an error is returned
// with the key in its message.
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Float64(key string, options ...Option) (f float64, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Float64(r.keyToValue, key, parseOptions...)
}

// BoolPtr returns a pointer to a `bool` from the value found at the given key.
//   - 'true' string values are: "enabled", "yes", "on", "true".
//   - 'false' string values are: "disabled", "no", "off", "false".
//
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
//
// Otherwise, if the value is not one of the above, an error is returned
// with the key in its message.
func (r *Reader) BoolPtr(key string, options ...Option) (boolPtr *bool, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.BoolPtr(r.keyToValue, key, parseOptions...)
}

// IntPtr returns a pointer to an `int` from the value found at the given key.
// If the value is not a valid integer string, an error is returned
// with the key in its message.
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) IntPtr(key string, options ...Option) (intPtr *int, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.IntPtr(r.keyToValue, key, parseOptions...)
}

// Uint8Ptr returns a pointer to an `uint8` from the value found at the given key.
// If the value is not a valid integer string between 0 and 255,
// an error is returned with the key in its message.
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Uint8Ptr(key string, options ...Option) (uint8Ptr *uint8, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Uint8Ptr(r.keyToValue, key, parseOptions...)
}

// Uint16Ptr returns a pointer to an `uint16` from the value found at the given key.
// If the value is not a valid integer string between 0 and 65535,
// an error is returned with the key its message.
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Uint16Ptr(key string, options ...Option) (
	uint16Ptr *uint16, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Uint16Ptr(r.keyToValue, key, parseOptions...)
}

// Uint32Ptr returns a pointer to an `uint32` from the value found at the given key.
// If the value is not a valid integer string between 0 and 4294967295
// an error is returned with the key in its message.
// The value is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) Uint32Ptr(key string, options ...Option) (
	uint32Ptr *uint32, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.Uint32Ptr(r.keyToValue, key, parseOptions...)
}
