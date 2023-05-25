//go:build go1.18
// +build go1.18

package validate

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNoChoice      = errors.New("one or more values is set but there is no possible value available")
	ErrValueNotOneOf = errors.New("value is not one of the possible choices")
)

func IsOneOf[T comparable](value T, possibilities ...T) (err error) {
	for _, possibility := range possibilities {
		if value == possibility {
			return nil
		}
	}

	return fmt.Errorf("%w: %v must be one of %s",
		ErrValueNotOneOf, value, orStrings(possibilities))
}

func AreAllOneOf[T comparable](values, choices []T) (err error) {
	if len(values) > 0 && len(choices) == 0 {
		return fmt.Errorf("%w", ErrNoChoice)
	}

	set := make(map[T]struct{}, len(choices))
	for _, choice := range choices {
		set[choice] = struct{}{}
	}

	for _, value := range values {
		_, ok := set[value]
		if !ok {
			choiceStrings := make([]string, len(choices))
			for i := range choices {
				choiceStrings[i] = fmt.Sprintf("%v", choices[i])
			}
			return fmt.Errorf("%w: value %v, choices available are %s",
				ErrValueNotOneOf, value, strings.Join(choiceStrings, ", "))
		}
	}

	return nil
}
