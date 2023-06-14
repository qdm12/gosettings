package validate

import (
	"errors"
	"testing"
)

func Test_ListeningAddress(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		address                string
		uid                    int
		allowedPrivilegedPorts []uint16
		errWrapped             error
		noErrWrap              bool
		errMessage             string
	}{
		"empty": {},
		"missing_semicolon": {
			address:    "1.2.3.4",
			noErrWrap:  true,
			errMessage: "splitting host and port: address 1.2.3.4: missing port in address",
		},
		"port_not_integer": {
			address:    "1.2.3.4:x",
			errWrapped: ErrPortNotAnInteger,
			errMessage: "port value is not an integer: x",
		},
		"port_negative": {
			address:    "1.2.3.4:-1",
			errWrapped: ErrPortNotAnInteger,
			errMessage: "port value is not an integer: -1",
		},
		"port_too_high": {
			address:    "1.2.3.4:65537",
			errWrapped: ErrPortTooHigh,
			errMessage: "port cannot be higher than 65535: 65537",
		},
		"zero_port": {
			address: "1.2.3.4:0",
		},
		"privileged_port_without_root": {
			address:    "1.2.3.4:100",
			uid:        1000,
			errWrapped: ErrPrivilegedPort,
			errMessage: "cannot use privileged ports (1 to 1023) when running without root: 100",
		},
		"allowed_privileged_port_without_root": {
			address:                "1.2.3.4:100",
			uid:                    1000,
			allowedPrivilegedPorts: []uint16{99, 100},
		},
		"privileged_port_with_root": {
			address: "1.2.3.4:100",
			uid:     0,
		},
		"valid_address": {
			address: "1.2.3.4:8000",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := ListeningAddress(testCase.address, testCase.uid, testCase.allowedPrivilegedPorts...)

			if !testCase.noErrWrap && !errors.Is(err, testCase.errWrapped) {
				t.Fatalf("expected error %q to be wrapped in %s", testCase.errWrapped, err)
			}
			if testCase.errWrapped != nil &&
				testCase.errMessage != err.Error() {
				t.Errorf("expected error %q but got %q", testCase.errMessage, err)
			}
		})
	}
}
