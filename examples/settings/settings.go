package settings

import (
	"fmt"

	"github.com/qdm12/gosettings"
)

// Settings holds all the settings.
type Settings struct {
	Enabled *bool
}

// SetDefaults sets the default values for the settings
// if they are not already set.
func (s *Settings) SetDefaults() {
	s.Enabled = gosettings.DefaultPointer(s.Enabled, true)
}

// Validate validates the settings and returns an error
// if one setting is not valid.
func (s *Settings) Validate() error {
	return nil
}

// Copy returns a copy of the settings.
func (s *Settings) Copy() Settings {
	return Settings{
		Enabled: gosettings.CopyPointer(s.Enabled),
	}
}

// MergeWith merges the receiver settings with the fields from
// another settings struct.
func (s *Settings) MergeWith(other Settings) {
	s.Enabled = gosettings.MergeWithPointer(s.Enabled, other.Enabled)
}

// OverrideWith overrides the settings with another settings struct.
func (s *Settings) OverrideWith(other Settings) {
	s.Enabled = gosettings.OverrideWithPointer(s.Enabled, other.Enabled)
}

// String returns a string representation of the settings.
func (s *Settings) String() string {
	return fmt.Sprintf(`Settings:
- Enabled: %t`, *s.Enabled)
}
