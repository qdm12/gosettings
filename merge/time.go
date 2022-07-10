package merge

import "time"

func Duration(existing, other *time.Duration) (result *time.Duration) {
	if existing != nil {
		return existing
	}
	return other
}
