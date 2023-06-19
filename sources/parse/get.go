package parse

import "strings"

// Get returns the value found at the given key in the given
// keyValues map, as a string pointer.
//
// The value is returned as `nil` if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AcceptEmpty option, if the
//     key is set in the mapping and its corresponding value is empty.
//
// Otherwise, the value may be modified depending on the parse
// default settings and the parse options given.
//
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func Get(keyValues map[string]string, key string, options ...Option) (value *string) {
	settings := settingsFromOptions(options)

	keysToTry := make([]string, 0, 1+len(settings.deprecatedKeys))
	keysToTry = append(keysToTry, settings.deprecatedKeys...)
	// Note we try the current key last because it might be set
	// in the released program (such as a Docker image), so we want
	// to take the older configuration from the user first.
	keysToTry = append(keysToTry, key)

	var firstKeySet string
	for _, keyToTry := range keysToTry {
		stringValue, isSet := keyValues[keyToTry]
		if !isSet {
			continue
		}
		firstKeySet = key
		value = new(string)
		*value = stringValue
		break
	}

	if firstKeySet == "" { // All keys are unset
		return nil
	}

	if firstKeySet != key {
		settings.handleDeprecatedKey(firstKeySet, key)
	}

	if !*settings.acceptEmpty && *value == "" {
		// value is set to the empty string, but the empty
		// string is not accepted so return nil.
		return nil
	}

	*value = postProcessValue(*value, settings)
	return value
}

func postProcessValue(value string, settings settings) string {
	if *settings.forceLowercase {
		value = strings.ToLower(value)
	}

	cutSet := map[string]struct{}{}
	if *settings.trimSpace {
		// Only latin charset
		spaceCharacters := []rune{'\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0}
		for _, r := range spaceCharacters {
			cutSet[string(r)] = struct{}{}
		}
	}

	if *settings.trimLineEndings {
		cutSet["\r"] = struct{}{}
		cutSet["\n"] = struct{}{}
	}

	if *settings.trimQuotes {
		cutSet[`"`] = struct{}{}
		cutSet[`'`] = struct{}{}
	}

	cutSetString := ""
	for s := range cutSet {
		cutSetString += s
	}

	return strings.Trim(value, cutSetString)
}
