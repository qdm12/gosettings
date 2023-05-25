package settings

import (
	"fmt"

	"github.com/qdm12/gosettings"
)

type Settings struct {
	Enabled *bool
}

func (s *Settings) SetDefaults() {
	s.Enabled = gosettings.DefaultPointer(s.Enabled, true)
}

func (s *Settings) Validate() error {
	return nil
}

func (s *Settings) Copy() Settings {
	return Settings{
		Enabled: gosettings.CopyPointer(s.Enabled),
	}
}

func (s *Settings) MergeWith(other Settings) {
	s.Enabled = gosettings.MergeWithPointer(s.Enabled, other.Enabled)
}

func (s *Settings) OverrideWith(other Settings) {
	s.Enabled = gosettings.OverrideWithPointer(s.Enabled, other.Enabled)
}

func (s *Settings) String() string {
	return fmt.Sprintf(`Settings:
- Enabled: %t`, *s.Enabled)
}
