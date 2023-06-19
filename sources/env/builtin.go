package env

import (
	"fmt"
	"strconv"
	"strings"
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
	settings := settingsFromOptions(options)

	keysToTry := make([]string, 0, 1+len(settings.retroKeys))
	keysToTry = append(keysToTry, settings.retroKeys...)
	// Note we try the current environment variable key last
	// because it might be set in a Docker image so we want to
	// take the older configuration from the user first.
	keysToTry = append(keysToTry, envKey)

	var firstEnvKeySet string
	for _, keyToTry := range keysToTry {
		envValue, isSet := e.environ[keyToTry]
		if !isSet {
			continue
		}
		firstEnvKeySet = envKey
		value = new(string)
		*value = envValue
		break
	}

	if firstEnvKeySet == "" { // All keys are unset
		return nil
	}

	if firstEnvKeySet != envKey {
		e.handleDeprecatedKey(firstEnvKeySet, envKey)
	}

	if !*settings.acceptEmpty && *value == "" {
		// environment variable value is set to the empty string,
		// but the empty string is not accepted so return nil.
		return nil
	}

	*value = postProcessValue(*value, settings)
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
	s := e.Get(envKey, options...)
	if s == nil {
		return ""
	}
	return *s
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
	csv := e.Get(envKey, options...)
	if csv == nil {
		return nil
	}
	return strings.Split(*csv, ",")
}

// Int returns an `int` from an environment variable value.
// If the value is not a valid integer string, an error is
// returned with the environment variable name in its message.
// The value is returned as `0` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
func (e *Env) Int(envKey string, options ...Option) (n int, err error) {
	s := e.Get(envKey, options...)
	if s == nil {
		return 0, nil
	}

	n, err = parseInt(*s)
	if err != nil {
		return 0, fmt.Errorf("environment variable %s: %w", envKey, err)
	}

	return n, nil
}

// Float64 returns a `float64` from an environment variable value.
// If the value is not a valid float64 string, an error is returned
// with the environment variable name in its message.
// The value is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
func (e *Env) Float64(envKey string, options ...Option) (f float64, err error) {
	s := e.Get(envKey, options...)
	if s == nil {
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
	value := e.Get(envKey, options...)
	if value == nil {
		return nil, nil //nolint:nilnil
	}

	boolPtr, err = parseBool(*value)
	if err != nil {
		return nil, fmt.Errorf("environment variable %s: %w", envKey, err)
	}
	return boolPtr, nil
}

// IntPtr returns a pointer to an `int` from an environment variable value.
// If the value is not a valid integer string, an error is returned
// with the environment variable name in its message.
// The value is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
func (e *Env) IntPtr(envKey string, options ...Option) (intPtr *int, err error) {
	s := e.Get(envKey, options...)
	if s == nil {
		return nil, nil //nolint:nilnil
	}
	value, err := parseInt(*s)
	if err != nil {
		return nil, fmt.Errorf("environment variable %s: %w", envKey, err)
	}
	return &value, nil
}

// Uint8Ptr returns a pointer to an `uint8` from an environment variable value.
// If the value is not a valid integer string between 0 and 255,
// an error is returned with the environment variable name in its message.
// The value is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
func (e *Env) Uint8Ptr(envKey string, options ...Option) (uint8Ptr *uint8, err error) {
	s := e.Get(envKey, options...)
	if s == nil {
		return nil, nil //nolint:nilnil
	}

	value, err := parseUint8(*s)
	if err != nil {
		return nil, fmt.Errorf("environment variable %s: %w", envKey, err)
	}

	return &value, nil
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
	s := e.Get(envKey, options...)
	if s == nil {
		return nil, nil //nolint:nilnil
	}

	value, err := parseUint16(*s)
	if err != nil {
		return nil, fmt.Errorf("environment variable %s: %w", envKey, err)
	}

	return &value, nil
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
	s := e.Get(envKey, options...)
	if s == nil {
		return nil, nil //nolint:nilnil
	}

	value, err := parseUint32(*s)
	if err != nil {
		return nil, fmt.Errorf("environment variable %s: %w", envKey, err)
	}

	return &value, nil
}
