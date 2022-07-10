package validate

import (
	"errors"
	"fmt"
)

var ErrValueOutOfBounds = errors.New("value is out of bounds")

func IntBetween(n, min, max int) (err error) {
	if n < min || n > max {
		return fmt.Errorf("%w: %d must be between %d and %d included",
			ErrValueOutOfBounds, n, min, max)
	}
	return nil
}
