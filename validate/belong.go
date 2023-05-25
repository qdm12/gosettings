//go:build go1.18
// +build go1.18

package validate

import (
	"errors"
	"fmt"
)

var ErrValueNotOneOf = errors.New("value is not one of the possible values")

func IsOneOf[T comparable](value T, possibilities ...T) (err error) {
	for _, possibility := range possibilities {
		if value == possibility {
			return nil
		}
	}

	return fmt.Errorf("%w: %v must be one of %s",
		ErrValueNotOneOf, value, orStrings(possibilities))
}
