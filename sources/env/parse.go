package env

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/qdm12/gosettings/validate"
	"golang.org/x/exp/constraints"
)

var (
	ErrValueNotInRange = errors.New("value is not in range")
)

func ptrTo[T any](x T) *T { return &x }

func parseBool(value string) (output *bool, err error) {
	if value == "" {
		return nil, nil //nolint:nilnil
	}

	lowercasedValue := strings.ToLower(value)
	enabledStrings := []string{"enabled", "yes", "on", "true"}
	disabledStrings := []string{"disabled", "no", "off", "false"}
	for _, enabledString := range enabledStrings {
		if lowercasedValue == enabledString {
			return ptrTo(true), nil
		}
	}

	for _, disabledString := range disabledStrings {
		if lowercasedValue == disabledString {
			return ptrTo(false), nil
		}
	}

	possibilities := make([]string, len(enabledStrings), len(enabledStrings)+len(disabledStrings))
	copy(possibilities, enabledStrings)
	possibilities = append(possibilities, disabledStrings...)
	err = validate.IsOneOf(lowercasedValue, possibilities...)
	return ptrTo(false), err
}

func parseInt(value string) (output int, err error) {
	const min, max = int64(math.MinInt), int64(math.MaxInt)
	return parseSigned[int](value, min, max)
}

func parseInt8(value string) (output int8, err error) {
	const min, max = math.MinInt8, math.MaxInt8
	return parseSigned[int8](value, min, max)
}

func parseInt16(value string) (output int16, err error) {
	const min, max = math.MinInt16, math.MaxInt16
	return parseSigned[int16](value, min, max)
}

func parseInt32(value string) (output int32, err error) {
	const min, max = math.MinInt32, math.MaxInt32
	return parseSigned[int32](value, min, max)
}

func parseInt64(value string) (output int64, err error) {
	const min, max = math.MinInt64, math.MaxInt64
	return parseSigned[int64](value, min, max)
}

func parseSigned[T constraints.Signed](value string, min, max int64) ( //nolint:ireturn
	n T, err error) {
	const base, bits = 10, 64
	xInt64, err := strconv.ParseInt(value, base, bits)
	if err != nil {
		return 0, err
	}
	if xInt64 < min || xInt64 > max {
		return 0, fmt.Errorf("%w: %d is not between %d and %d",
			ErrValueNotInRange, xInt64, min, max)
	}

	return T(xInt64), nil
}

func parseUint(value string) (output uint, err error) {
	const min, max = 0, math.MaxUint
	return parseUnsigned[uint](value, min, max)
}

func parseUint8(value string) (output uint8, err error) {
	const min, max = 0, math.MaxUint8
	return parseUnsigned[uint8](value, min, max)
}

func parseUint16(value string) (output uint16, err error) {
	const min, max = 0, math.MaxUint16
	return parseUnsigned[uint16](value, min, max)
}

func parseUint32(value string) (output uint32, err error) {
	const min, max = 0, math.MaxUint32
	return parseUnsigned[uint32](value, min, max)
}

func parseUint64(value string) (output uint64, err error) {
	const min, max = 0, math.MaxUint64
	return parseUnsigned[uint64](value, min, max)
}

func parseUnsigned[T constraints.Unsigned](value string, min, max uint64) ( //nolint:ireturn
	n T, err error) {
	const base, bits = 10, 64
	xUint64, err := strconv.ParseUint(value, base, bits)
	if err != nil {
		return 0, err
	}
	if xUint64 < min || xUint64 > max {
		return 0, fmt.Errorf("%w: %d is not between %d and %d",
			ErrValueNotInRange, xUint64, min, max)
	}

	return T(xUint64), nil
}
