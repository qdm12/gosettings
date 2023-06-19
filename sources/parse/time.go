package parse

import (
	"time"
)

// DurationPtr returns a pointer to a `time.Duration`
// from the value found at the given key in the given
// keyValues map.
// If the value is not a valid time.Duration string, an error
// is returned with the key in its message.
// The value is returned as `nil` if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func DurationPtr(keyValues map[string]string, key string,
	options ...Option) (durationPtr *time.Duration, err error) {
	return GetParsePtr(keyValues, key, time.ParseDuration, options...)
}

// Duration returns a `time.Duration` from the value found
// at the given key in the given keyValues map.
// If the value is not a valid time.Duration string, an error
// is returned with the key in its message.
// The value is returned as `0` if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func Duration(keyValues map[string]string, key string,
	options ...Option) (duration time.Duration, err error) {
	return GetParse(keyValues, key, time.ParseDuration, options...)
}
