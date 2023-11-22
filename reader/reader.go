package reader

import (
	"os"

	"github.com/qdm12/gosettings"
	"github.com/qdm12/gosettings/reader/parse"
	"github.com/qdm12/gosettings/reader/sources/env"
)

// Reader is a settings sources reader and parser.
type Reader struct {
	sources             []parse.Source
	handleDeprecatedKey func(source, deprecatedKey, currentKey string)
	defaultReadSettings settings
}

// New creates a new reader using the settings given.
func New(readerSettings Settings) *Reader {
	readerSettings.setDefaults()

	var defaultReadSettings settings
	for _, defaultOption := range readerSettings.DefaultOptions {
		defaultOption(&defaultReadSettings)
	}

	parseSources := make([]parse.Source, len(readerSettings.Sources))
	for i, source := range readerSettings.Sources {
		parseSources[i] = source
	}

	return &Reader{
		sources:             parseSources,
		handleDeprecatedKey: readerSettings.HandleDeprecatedKey,
		defaultReadSettings: defaultReadSettings,
	}
}

// Settings is the settings to create a new reader.
type Settings struct {
	// Sources is a slice of sources where a source at
	// a lower index has a higher priority.
	// It defaults to:
	// []reader.Source{env.New(os.Environ())}
	Sources []Source
	// HandleDeprecatedKey is called when encountering a deprecated
	// key, and defaults to a no-op function.
	HandleDeprecatedKey func(source, deprecatedKey, currentKey string)
	// DefaultOptions are the default options to use for every method call.
	// They default to ForceLowercase(false), AcceptEmpty(false).
	DefaultOptions []Option
}

func (s *Settings) setDefaults() {
	s.Sources = gosettings.DefaultSlice(s.Sources,
		[]Source{env.New(os.Environ())})

	if s.HandleDeprecatedKey == nil { // Note: cannot use DefaultInterface
		s.HandleDeprecatedKey = func(source, deprecatedKey, currentKey string) {}
	}
	s.DefaultOptions = gosettings.DefaultSlice(s.DefaultOptions,
		[]Option{ForceLowercase(false), AcceptEmpty(false)})
}
