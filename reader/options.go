package reader

import (
	"github.com/qdm12/gosettings"
	"github.com/qdm12/gosettings/reader/parse"
)

// Option is an option to modify the behavior of
// the `Get(key string, options ...Option)` method
// which is used by all methods.
type Option func(s *settings)

// ForceLowercase forces the string values read from the
// reader to be lowercased or not, depending on the
// `lowercase` argument given.
func ForceLowercase(lowercase bool) Option {
	return func(s *settings) {
		s.forceLowercase = &lowercase
	}
}

// AcceptEmpty, if set to true, makes the code distinguish
// between unset keys and empty values.
// By default, the code does not distinguish between the two cases.
func AcceptEmpty(accept bool) Option {
	return func(s *settings) {
		s.acceptEmpty = &accept
	}
}

// RetroKeys specifies a list of keys that are deprecated
// and replaced by the current key.
// The oldest retro-compatible key should be placed first
// in the list of retro keys, so it gets checked first, which
// is especially important if for example default variable values
// are set in the program or operating system matching more recent keys.
// The `handleDeprecatedKey` function is called when a deprecated
// key is used, with the source name, the deprecated key and
// the current key as arguments.
func RetroKeys(retroKeys ...string) Option {
	return func(s *settings) {
		s.retroKeys = retroKeys
	}
}

type settings struct {
	forceLowercase *bool
	acceptEmpty    *bool
	retroKeys      []string
}

func (s settings) copy() settings {
	return settings{
		forceLowercase: gosettings.CopyPointer(s.forceLowercase),
		acceptEmpty:    gosettings.CopyPointer(s.acceptEmpty),
		retroKeys:      gosettings.CopySlice(s.retroKeys),
	}
}

func (r *Reader) makeParseOptions(options []Option) (parseOptions []parse.Option) {
	settings := r.defaultReadSettings.copy()
	for _, option := range options {
		option(&settings)
	}

	const maxOptions = 3
	parseOptions = make([]parse.Option, 0, maxOptions)
	if settings.forceLowercase != nil {
		parseOption := parse.ForceLowercase(*settings.forceLowercase)
		parseOptions = append(parseOptions, parseOption)
	}
	if settings.acceptEmpty != nil {
		parseOption := parse.AcceptEmpty(*settings.acceptEmpty)
		parseOptions = append(parseOptions, parseOption)
	}
	if len(settings.retroKeys) > 0 {
		parseOption := parse.RetroKeys(r.handleDeprecatedKey, settings.retroKeys...)
		parseOptions = append(parseOptions, parseOption)
	}

	return parseOptions
}
