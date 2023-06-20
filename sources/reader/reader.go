package reader

// Reader is an environment variables source parser
// based on functions from the sources/parser package.
type Reader struct {
	keyToValue          map[string]string
	handleDeprecatedKey func(deprecatedKey string, currentKey string)
}

// New creates a new environment variables reader
// and initializes it with the given slice of environment
// variable strings, where each string is in the form
// "key=value". The functional argument `handleDeprecatedKey`
// is called when encountering a deprecated environment variable
// key, and defaults to a no-op function if left to `nil`.
func New(settings Settings) *Reader {
	settings.setDefaults()

	keyToValue := keyToValueFromEnviron(settings.Environ)
	return &Reader{
		keyToValue:          keyToValue,
		handleDeprecatedKey: settings.HandleDeprecatedKey,
	}
}

// Settings is the settings to create a new reader.
type Settings struct {
	// Environ is a slice of environment variable strings,
	// where each string is in the form "key=value".
	Environ []string
	// HandleDeprecatedKey is called when encountering a deprecated
	// key, and defaults to a no-op function.
	HandleDeprecatedKey func(deprecatedKey string, currentKey string)
}

func (s *Settings) setDefaults() {
	if s.HandleDeprecatedKey == nil { // Note: cannot use DefaultInterface
		s.HandleDeprecatedKey = func(deprecatedKey string, currentKey string) {}
	}
}
