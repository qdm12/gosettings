package settings

import (
	"fmt"

	"github.com/qdm12/gosettings/copier"
	"github.com/qdm12/gosettings/defaults"
	"github.com/qdm12/gosettings/merge"
	"github.com/qdm12/gosettings/override"
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

func (s *Settings) String() string {
	return fmt.Sprintf(`Settings:
- Enabled: %t`, *s.Enabled)
}
