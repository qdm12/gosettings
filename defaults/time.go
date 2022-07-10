package defaults

import "time"

func Duration(existing *time.Duration,
	defaultValue time.Duration) (result *time.Duration) {
	if existing != nil {
		return existing
	}
	result = new(time.Duration)
	*result = defaultValue
	return result
}
