package env

import "testing"

func ptrTo[T any](x T) *T { return &x }

func Test_postProcessValue(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		value    string
		settings settings
		result   string
	}{
		"no_post_processing": {
			value: " Value\n ",
			settings: settings{
				trimLineEndings: ptrTo(false),
				trimSpace:       ptrTo(false),
				trimQuotes:      ptrTo(false),
				forceLowercase:  ptrTo(false),
			},
			result: " Value\n ",
		},
		"remove_line_endings": {
			value: " Value\n\r\n\n",
			settings: settings{
				trimLineEndings: ptrTo(true),
				trimSpace:       ptrTo(false),
				trimQuotes:      ptrTo(false),
				forceLowercase:  ptrTo(false),
			},
			result: " Value",
		},
		"remove_spaces": {
			value: " Value \t\v",
			settings: settings{
				trimLineEndings: ptrTo(false),
				trimSpace:       ptrTo(true),
				trimQuotes:      ptrTo(false),
				forceLowercase:  ptrTo(false),
			},
			result: "Value",
		},
		"trim_single_quotes": {
			value: "'Value'",
			settings: settings{
				trimLineEndings: ptrTo(false),
				trimSpace:       ptrTo(false),
				trimQuotes:      ptrTo(true),
				forceLowercase:  ptrTo(false),
			},
			result: "Value",
		},
		"trim_double_quotes": {
			value: "\"Value\"",
			settings: settings{
				trimLineEndings: ptrTo(false),
				trimSpace:       ptrTo(false),
				trimQuotes:      ptrTo(true),
				forceLowercase:  ptrTo(false),
			},
			result: "Value",
		},
		"force_lowercase": {
			value: "Value",
			settings: settings{
				trimLineEndings: ptrTo(false),
				trimSpace:       ptrTo(false),
				trimQuotes:      ptrTo(false),
				forceLowercase:  ptrTo(true),
			},
			result: "value",
		},
		"combined": {
			value: "\t\"'Value\"'\t\n\t\r\n\t",
			settings: settings{
				trimLineEndings: ptrTo(true),
				trimSpace:       ptrTo(true),
				trimQuotes:      ptrTo(true),
				forceLowercase:  ptrTo(true),
			},
			result: "value",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := postProcessValue(testCase.value, testCase.settings)
			if result != testCase.result {
				t.Errorf("expected %s, got %s", testCase.result, result)
			}
		})
	}
}
