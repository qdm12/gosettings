package validate

import "testing"

func Test_orStrings(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		elements []string
		result   string
	}{
		"empty": {},
		"one": {
			elements: []string{"a"},
			result:   "a",
		},
		"two": {
			elements: []string{"a", "b"},
			result:   "a or b",
		},
		"three": {
			elements: []string{"a", "b", "c"},
			result:   "a, b or c",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := orStrings(testCase.elements)
			if result != testCase.result {
				t.Errorf("expected %q, got %q", testCase.result, result)
			}
		})
	}
}
