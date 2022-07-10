package override

import "time"

func Duration(existing, other *time.Duration) (result *time.Duration) {
	if other == nil {
		return existing
	}
	result = new(time.Duration)
	*result = *other
	return result
}
