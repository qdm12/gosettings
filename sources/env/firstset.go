package env

import "github.com/qdm12/gosettings/sources/parse"

// FirstKeySet returns the first environment variable key set
// with a non empty value from a list of keys, or an empty string
// if none of the keys are set.
func (e *Env) FirstKeySet(keys ...string) (firstKeySet string) {
	return parse.FirstKeySet(e.keyToValue, keys...)
}
