package env

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_New(t *testing.T) {
	t.Parallel()

	testKeys := make(map[string]struct{})

	emptyKey := setTestEnv(t, "EMPTY", "")
	testKeys[emptyKey] = struct{}{}

	filledKey := setTestEnv(t, "FILLED", "value")
	testKeys[filledKey] = struct{}{}

	lowercaseKey := setTestEnv(t, "lowercase", "value2")
	transformedLowercaseKey := strings.ToUpper(lowercaseKey)
	testKeys[transformedLowercaseKey] = struct{}{}

	settings := Settings{
		Environ: os.Environ(),
	}

	env := New(settings)

	// Remove other test irrelevant environment variables
	for k := range env.keyToValue {
		_, isTestKey := testKeys[k]
		if !isTestKey {
			delete(env.keyToValue, k)
		}
	}

	expectedEnv := &Source{
		keyToValue: map[string]string{
			emptyKey:                "",
			filledKey:               "value",
			transformedLowercaseKey: "value2",
		},
	}

	if !reflect.DeepEqual(expectedEnv, env) {
		t.Errorf("expected: %#v, got: %#v", expectedEnv, env)
	}
}

func Test_Env_String(t *testing.T) {
	t.Parallel()

	env := &Source{}
	someErrorMessage := fmt.Sprintf("%s %s: %s",
		env, "ENV_KEY", "some problem")
	const expectedErrorMessage = "environment variable ENV_KEY: " +
		"some problem"
	if someErrorMessage != expectedErrorMessage {
		t.Errorf("expected %s, got %s", expectedErrorMessage,
			someErrorMessage)
	}
}

func Test_Env_Get(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		env   *Source
		key   string
		value string
		isSet bool
	}{
		"key_not_found": {
			env: New(Settings{Environ: []string{}}),
			key: "KEY",
		},
		"empty_value": {
			env:   New(Settings{Environ: []string{"KEY="}}),
			key:   "KEY",
			isSet: true,
		},
		"non_empty_value": {
			env:   New(Settings{Environ: []string{"KEY=value"}}),
			key:   "KEY",
			value: "value",
			isSet: true,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			value, isSet := testCase.env.Get(testCase.key)

			if value != testCase.value {
				t.Errorf("expected value %s, got %s", testCase.value, value)
			}
			if isSet != testCase.isSet {
				t.Errorf("expected isSet %t, got %t", testCase.isSet, isSet)
			}
		})
	}
}

func Test_Env_KeyTransform(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		key    string
		newKey string
	}{
		"empty": {},
		"no_change": {
			key:    "ENV_KEY",
			newKey: "ENV_KEY",
		},
		"flag_like": {
			key:    "env-key",
			newKey: "ENV_KEY",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			env := &Source{}
			newKey := env.KeyTransform(testCase.key)
			if newKey != testCase.newKey {
				t.Errorf("expected %s, got %s", testCase.newKey, newKey)
			}
		})
	}
}

// setTestEnv is used to set environment variables in
// parallel tests, and restores or clears set variables
// when the test finishes.
// The environment variable key is computed as:
// keyPrefix + "_" + UPPER(test name)
// and is returned to the caller.
func setTestEnv(t *testing.T, keyPrefix, value string) (key string) {
	t.Helper()
	keySuffix := "_" + strings.ToUpper(t.Name())
	key = keyPrefix + keySuffix
	existing, wasSet := os.LookupEnv(key)

	t.Cleanup(func() {
		var err error
		if wasSet {
			if strings.HasSuffix(existing, keySuffix) {
				// the cleanup associated with the first set
				// will take care of cleaning up.
				return
			}
			err = os.Setenv(key, existing)
		} else {
			err = os.Unsetenv(key)
		}
		if err != nil {
			t.Error(err)
		}
	})
	err := os.Setenv(key, value) //nolint:tenv
	if err != nil {
		t.Fatal(err)
	}
	return key
}
