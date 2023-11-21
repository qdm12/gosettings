package reader

import (
	"os"
	"strings"
	"testing"
)

func ptrTo[T any](x T) *T { return &x }

type testSource struct {
	keyValue map[string]string
}

func (t *testSource) String() string { return "test" }

func (t *testSource) Get(key string) (value string, isSet bool) {
	value, isSet = t.keyValue[key]
	return value, isSet
}

func (t *testSource) KeyTransform(key string) string { return key }

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
