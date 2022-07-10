package settings

import (
	"github.com/qdm12/gosettings/copier"
	"github.com/qdm12/gosettings/defaults"
	"github.com/qdm12/gosettings/merge"
	"github.com/qdm12/gosettings/override"
	"github.com/qdm12/gotree"
)

type Settings struct {
	Enabled *bool
}

func (s *Settings) SetDefaults() {
	s.Enabled = defaults.Bool(s.Enabled, true)
}

func (s *Settings) Validate() error {
	return nil
}

func (s *Settings) Copy() Settings {
	return Settings{
		Enabled: copier.Bool(s.Enabled),
	}
}

func (s *Settings) MergeWith(other Settings) {
	s.Enabled = merge.Bool(s.Enabled, other.Enabled)
}

func (s *Settings) OverrideWith(other Settings) {
	s.Enabled = override.Bool(s.Enabled, other.Enabled)
}

func (s *Settings) ToLinesNode() *gotree.Node {
	node := gotree.New("Settings:")
	node.Appendf("Enabled: %t", *s.Enabled)
	return node
}

func (s *Settings) String() string {
	return s.ToLinesNode().String()
}
