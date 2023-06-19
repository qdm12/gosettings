package parse

// FirstKeySet returns the first key set from a list of keys,
// or an empty string if none of the keys are set.
func FirstKeySet(keyValues map[string]string, keys ...string) (firstKeySet string) {
	for _, key := range keys {
		_, set := keyValues[key]
		if set {
			return key
		}
	}
	return ""
}
