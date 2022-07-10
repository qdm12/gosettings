package validate

import (
	"errors"
	"fmt"
	"regexp"
)

var ErrValueMismatchRegex = errors.New("value does not match regular expression")

func MatchRegex(value string, regex *regexp.Regexp) (err error) {
	if regex.MatchString(value) {
		return nil
	}
	return fmt.Errorf("%w: %q does not match regular expression %s",
		ErrValueMismatchRegex, value, regex)
}

func AllMatchRegex(values []string, regex *regexp.Regexp) (err error) {
	for _, value := range values {
		if err = MatchRegex(value, regex); err != nil {
			return err
		}
	}

	return nil
}
