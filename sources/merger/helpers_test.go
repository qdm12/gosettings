package merger

import (
	"testing"

	"github.com/qdm12/gosettings"
)

type settings struct {
	x int
}

func (s settings) MergeWith(other settings) (merged settings) {
	merged.x = gosettings.MergeWithComparable(s.x, other.x)
	return merged
}

type sourceA struct{}

func (s *sourceA) String() string {
	return "A"
}

func (s *sourceA) Read() (settings, error) {
	return settings{x: 0}, nil
}

type sourceB struct{}

func (s *sourceB) String() string {
	return "B"
}

func (s *sourceB) Read() (settings, error) {
	return settings{x: 1}, nil
}

type sourceError struct{ err error }

func (s *sourceError) String() string {
	return "SourceError"
}

func (s *sourceError) Read() (settings, error) {
	return settings{}, s.err
}

func Test_emptySettings_MergeWith(t *testing.T) { // just for 100% coverage
	t.Parallel()

	empty := emptySettings{}
	_ = empty.MergeWith(empty)
}
