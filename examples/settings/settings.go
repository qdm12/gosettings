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
	s.Enabled = defaults.Pointer(s.Enabled, true)
}

func (s *Settings) Validate() error {
	return nil
}

func (s *Settings) Copy() Settings {
	return Settings{
		Enabled: copier.Pointer(s.Enabled),
	}
}

func (s *Settings) MergeWith(other Settings) {
	s.Enabled = merge.WithPointer(s.Enabled, other.Enabled)
}

func (s *Settings) OverrideWith(other Settings) {
	s.Enabled = override.WithPointer(s.Enabled, other.Enabled)
}

func (s *Settings) String() string {
	return fmt.Sprintf(`Settings:
- Enabled: %t`, *s.Enabled)
}
