package env

// Option is an option to modify the behavior of
// `env.Get(key string, options ...Option)`.
type Option func(s *settings)

// ForceLowercase forces the string values read from the
// environment to be lowercased or not, depending on the
// `lowercase` argument given.
func ForceLowercase(lowercase bool) Option {
	return func(s *settings) {
		s.forceLowercase = &lowercase
	}
}

// AcceptEmpty allows to have set but empty environment
// variables.
func AcceptEmpty(accept bool) Option {
	return func(s *settings) {
		s.acceptEmpty = &accept
	}
}

// RetroKeys specifies a list of environment variable keys
// that are deprecated and replaced by the current key.
// The `handleDeprecatedKey` function is called when a deprecated
// key is used, with the deprecated key and the current key as
// arguments.
func RetroKeys(retroKeys ...string) Option {
	return func(s *settings) {
		s.retroKeys = retroKeys
	}
}
