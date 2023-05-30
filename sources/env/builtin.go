package env

import (
	"os"
	"strconv"
	"strings"

	"github.com/qdm12/govalid/binary"
	"github.com/qdm12/govalid/integer"
)

// Get returns an environment variable value modified
// depending on the options given.
// By default, and unless an option specifies otherwise,
// the following options are applied on the value string:
// - Trim line endings suffixes \r\n and \n.
// - Trim spaces.
// - Trim quotes.
// - Force lowercase.
func Get(envKey string, options ...Option) (value string) {
	settings := settingsFromOptions(options)

	value = os.Getenv(envKey)

	return postProcessValue(value, settings)
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

// CSV returns a slice of strings from a comma separated
// environment variable value. If the value is the empty
// string, `nil` is returned.
func CSV(envKey string, options ...Option) (values []string) {
	csv := Get(envKey, options...)
	if csv == "" {
		return nil
	}
	return strings.Split(csv, ",")
}

// Int returns an `int` from an environment variable value.
// If the value is the empty string, `0` is returned.
// Otherwise, if the value is not a valid integer string, an
// error is returned.
func Int(envKey string, options ...Option) (n int, err error) {
	s := Get(envKey, options...)
	if s == "" {
		return 0, nil
	}
	return strconv.Atoi(s)
}

// Float64 returns a `float64` from an environment variable value.
// If the value is the empty string, `0` is returned.
// Otherwise, if the value is not a valid float64 string, an error is returned.
func Float64(envKey string, options ...Option) (f float64, err error) {
	s := Get(envKey, options...)
	if s == "" {
		return 0, nil
	}
	const bits = 64
	return strconv.ParseFloat(s, bits)
}

// StringPtr returns a pointer to a `string` from an environment variable value.
// If the value is the empty string, `nil` is returned.
func StringPtr(envKey string, options ...Option) (stringPtr *string) {
	s := Get(envKey, options...)
	if s == "" {
		return nil
	}
	return &s
}

// BoolPtr returns a pointer to a `bool` from an environment variable value.
// 'true' string values are: "enabled", "yes", "on".
// 'false' string values are: "disabled", "no", "off".
// If the value is the empty string, `nil` is returned.
// Otherwise, if the value is not one of the above, an error is returned.
func BoolPtr(envKey string, options ...Option) (boolPtr *bool, err error) {
	s := Get(envKey, options...)
	value, err := binary.Validate(s)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// IntPtr returns a pointer to an `int` from an environment variable value.
// If the value is the empty string, `nil` is returned.
// Otherwise, if the value is not a valid integer string, an error is returned.
func IntPtr(envKey string, options ...Option) (intPtr *int, err error) {
	s := Get(envKey, options...)
	if s == "" {
		return nil, nil //nolint:nilnil
	}
	value, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

// Uint8Ptr returns a pointer to an `uint8` from an environment variable value.
// If the value is the empty string, `nil` is returned.
// Otherwise, if the value is not a valid integer string between 0 and 255,
// an error is returned.
func Uint8Ptr(envKey string, options ...Option) (uint8Ptr *uint8, err error) {
	s := Get(envKey, options...)
	if s == "" {
		return nil, nil //nolint:nilnil
	}

	const min, max = 0, 255
	value, err := integer.Validate(s, integer.OptionRange(min, max))
	if err != nil {
		return nil, err
	}

	uint8Ptr = new(uint8)
	*uint8Ptr = uint8(value)
	return uint8Ptr, nil
}

// Uint16Ptr returns a pointer to an `uint16` from an environment variable value.
// If the value is the empty string, `nil` is returned.
// Otherwise, if the value is not a valid integer string between 0 and 65535,
// an error is returned.
func Uint16Ptr(envKey string, options ...Option) (
	uint16Ptr *uint16, err error) {
	s := Get(envKey, options...)
	if s == "" {
		return nil, nil //nolint:nilnil
	}

	const min, max = 0, 65535
	value, err := integer.Validate(s, integer.OptionRange(min, max))
	if err != nil {
		return nil, err
	}

	uint16Ptr = new(uint16)
	*uint16Ptr = uint16(value)
	return uint16Ptr, nil
}
