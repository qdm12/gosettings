package gosettings

import (
	"testing"
)

func Test_ObfuscateKey(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		key           string
		obfuscatedKey string
	}{
		"empty": {
			obfuscatedKey: "[not set]",
		},
		"124_bits_security": {
			// 16^31 gives 124 bits security hidden
			key:           "0123456789abcdef0123456789abcde",
			obfuscatedKey: "[set]",
		},
		"128_bits_security": {
			// 16^32 gives 128 bits security hidden
			key:           "0123456789abcdef0123456789abcdef",
			obfuscatedKey: "[set]",
		},
		"132_bits_security": {
			// 16^33 gives 128 bits security hidden + 1 shown character
			key:           "0123456789abcdef0123456789abcdef0",
			obfuscatedKey: "[set]",
		},
		"136_bits_security": {
			// 16^34 gives 128 bits security hidden + 2 shown characters
			key:           "0123456789abcdef0123456789abcdef01",
			obfuscatedKey: "0...1",
		},
		"296_bits_security": {
			// 52^52 combinations ~= 2^296
			key:           "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			obfuscatedKey: "abc...789",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			obfuscatedKey := ObfuscateKey(testCase.key)

			if obfuscatedKey != testCase.obfuscatedKey {
				t.Errorf("expected %s, got %s", testCase.obfuscatedKey, obfuscatedKey)
			}
		})
	}
}
