package settings

import (
	"errors"
	"fmt"
	"strings"

	"github.com/qdm12/gosettings"
)

// Settings holds all the settings.
type Settings struct {
	Enabled *bool
	Names   []string
}

// SetDefaults sets the default values for the settings
// if they are not already set.
func (s *Settings) SetDefaults() {
	s.Enabled = gosettings.DefaultPointer(s.Enabled, true)
	s.Names = gosettings.DefaultSlice(s.Names, []string{"Alice", "Bob"})
}

var (
	ErrNameContainsSpace = errors.New("name contains a space")
)

// Validate validates the settings and returns an error
// if one setting is not valid.
func (s *Settings) Validate() error {
	// Names cannot contain spaces
	for _, name := range s.Names {
		if strings.ContainsRune(name, ' ') {
			return fmt.Errorf("%w: %s", ErrNameContainsSpace, name)
		}
	}
	return nil
}

// Copy returns a copy of the settings.
func (s *Settings) Copy() Settings {
	return Settings{
		Enabled: gosettings.CopyPointer(s.Enabled),
		Names:   gosettings.CopySlice(s.Names),
	}
}

// MergeWith merges the settings with another settings struct
// and returns the result as a new settings struct.
func (s Settings) MergeWith(other Settings) (merged Settings) {
	merged.Enabled = gosettings.MergeWithPointer(s.Enabled, other.Enabled)
	merged.Names = gosettings.MergeWithSlice(s.Names, other.Names)
	return merged
}

// OverrideWith overrides the settings with another settings struct.
func (s *Settings) OverrideWith(other Settings) {
	s.Enabled = gosettings.OverrideWithPointer(s.Enabled, other.Enabled)
	s.Names = gosettings.OverrideWithSlice(s.Names, other.Names)
}

// String returns a string representation of the settings.
func (s Settings) String() string {
	return fmt.Sprintf(`Settings:
	- Enabled: %t
	- Names: %s`, *s.Enabled, strings.Join(s.Names, ", "))
}
