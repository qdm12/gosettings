package env

import (
	"fmt"
	"time"
)

// DurationPtr returns a pointer to a `time.Duration`
// from an environment variable value.
// If the environment variable is not set or its value is
// the empty string, `nil` is returned.
// Otherwise, if the value is not a valid duration string,
// an error is returned with the environment variable name
// in its message.
func (e *Env) DurationPtr(envKey string, options ...Option) (
	durationPtr *time.Duration, err error) {
	s := e.Get(envKey, options...)
	if s == nil || *s == "" {
		// note: no point accepting the empty string in this case
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
// If the environment variable is not set or its value is
// the empty string, `0` is returned.
// Otherwise, if the value is not a valid duration string,
// an error is returned with the environment variable name
// in its message.
func (e *Env) Duration(envKey string, options ...Option) (
	duration time.Duration, err error) {
	s := e.Get(envKey, options...)
	if s == nil || *s == "" {
		// note: no point accepting the empty string in this case
		return 0, nil
	}

	duration, err = time.ParseDuration(*s)
	if err != nil {
		return 0, fmt.Errorf("environment variable %s: %w", envKey, err)
	}

	return duration, nil
}
