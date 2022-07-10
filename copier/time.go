package copier

import "time"

func Duration(original *time.Duration) (copied *time.Duration) {
	if original == nil {
		return nil
	}
	copied = new(time.Duration)
	*copied = *original
	return copied
}
