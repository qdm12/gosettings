package gosettings

import "testing"

func Test_MergeWithComparable(t *testing.T) {
	t.Parallel()

	t.Run("no_type_assertion_needed", func(t *testing.T) {
		t.Parallel()

		var existing testInterface
		existing = MergeWithComparable[testInterface](
			existing, &testInterfaceImplA{})
		existing.F() // make use of variable
	})

	testCases := map[string]struct {
		existing any
		other    any
		result   any
	}{
		"all_untyped_nils": {},
		"all_typed_nils": {
			existing: (testInterface)(nil),
			other:    (testInterface)(nil),
			result:   (testInterface)(nil),
		},
		"no_merge": {
			existing: &testInterfaceImplA{},
			other:    &testInterfaceImplB{},
			result:   &testInterfaceImplA{},
		},
		"merge": {
			existing: (testInterface)(nil),
			other:    &testInterfaceImplB{},
			result:   &testInterfaceImplB{},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := MergeWithComparable(testCase.existing, testCase.other)
			if result != testCase.result {
				t.Errorf("expected %v, got %v", testCase.result, result)
			}
		})
	}
}
