package envhelpers

import "time"

func DurationPtr(envKey string, options ...Option) (
	durationPtr *time.Duration, err error) {
	s := Get(envKey, options...)
	if s == "" {
		return nil, nil //nolint:nilnil
	}

	durationPtr = new(time.Duration)
	*durationPtr, err = time.ParseDuration(s)
	if err != nil {
		return nil, err
	}

	return durationPtr, nil
}
