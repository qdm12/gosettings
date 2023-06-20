package env

import "strings"

// Env is an environment variables source parser
// based on functions from the sources/parser package.
type Env struct {
	keyToValue          map[string]string
	handleDeprecatedKey func(deprecatedKey string, currentKey string)
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
	keyToValue := make(map[string]string, len(environ))
	for _, keyValue := range environ {
		const maxParts = 2
		parts := strings.SplitN(keyValue, "=", maxParts)
		if len(parts) != maxParts {
			panic("invalid environment variable: " + keyValue)
		}
		keyToValue[parts[0]] = parts[1]
	}

	if handleDeprecatedKey == nil {
		handleDeprecatedKey = func(oldKey string, currentKey string) {}
	}

	return &Env{
		keyToValue:          keyToValue,
		handleDeprecatedKey: handleDeprecatedKey,
	}
}
