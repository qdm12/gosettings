package gosettings

import (
	"net/netip"
	"testing"
)

func Test_DefaultInterface(t *testing.T) {
	t.Parallel()

	t.Run("no_type_assertion_needed", func(t *testing.T) {
		t.Parallel()

		var existing testInterface
		existing = DefaultInterface(existing, &testInterfaceImplA{})
		existing.F() // make use of variable
	})

	testCases := map[string]struct {
		existing     any
		defaultValue any
		result       any
	}{
		"all_untyped_nils": {},
		"all_typed_nils": {
			existing:     (testInterface)(nil),
			defaultValue: (testInterface)(nil),
			result:       (testInterface)(nil),
		},
		"nil_default": {
			existing: &testInterfaceImplA{},
			result:   &testInterfaceImplA{},
		},
		"no_default": {
			existing:     &testInterfaceImplA{},
			defaultValue: &testInterfaceImplB{},
			result:       &testInterfaceImplA{},
		},
		"default": {
			existing:     (testInterface)(nil),
			defaultValue: &testInterfaceImplB{},
			result:       &testInterfaceImplB{},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := DefaultInterface(testCase.existing, testCase.defaultValue)
			if result != testCase.result {
				t.Errorf("expected %v, got %v", testCase.result, result)
			}
		})
	}
}

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
