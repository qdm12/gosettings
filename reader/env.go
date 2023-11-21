package reader

import (
	"strings"
)

// Env implements an environment variables source.
// Note all keys are transformed using its KeyTransform
// method.
type Env struct {
	keyToValue map[string]string
}

// NewEnv creates a new environment variable source
// given the environment as a slice of key-value pairs,
// which can generally be obtained with os.Environ().
// Environ values with no '=' sign are ignored.
// All environment variable keys read are eventually
// transformed using the KeyTransform method.
func NewEnv(environ []string) (env *Env) {
	env = &Env{
		keyToValue: make(map[string]string, len(environ)),
	}

	for _, keyValue := range environ {
		const maxParts = 2
		parts := strings.SplitN(keyValue, "=", maxParts)
		if len(parts) != maxParts {
			continue
		}

		key := env.KeyTransform(parts[0])
		value := parts[1]
		env.keyToValue[key] = value
	}

	return env
}

func (e *Env) String() string {
	return "environment variable"
}

// Get returns the value of the environment variable
// found at the given key, and a boolean `isSet` to
// indicate if it is set or not.
func (e *Env) Get(key string) (value string, isSet bool) {
	value, isSet = e.keyToValue[key]
	return value, isSet
}

// KeyTransform transforms a generic key to an environment
// variable key. It notably:
// - Changes all characters to be uppercase
// - Replaces all dashes with underscores.
func (e *Env) KeyTransform(key string) (newKey string) {
	newKey = strings.ToUpper(key)
	newKey = strings.ReplaceAll(newKey, "-", "_")
	return newKey
}
