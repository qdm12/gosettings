package env

import "strings"

// Env is an environment variables reader
// with helping methods to parse values.
type Env struct {
	environ map[string]string
}

// New creates a new environment variables reader
// and initializes it with the given slice of environment
// variable strings, where each string is in the form
// "key=value".
func New(environ []string) *Env {
	environMap := make(map[string]string, len(environ))
	for _, keyValue := range environ {
		const maxParts = 2
		parts := strings.SplitN(keyValue, "=", maxParts)
		if len(parts) != maxParts {
			panic("invalid environment variable: " + keyValue)
		}
		environMap[parts[0]] = parts[1]
	}

	return &Env{
		environ: environMap,
	}
}
