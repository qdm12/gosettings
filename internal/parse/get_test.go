package parse

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func Test_get(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		makeSources func(ctrl *gomock.Controller) []Source
		key         string
		options     []Option
		value       *string
		sourceKind  string
	}{
		"no_source": {
			key: "KEY",
		},
		"not_found_in_single_source": {
			makeSources: func(ctrl *gomock.Controller) []Source {
				source := NewMockSource(ctrl)
				source.EXPECT().KeyTransform("KEY").Return("key_transformed")
				source.EXPECT().Get("key_transformed").Return("", false)
				return []Source{source}
			},
			key:   "KEY",
			value: nil,
		},
		"found_in_single_source": {
			makeSources: func(ctrl *gomock.Controller) []Source {
				source := NewMockSource(ctrl)
				source.EXPECT().KeyTransform("KEY").Return("key_transformed").Times(2)
				source.EXPECT().Get("key_transformed").Return("value", true)
				source.EXPECT().String().Return("A")
				return []Source{source}
			},
			key:        "KEY",
			value:      ptrTo("value"),
			sourceKind: "A",
		},
		"found_in_2_of_3_sources": {
			makeSources: func(ctrl *gomock.Controller) []Source {
				sourceA := NewMockSource(ctrl)
				sourceA.EXPECT().KeyTransform("KEY").Return("key_transformed")
				sourceA.EXPECT().Get("key_transformed").Return("", false)
				sourceB := NewMockSource(ctrl)
				sourceB.EXPECT().KeyTransform("KEY").Return("key_transformed").Times(2)
				sourceB.EXPECT().Get("key_transformed").Return("value", true)
				sourceB.EXPECT().String().Return("B")
				sourceC := NewMockSource(ctrl)
				return []Source{sourceA, sourceB, sourceC}
			},
			key:        "KEY",
			value:      ptrTo("value"),
			sourceKind: "B",
		},
		"found_current_key_with_retro_keys": {
			makeSources: func(ctrl *gomock.Controller) []Source {
				source := NewMockSource(ctrl)
				source.EXPECT().KeyTransform("OLDEST_DEPRECATED_KEY").Return("oldest_deprecated_key_transformed")
				source.EXPECT().Get("oldest_deprecated_key_transformed").Return("", false)
				source.EXPECT().KeyTransform("NEWEST_DEPRECATED_KEY").Return("newest_deprecated_key_transformed")
				source.EXPECT().Get("newest_deprecated_key_transformed").Return("", false)
				source.EXPECT().KeyTransform("KEY").Return("key_transformed").Times(2)
				source.EXPECT().Get("key_transformed").Return("value", true)
				source.EXPECT().String().Return("A")
				return []Source{source}
			},
			key: "KEY",
			options: []Option{
				RetroKeys(func(source string, deprecateKey string, currentKey string) {},
					"OLDEST_DEPRECATED_KEY", "NEWEST_DEPRECATED_KEY"),
			},
			value:      ptrTo("value"),
			sourceKind: "A",
		},
		"found_after_empty_set_in_first_retrokey": {
			makeSources: func(ctrl *gomock.Controller) []Source {
				source := NewMockSource(ctrl)
				source.EXPECT().KeyTransform("DEPRECATED_KEY").Return("deprecated_key_transformed")
				source.EXPECT().Get("deprecated_key_transformed").Return("", true) // empty value
				source.EXPECT().KeyTransform("KEY").Return("key_transformed").Times(2)
				source.EXPECT().Get("key_transformed").Return("value", true)
				source.EXPECT().String().Return("A")
				return []Source{source}
			},
			key: "KEY",
			options: []Option{
				RetroKeys(func(source string, deprecateKey string, currentKey string) {},
					"DEPRECATED_KEY"),
			},
			value:      ptrTo("value"),
			sourceKind: "A",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)

			var sources []Source
			if testCase.makeSources != nil {
				sources = testCase.makeSources(ctrl)
			}

			value, sourceKind := get(sources, testCase.key, testCase.options...)
			if (value == nil && testCase.value != nil) ||
				(value != nil && testCase.value == nil) ||
				(value != nil && *value != *testCase.value) {
				t.Errorf("expected %v, got %v", testCase.value, value)
			}
			if sourceKind != testCase.sourceKind {
				t.Errorf("expected %s, got %s", testCase.sourceKind, sourceKind)
			}
		})
	}
}

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
			value: "\t\"'Value 'hello'\"'\t\n\t\r\n\t",
			settings: settings{
				trimLineEndings: ptrTo(true),
				trimSpace:       ptrTo(true),
				trimQuotes:      ptrTo(true),
				forceLowercase:  ptrTo(true),
			},
			result: "value 'hello'",
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
