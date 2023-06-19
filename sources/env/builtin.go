package env

import (
	"github.com/qdm12/gosettings/sources/parse"
)

// Get returns an environment variable value as a string pointer.
//
// The value is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option, if the
//     environment variable is set and its value is empty.
//
// Otherwise, the value may be modified depending on the parse
// default settings and the parse options given.
//
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func (e *Env) Get(envKey string, options ...Option) (value *string) {
	parseOptions := e.makeParseOptions(options)
	return parse.Get(e.environ, envKey, parseOptions...)
}

// String returns a string from an environment variable value,
// which may be modified depending on the parse default settings
// and the parse options given. The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
//
// If the environment variable is not set, the empty string is
// returned.
func (e *Env) String(envKey string, options ...Option) (value string) {
	parseOptions := e.makeParseOptions(options)
	return parse.String(e.environ, envKey, parseOptions...)
}

// CSV returns a slice of strings from a comma separated
// environment variable value.
// The entire CSV string value may be modified depending on the
// parse default settings and the parse options given.
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
//
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (e *Env) CSV(envKey string, options ...Option) (values []string) {
	parseOptions := e.makeParseOptions(options)
	return parse.CSV(e.environ, envKey, parseOptions...)
}

// Int returns an `int` from an environment variable value.
// If the value is not a valid integer string, an error is
// returned with the environment variable name in its message.
// The value is returned as `0` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
func (e *Env) Int(envKey string, options ...Option) (n int, err error) {
	parseOptions := e.makeParseOptions(options)
	return parse.Int(e.environ, envKey, parseOptions...)
}

// Float64 returns a `float64` from an environment variable value.
// If the value is not a valid float64 string, an error is returned
// with the environment variable name in its message.
// The value is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
func (e *Env) Float64(envKey string, options ...Option) (f float64, err error) {
	parseOptions := e.makeParseOptions(options)
	return parse.Float64(e.environ, envKey, parseOptions...)
}

// BoolPtr returns a pointer to a `bool` from an environment variable value.
//   - 'true' string values are: "enabled", "yes", "on", "true".
//   - 'false' string values are: "disabled", "no", "off", "false".
//
// The value is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
//
// Otherwise, if the value is not one of the above, an error is returned
// with the environment variable name in its message.
func (e *Env) BoolPtr(envKey string, options ...Option) (boolPtr *bool, err error) {
	parseOptions := e.makeParseOptions(options)
	return parse.BoolPtr(e.environ, envKey, parseOptions...)
}

// IntPtr returns a pointer to an `int` from an environment variable value.
// If the value is not a valid integer string, an error is returned
// with the environment variable name in its message.
// The value is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
func (e *Env) IntPtr(envKey string, options ...Option) (intPtr *int, err error) {
	parseOptions := e.makeParseOptions(options)
	return parse.IntPtr(e.environ, envKey, parseOptions...)
}

// Uint8Ptr returns a pointer to an `uint8` from an environment variable value.
// If the value is not a valid integer string between 0 and 255,
// an error is returned with the environment variable name in its message.
// The value is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
func (e *Env) Uint8Ptr(envKey string, options ...Option) (uint8Ptr *uint8, err error) {
	parseOptions := e.makeParseOptions(options)
	return parse.Uint8Ptr(e.environ, envKey, parseOptions...)
}

// Uint16Ptr returns a pointer to an `uint16` from an environment variable value.
// If the value is not a valid integer string between 0 and 65535,
// an error is returned with the environment variable name its message.
// The value is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
func (e *Env) Uint16Ptr(envKey string, options ...Option) (
	uint16Ptr *uint16, err error) {
	parseOptions := e.makeParseOptions(options)
	return parse.Uint16Ptr(e.environ, envKey, parseOptions...)
}

// Uint32Ptr returns a pointer to an `uint32` from an environment variable value.
// If the value is not a valid integer string between 0 and 4294967295
// an error is returned with the environment variable name in its message.
// The value is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
func (e *Env) Uint32Ptr(envKey string, options ...Option) (
	uint32Ptr *uint32, err error) {
	parseOptions := e.makeParseOptions(options)
	return parse.Uint32Ptr(e.environ, envKey, parseOptions...)
}
