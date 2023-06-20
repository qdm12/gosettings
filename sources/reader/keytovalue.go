package reader

import "strings"

func keyToValueFromEnviron(environ []string) (keyToValue map[string]string) {
	keyToValue = make(map[string]string, len(environ))
	for _, keyValue := range environ {
		const maxParts = 2
		parts := strings.SplitN(keyValue, "=", maxParts)
		if len(parts) != maxParts {
			continue
		}
		keyToValue[parts[0]] = parts[1]
	}
	return keyToValue
}
