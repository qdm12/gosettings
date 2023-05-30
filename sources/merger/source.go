package merger

import (
	"fmt"
	"strings"
)

// Merger is a settings source which merges multiple
// settings obtained from multiple settings sources
// together. It is generic over the settings type.
type Merger[T MergeableSettings[T]] struct {
	sources []Source[T]
}

// New creates a new merger source using the settings
// type and multiple sources given. It can be used with for example:
// `merger := New[Settings](envSource, flagsSource)`.
func New[T MergeableSettings[T]](sources ...Source[T]) *Merger[T] {
	return &Merger[T]{
		sources: sources,
	}
}

// String returns 'merger' with the string representation
// of the sources it merges.
func (m *Merger[T]) String() string {
	switch len(m.sources) {
	case 0:
		return "merger"
	case 1:
		return "merger of " + m.sources[0].String()
	}

	subSources := make([]string, len(m.sources))
	for i, source := range m.sources {
		subSources[i] = source.String()
	}
	return "merger of " + strings.Join(subSources, ", ")
}

// Read obtains the settings from each source and merges
// them together. If an error occurs, it is returned
// with the source name prepended to it, and the merging
// stops.
func (m *Merger[T]) Read() (settings T, err error) { //nolint:ireturn
	for _, source := range m.sources {
		newSettings, err := source.Read()
		if err != nil {
			return settings, fmt.Errorf("%s source: %w", source, err)
		}
		settings = settings.MergeWith(newSettings)
	}
	return settings, nil
}
