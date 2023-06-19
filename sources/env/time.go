package env

import (
	"time"

	"github.com/qdm12/gosettings/sources/parse"
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
	parseOptions := e.makeParseOptions(options)
	return parse.DurationPtr(e.environ, envKey, parseOptions...)
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
	parseOptions := e.makeParseOptions(options)
	return parse.Duration(e.environ, envKey, parseOptions...)
}
