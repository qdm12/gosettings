package gosettings

import (
	"net/netip"
	"testing"
)

func Test_DefaultValidator(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		existing     SelfValidator
		defaultValue SelfValidator
		result       SelfValidator
	}{
		"netip.Addr keep existing": {
			existing:     netip.AddrFrom4([4]byte{1, 2, 3, 4}),
			defaultValue: netip.AddrFrom4([4]byte{5, 6, 7, 8}),
			result:       netip.AddrFrom4([4]byte{1, 2, 3, 4}),
		},
		"netip.Addr use default existing": {
			existing:     netip.Addr{},
			defaultValue: netip.AddrFrom4([4]byte{5, 6, 7, 8}),
			result:       netip.AddrFrom4([4]byte{5, 6, 7, 8}),
		},
		"netip.Prefix keep existing": {
			existing:     netip.PrefixFrom(netip.AddrFrom4([4]byte{1, 2, 3, 4}), 32),
			defaultValue: netip.PrefixFrom(netip.AddrFrom4([4]byte{5, 6, 7, 8}), 32),
			result:       netip.PrefixFrom(netip.AddrFrom4([4]byte{1, 2, 3, 4}), 32),
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := DefaultValidator(testCase.existing, testCase.defaultValue)
			if result != testCase.result {
				t.Fatalf("expected %v, got %v", testCase.result, result)
			}
		})
	}
}
