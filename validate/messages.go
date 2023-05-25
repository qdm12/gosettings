package validate

import "fmt"

func orStrings[T comparable](elements []T) (result string) {
	return joinStrings(elements, "or")
}

func joinStrings[T comparable](elements []T, lastJoin string) (result string) {
	if len(elements) == 0 {
		return ""
	}

	result = fmt.Sprint(elements[0])
	for i := 1; i < len(elements); i++ {
		if i < len(elements)-1 {
			result += fmt.Sprint(elements[i]) + ", "
		} else {
			result += " " + lastJoin + " " + fmt.Sprint(elements[i])
		}
	}

	return result
}
