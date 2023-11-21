package reader

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_NewEnv(t *testing.T) {
	t.Parallel()

	testKeys := make(map[string]struct{})

	emptyKey := setTestEnv(t, "EMPTY", "")
	testKeys[emptyKey] = struct{}{}

	filledKey := setTestEnv(t, "FILLED", "value")
	testKeys[filledKey] = struct{}{}

	lowercaseKey := setTestEnv(t, "lowercase", "value2")
	transformedLowercaseKey := strings.ToUpper(lowercaseKey)
	testKeys[transformedLowercaseKey] = struct{}{}

	env := NewEnv(os.Environ())

	// Remove other test irrelevant environment variables
	for k := range env.keyToValue {
		_, isTestKey := testKeys[k]
		if !isTestKey {
			delete(env.keyToValue, k)
		}
	}

	expectedEnv := &Env{
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

	env := &Env{}
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
		env   *Env
		key   string
		value string
		isSet bool
	}{
		"key_not_found": {
			env: NewEnv([]string{}),
			key: "KEY",
		},
		"empty_value": {
			env:   NewEnv([]string{"KEY="}),
			key:   "KEY",
			isSet: true,
		},
		"non_empty_value": {
			env:   NewEnv([]string{"KEY=value"}),
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

			env := &Env{}
			newKey := env.KeyTransform(testCase.key)
			if newKey != testCase.newKey {
				t.Errorf("expected %s, got %s", testCase.newKey, newKey)
			}
		})
	}
}
