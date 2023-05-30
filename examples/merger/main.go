package main

import (
	"fmt"

	"github.com/qdm12/gosettings/examples/settings"
	"github.com/qdm12/gosettings/sources/merger"
)

func main() {
	sourceA := new(SourceA)
	sourceB := new(SourceB)
	merger := merger.New[settings.Settings](sourceA, sourceB)
	settings, err := merger.Read()
	if err != nil {
		panic(err)
	}
	settings.SetDefaults() // set remaining unset fields
	fmt.Println(settings)
}

// SourceA is a settings source.
type SourceA struct{}

func (s *SourceA) String() string {
	return "A"
}

func (s *SourceA) Read() (settings settings.Settings, err error) {
	settings.Enabled = nil // no setting found
	return settings, err
}

// SourceB is a settings source.
type SourceB struct{}

func (s *SourceB) String() string {
	return "B"
}

func (f *SourceB) Read() (settings settings.Settings, err error) {
	enabled := false
	settings.Enabled = &enabled // found enabled as false
	return settings, nil
}
