package env

import "github.com/qdm12/gosettings/sources/parse"

// CSVInt returns a slice of int from a comma separated
// environment variable value and returns an error if any value
// is not a valid int string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (r *Reader) CSVInt(envKey string, options ...Option) (values []int, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVInt(r.keyToValue, envKey, parseOptions...)
}

// CSVInt8 returns a slice of int8 from a comma separated
// environment variable value and returns an error if any value
// is not a valid int8 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (r *Reader) CSVInt8(envKey string, options ...Option) (values []int8, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVInt8(r.keyToValue, envKey, parseOptions...)
}

// CSVInt16 returns a slice of int16 from a comma separated
// environment variable value and returns an error if any value
// is not a valid int16 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (r *Reader) CSVInt16(envKey string, options ...Option) (values []int16, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVInt16(r.keyToValue, envKey, parseOptions...)
}

// CSVInt32 returns a slice of int32 from a comma separated
// environment variable value and returns an error if any value
// is not a valid int32 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (r *Reader) CSVInt32(envKey string, options ...Option) (values []int32, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVInt32(r.keyToValue, envKey, parseOptions...)
}

// CSVInt64 returns a slice of int64 from a comma separated
// environment variable value and returns an error if any value
// is not a valid int64 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (r *Reader) CSVInt64(envKey string, options ...Option) (values []int64, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVInt64(r.keyToValue, envKey, parseOptions...)
}

// CSVUint returns a slice of uint from a comma separated
// environment variable value and returns an error if any value
// is not a valid uint string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (r *Reader) CSVUint(envKey string, options ...Option) (values []uint, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVUint(r.keyToValue, envKey, parseOptions...)
}

// CSVUint8 returns a slice of uint8 from a comma separated
// environment variable value and returns an error if any value
// is not a valid uint8 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (r *Reader) CSVUint8(envKey string, options ...Option) (values []uint8, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVUint8(r.keyToValue, envKey, parseOptions...)
}

// CSVUint16 returns a slice of uint16 from a comma separated
// environment variable value and returns an error if any value
// is not a valid uint16 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (r *Reader) CSVUint16(envKey string, options ...Option) (values []uint16, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVUint16(r.keyToValue, envKey, parseOptions...)
}

// CSVUint32 returns a slice of uint32 from a comma separated
// environment variable value and returns an error if any value
// is not a valid uint32 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (r *Reader) CSVUint32(envKey string, options ...Option) (values []uint32, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVUint32(r.keyToValue, envKey, parseOptions...)
}

// CSVUint64 returns a slice of uint64 from a comma separated
// environment variable value and returns an error if any value
// is not a valid uint64 string.
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
func (r *Reader) CSVUint64(envKey string, options ...Option) (values []uint64, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVUint64(r.keyToValue, envKey, parseOptions...)
}
