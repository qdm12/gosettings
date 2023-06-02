package env

import "strings"

// Env is an environment variables reader
// with helping methods to parse values.
type Env struct {
	environ             map[string]string
	handleDeprecatedKey func(oldKey string, currentKey string)
}

// New creates a new environment variables reader
// and initializes it with the given slice of environment
// variable strings, where each string is in the form
// "key=value". The functional argument `handleDeprecatedKey`
// is called when encountering a deprecated environment variable
// key, and defaults to a no-op function if left to `nil`.
func New(environ []string,
	handleDeprecatedKey func(deprecatedKey string, currentKey string),
) *Env {
	environMap := make(map[string]string, len(environ))
	for _, keyValue := range environ {
		const maxParts = 2
		parts := strings.SplitN(keyValue, "=", maxParts)
		if len(parts) != maxParts {
			panic("invalid environment variable: " + keyValue)
		}
		environMap[parts[0]] = parts[1]
	}

	if handleDeprecatedKey == nil {
		handleDeprecatedKey = func(oldKey string, currentKey string) {}
	}

	return &Env{
		environ:             environMap,
		handleDeprecatedKey: handleDeprecatedKey,
	}
}
