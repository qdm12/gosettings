package merger

// MergeableSettings is an interface for settings
// that can be merged with another settings struct.
// It is defined largely so the merger source can
// be generic over the settings type.
type MergeableSettings[T any] interface {
	MergeWith(other T) (merged T)
}

// Source is a source of settings for a particular
// concrete settings type implementing the MergeableSettings
// interface, and in particular the method:
// `MergeWith(other T) (merged T)`.
type Source[T MergeableSettings[T]] interface {
	Read() (settings T, err error)
	String() string
}

type emptySettings struct{}

func (s emptySettings) MergeWith(other emptySettings) (merged emptySettings) {
	return merged
}

var _ Source[emptySettings] = (*Merger[emptySettings])(nil)
