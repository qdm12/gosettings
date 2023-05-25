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
	s.trimLineEndings = defaults.Pointer(s.trimLineEndings, true)
	s.trimSpace = defaults.Pointer(s.trimSpace, true)
	s.forceLowercase = defaults.Pointer(s.forceLowercase, true)
}
