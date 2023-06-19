package env

import (
	"fmt"
	"time"
)

// DurationPtr returns a pointer to a `time.Duration`
// from an environment variable value.
// If the value is not a valid time.Duration string, an error
// is returned with the environment variable name in its message.
// The value is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
func (e *Env) DurationPtr(envKey string, options ...Option) (
	durationPtr *time.Duration, err error) {
	s := e.Get(envKey, options...)
	if s == nil {
		return nil, nil //nolint:nilnil
	}

	durationPtr = new(time.Duration)
	*durationPtr, err = time.ParseDuration(*s)
	if err != nil {
		return nil, fmt.Errorf("environment variable %s: %w", envKey, err)
	}

	return durationPtr, nil
}

// Duration returns a `time.Duration` from an environment
// variable value.
// If the value is not a valid time.Duration string, an error
// is returned with the environment variable name in its message.
// The value is returned as `0` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
func (e *Env) Duration(envKey string, options ...Option) (
	duration time.Duration, err error) {
	s := e.Get(envKey, options...)
	if s == nil {
		return 0, nil
	}

	duration, err = time.ParseDuration(*s)
	if err != nil {
		return 0, fmt.Errorf("environment variable %s: %w", envKey, err)
	}

	return duration, nil
}
