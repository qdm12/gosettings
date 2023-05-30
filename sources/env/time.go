package env

import (
	"fmt"
	"time"
)

// DurationPtr returns a pointer to a `time.Duration`
// from an environment variable value.
// If the value is the empty string, `nil` is returned.
// Otherwise, if the value is not a valid duration string,
// an error is returned with the environment variable name
// in the error context.
func DurationPtr(envKey string, options ...Option) (
	durationPtr *time.Duration, err error) {
	s := Get(envKey, options...)
	if s == "" {
		return nil, nil //nolint:nilnil
	}

	durationPtr = new(time.Duration)
	*durationPtr, err = time.ParseDuration(s)
	if err != nil {
		return nil, fmt.Errorf("environment variable %s: %w", envKey, err)
	}

	return durationPtr, nil
}
