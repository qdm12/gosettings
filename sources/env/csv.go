package env

import (
	"fmt"
)

// CSVInt returns a slice of int from a comma separated
// environment variable value and returns an error if any value
// is not a valid int string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (e *Env) CSVInt(envKey string, options ...Option) (values []int, err error) {
	return csvParse(e, envKey, parseInt, options...)
}

// CSVInt8 returns a slice of int8 from a comma separated
// environment variable value and returns an error if any value
// is not a valid uint8 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (e *Env) CSVInt8(envKey string, options ...Option) (values []int8, err error) {
	return csvParse(e, envKey, parseInt8, options...)
}

// CSVInt16 returns a slice of int16 from a comma separated
// environment variable value and returns an error if any value
// is not a valid uint16 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (e *Env) CSVInt16(envKey string, options ...Option) (values []int16, err error) {
	return csvParse(e, envKey, parseInt16, options...)
}

// CSVInt32 returns a slice of int32 from a comma separated
// environment variable value and returns an error if any value
// is not a valid int32 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (e *Env) CSVInt32(envKey string, options ...Option) (values []int32, err error) {
	return csvParse(e, envKey, parseInt32, options...)
}

// CSVInt64 returns a slice of int64 from a comma separated
// environment variable value and returns an error if any value
// is not a valid int64 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (e *Env) CSVInt64(envKey string, options ...Option) (values []int64, err error) {
	return csvParse(e, envKey, parseInt64, options...)
}

// CSVUint returns a slice of uint from a comma separated
// environment variable value and returns an error if any value
// is not a valid uint string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (e *Env) CSVUint(envKey string, options ...Option) (values []uint, err error) {
	return csvParse(e, envKey, parseUint, options...)
}

// CSVUint8 returns a slice of uint8 from a comma separated
// environment variable value and returns an error if any value
// is not a valid uint8 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (e *Env) CSVUint8(envKey string, options ...Option) (values []uint8, err error) {
	return csvParse(e, envKey, parseUint8, options...)
}

// CSVUint16 returns a slice of uint16 from a comma separated
// environment variable value and returns an error if any value
// is not a valid uint16 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (e *Env) CSVUint16(envKey string, options ...Option) (values []uint16, err error) {
	return csvParse(e, envKey, parseUint16, options...)
}

// CSVUint32 returns a slice of uint32 from a comma separated
// environment variable value and returns an error if any value
// is not a valid uint32 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (e *Env) CSVUint32(envKey string, options ...Option) (values []uint32, err error) {
	return csvParse(e, envKey, parseUint32, options...)
}

// CSVUint64 returns a slice of uint64 from a comma separated
// environment variable value and returns an error if any value
// is not a valid uint64 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (e *Env) CSVUint64(envKey string, options ...Option) (values []uint64, err error) {
	return csvParse(e, envKey, parseUint64, options...)
}

func csvParse[T any](env *Env, envKey string,
	parse func(value string) (output T, err error),
	options ...Option) (
	values []T, err error) {
	stringValues := env.CSV(envKey, options...)
	if stringValues == nil {
		return nil, nil
	}

	values = make([]T, len(stringValues))
	for i, stringValue := range stringValues {
		values[i], err = parse(stringValue)
		if err != nil {
			return nil, fmt.Errorf("environment variable %s: %w", envKey, err)
		}
	}

	return values, nil
}
