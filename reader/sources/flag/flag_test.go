package flag

import (
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		osArgs []string
		source *Source
	}{
		"empty": {
			osArgs: []string{"program"},
			source: &Source{
				keyToValue: map[string]string{},
			},
		},
		"commands_no_flags": {
			osArgs: []string{"program", "do-this", "with-that"},
			source: &Source{
				keyToValue: map[string]string{},
			},
		},
		"command_then_flag": {
			osArgs: []string{"program", "do-this", "-log=yes"},
			source: &Source{
				keyToValue: map[string]string{
					"log": "yes",
				},
			},
		},
		"flag_then_command": {
			osArgs: []string{"program", "-log=yes", "do-this"},
			source: &Source{
				keyToValue: map[string]string{
					"log": "yes",
				},
			},
		},
		"last_flag_override": {
			osArgs: []string{"program", "-log=yes", "do-this", "-log=no"},
			source: &Source{
				keyToValue: map[string]string{
					"log": "no",
				},
			},
		},
		"mix": {
			osArgs: []string{"program", "-log=yes",
				"do-this", "-log=no", "--metrics", "prometheus",
				"--rpc", "--port", "8000", "not-a-flag",
			},
			source: &Source{
				keyToValue: map[string]string{
					"log":     "no",
					"metrics": "prometheus",
					"rpc":     "true",
					"port":    "8000",
				},
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			source := New(testCase.osArgs)

			if !reflect.DeepEqual(source, testCase.source) {
				t.Errorf("expected source %#v, got %#v", testCase.source, source)
			}
		})
	}
}

func Test_isFlag(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		osArg string
		ok    bool
	}{
		"empty": {},
		"one_character": {
			osArg: "-",
		},
		"no_dash_prefix": {
			osArg: "abc",
		},
		"two_dashes_only": {
			osArg: "--",
		},
		"dash_equal": {
			osArg: "-=a",
		},
		"three_dashes": {
			osArg: "---a",
		},
		"two_dashes_one_equal": {
			osArg: "--=a",
		},
		"valid_single_dash": {
			osArg: "-a",
			ok:    true,
		},
		"valid_two_dashes": {
			osArg: "--a",
			ok:    true,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ok := isFlag(testCase.osArg)

			if ok != testCase.ok {
				t.Errorf("expected ok %t, got %t", testCase.ok, ok)
			}
		})
	}
}
