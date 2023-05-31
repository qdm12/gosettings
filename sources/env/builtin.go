package env

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/qdm12/govalid/binary"
	"github.com/qdm12/govalid/integer"
)

// Get returns an environment variable value as a string pointer.
// The pointer is returned as `nil` if the environment variable key
// given is NOT set. By default and unless changed by the AllowEmpty
// option, if the environment variable is set but its value is empty,
// the pointer is returned as `nil`.
// Otherwise, the value string is then modified depending on the
// defaults settings and options given, and returned as a pointer.
// By default, and unless an option specifies otherwise,
// the following options are applied on the value string:
// - Trim line endings suffixes \r\n and \n.
// - Trim spaces.
// - Trim quotes.
// - Force lowercase.
func (e *Env) Get(envKey string, options ...Option) (value *string) {
	settings := settingsFromOptions(options)

	envValue, isSet := e.environ[envKey]
	if !isSet || (!*settings.acceptEmpty && envValue == "") {
		return nil
	}

	value = new(string)
	*value = postProcessValue(envValue, settings)
	return value
}

func postProcessValue(value string, settings settings) string {
	if *settings.forceLowercase {
		value = strings.ToLower(value)
	}

	cutSet := map[string]struct{}{}
	if *settings.trimSpace {
		// Only latin charset
		spaceCharacters := []rune{'\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0}
		for _, r := range spaceCharacters {
			cutSet[string(r)] = struct{}{}
		}
	}

	if *settings.trimLineEndings {
		cutSet["\r"] = struct{}{}
		cutSet["\n"] = struct{}{}
	}

	if *settings.trimQuotes {
		cutSet[`"`] = struct{}{}
		cutSet[`'`] = struct{}{}
	}

	cutSetString := ""
	for s := range cutSet {
		cutSetString += s
	}

	return strings.Trim(value, cutSetString)
}

// String returns a string from an environment variable value.
// If the environment variable is not set, the empty string is
// returned.
// By default, and unless an option specifies otherwise,
// the following options are applied on the value string:
// - Trim line endings suffixes \r\n and \n.
// - Trim spaces.
// - Trim quotes.
// - Force lowercase.
func (e *Env) String(envKey string, options ...Option) (value string) {
	s := e.Get(envKey, options...)
	if s == nil {
		return ""
	}
	return *s
}

// CSV returns a slice of strings from a comma separated
// environment variable value.
// If the environment variable is not set, `nil` is returned.
// By default and unless specified by an option, if the
// environment variable is set but its value is empty, `nil`
// is returned.
// Otherwise, the value is split by commas and returned as a
// slice of strings.
func (e *Env) CSV(envKey string, options ...Option) (values []string) {
	csv := e.Get(envKey, options...)
	if csv == nil {
		return nil
	}
	return strings.Split(*csv, ",")
}

// Int returns an `int` from an environment variable value.
// If the environment variable is not set or its value is
// the empty string, `0` is returned.
// Otherwise, if the value is not a valid integer string, an
// error is returned with the environment variable name in the
// error context.
func (e *Env) Int(envKey string, options ...Option) (n int, err error) {
	s := e.Get(envKey, options...)
	if s == nil || *s == "" {
		// note: no point accepting the empty string in this case
		return 0, nil
	}

	n, err = strconv.Atoi(*s)
	if err != nil {
		return 0, fmt.Errorf("environment variable %s: %w", envKey, err)
	}

	return n, nil
}

// Float64 returns a `float64` from an environment variable value.
// If the environment variable is not set or its value is
// the empty string, `0` is returned.
// Otherwise, if the value is not a valid float64 string, an error is
// returned with the environment variable name in the error context.
func (e *Env) Float64(envKey string, options ...Option) (f float64, err error) {
	s := e.Get(envKey, options...)
	if s == nil || *s == "" {
		// note: no point accepting the empty string in this case
		return 0, nil
	}

	const bits = 64
	f, err = strconv.ParseFloat(*s, bits)
	if err != nil {
		return 0, fmt.Errorf("environment variable %s: %w", envKey, err)
	}

	return f, nil
}

// BoolPtr returns a pointer to a `bool` from an environment variable value.
// 'true' string values are: "enabled", "yes", "on".
// 'false' string values are: "disabled", "no", "off".
// If the environment variable is not set or its value is the empty string,
// `nil` is returned.
// Otherwise, if the value is not one of the above, an error is returned
// with the environment variable name in the error context.
func (e *Env) BoolPtr(envKey string, options ...Option) (boolPtr *bool, err error) {
	s := e.Get(envKey, options...)
	if s == nil || *s == "" {
		// note: no point accepting the empty string in this case
		return nil, nil //nolint:nilnil
	}

	value, err := binary.Validate(*s)
	if err != nil {
		return nil, fmt.Errorf("environment variable %s: %w", envKey, err)
	}
	return value, nil
}

// IntPtr returns a pointer to an `int` from an environment variable value.
// If the environment variable is not set or its value is the empty string,
// `nil` is returned.
// Otherwise, if the value is not a valid integer string, an error is returned
// with the environment variable name in the error context.
func (e *Env) IntPtr(envKey string, options ...Option) (intPtr *int, err error) {
	s := e.Get(envKey, options...)
	if s == nil || *s == "" {
		// note: no point accepting the empty string in this case
		return nil, nil //nolint:nilnil
	}
	value, err := strconv.Atoi(*s)
	if err != nil {
		return nil, fmt.Errorf("environment variable %s: %w", envKey, err)
	}
	return &value, nil
}

// Uint8Ptr returns a pointer to an `uint8` from an environment variable value.
// If the environment variable is not set or its value is the empty string,
// `nil` is returned.
// Otherwise, if the value is not a valid integer string between 0 and 255,
// an error is returned with the environment variable name in the error context.
func (e *Env) Uint8Ptr(envKey string, options ...Option) (uint8Ptr *uint8, err error) {
	s := e.Get(envKey, options...)
	if s == nil || *s == "" {
		// note: no point accepting the empty string in this case
		return nil, nil //nolint:nilnil
	}

	const min, max = 0, 255
	value, err := integer.Validate(*s, integer.OptionRange(min, max))
	if err != nil {
		return nil, fmt.Errorf("environment variable %s: %w", envKey, err)
	}

	uint8Ptr = new(uint8)
	*uint8Ptr = uint8(value)
	return uint8Ptr, nil
}

// Uint16Ptr returns a pointer to an `uint16` from an environment variable value.
// If the environment variable is not set or its value is the empty string,
// `nil` is returned.
// Otherwise, if the value is not a valid integer string between 0 and 65535,
// an error is returned with the environment variable name in the error context.
func (e *Env) Uint16Ptr(envKey string, options ...Option) (
	uint16Ptr *uint16, err error) {
	s := e.Get(envKey, options...)
	if s == nil || *s == "" {
		// note: no point accepting the empty string in this case
		return nil, nil //nolint:nilnil
	}

	const min, max = 0, 65535
	value, err := integer.Validate(*s, integer.OptionRange(min, max))
	if err != nil {
		return nil, fmt.Errorf("environment variable %s: %w", envKey, err)
	}

	uint16Ptr = new(uint16)
	*uint16Ptr = uint16(value)
	return uint16Ptr, nil
}
