package env

import "github.com/qdm12/gosettings/defaults"

type settings struct {
	trimLineEndings *bool
	trimSpace       *bool
	forceLowercase  *bool
}

func settingsFromOptions(options []Option) (s settings) {
	for _, option := range options {
		option(&s)
	}
	s.setDefaults()
	return s
}

func (s *settings) setDefaults() {
	s.trimLineEndings = defaults.Bool(s.trimLineEndings, true)
	s.trimSpace = defaults.Bool(s.trimSpace, true)
	s.forceLowercase = defaults.Bool(s.forceLowercase, true)
}
