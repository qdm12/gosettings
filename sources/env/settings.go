package env

import (
	"github.com/qdm12/gosettings"
)

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
	s.trimLineEndings = gosettings.DefaultPointer(s.trimLineEndings, true)
	s.trimSpace = gosettings.DefaultPointer(s.trimSpace, true)
	s.forceLowercase = gosettings.DefaultPointer(s.forceLowercase, true)
}
