package reader

import (
	"os"
	"reflect"
	"testing"
)

// setTestEnv is used to set environment variables in
// parallel tests.
// The key is created as keyPrefix + test name, so make
// sure to use different key prefixes if you need multiple
// environment variables in a single test.
func setTestEnv(t *testing.T, keyPrefix, value string) (key string) {
	t.Helper()
	key = keyPrefix + t.Name()
	existing, wasSet := os.LookupEnv(key)
	t.Cleanup(func() {
		var err error
		if wasSet {
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

func Test_New(t *testing.T) {
	t.Parallel()

	testKeys := make(map[string]struct{})

	emptyKey := setTestEnv(t, "EMPTY", "")
	testKeys[emptyKey] = struct{}{}

	filledKey := setTestEnv(t, "FILLED", "value")
	testKeys[filledKey] = struct{}{}

	notExistsKey := "NOTEXISTS" + t.Name()
	_, isSet := os.LookupEnv(notExistsKey)
	if isSet {
		t.Fatal("NOTEXISTS environment variable should not be set")
	}
	testKeys[notExistsKey] = struct{}{}

	settings := Settings{
		Environ: os.Environ(),
	}
	reader := New(settings)

	// Remove other test irrelevant environment variables
	for k := range reader.keyToValue {
		_, isTestKey := testKeys[k]
		if !isTestKey {
			delete(reader.keyToValue, k)
		}
	}

	if reader.handleDeprecatedKey == nil {
		t.Error("expected handleDeprecatedKey to be set")
	}
	reader.handleDeprecatedKey = nil

	expectedEnv := &Reader{
		keyToValue: map[string]string{
			emptyKey:  "",
			filledKey: "value",
		},
	}

	if !reflect.DeepEqual(expectedEnv, reader) {
		t.Errorf("expected: %v, got: %v", expectedEnv, reader)
	}
}
