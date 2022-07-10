package validate

import (
	"errors"
	"fmt"
)

var ErrValueNotOneOf = errors.New("value is not one of the possible values")

func IsOneOf(value string, possibilities ...string) (err error) {
	for _, possibility := range possibilities {
		if value == possibility {
			return nil
		}
	}

	return fmt.Errorf("%w: %q must be one of %s",
		ErrValueNotOneOf, value, orStrings(possibilities))
}
