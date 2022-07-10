package envhelpers

import (
	"os"
	"strconv"
	"strings"

	"github.com/qdm12/govalid/binary"
	"github.com/qdm12/govalid/integer"
)

// Get returns an environment variable value modified
// depending on the options given.
func Get(envKey string, options ...Option) (value string) {
	settings := settingsFromOptions(options)

	value = os.Getenv(envKey)

	if *settings.forceLowercase {
		value = strings.ToLower(value)
	}

	if *settings.trimSpace {
		value = strings.TrimSpace(value)
	}

	if *settings.trimLineEndings {
		for strings.HasSuffix(value, "\n") {
			value = strings.TrimSuffix(value, "\r\n")
			value = strings.TrimSuffix(value, "\n")
		}
	}

	return value
}

func CSV(envKey string, options ...Option) (values []string) {
	csv := Get(envKey, options...)
	if csv == "" {
		return nil
	}
	return strings.Split(csv, ",")
}

func Int(envKey string, options ...Option) (n int, err error) {
	s := Get(envKey, options...)
	if s == "" {
		return 0, nil
	}
	return strconv.Atoi(s)
}

func Float64(envKey string, options ...Option) (f float64, err error) {
	s := Get(envKey, options...)
	if s == "" {
		return 0, nil
	}
	const bits = 64
	return strconv.ParseFloat(s, bits)
}

func StringPtr(envKey string, options ...Option) (stringPtr *string) {
	s := Get(envKey, options...)
	if s == "" {
		return nil
	}
	return &s
}

func BoolPtr(envKey string, options ...Option) (boolPtr *bool, err error) {
	s := Get(envKey, options...)
	if s == "" {
		return nil, nil //nolint:nilnil
	}
	value, err := binary.Validate(s)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

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
