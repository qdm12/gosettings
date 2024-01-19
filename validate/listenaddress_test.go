package validate

import (
	"errors"
	"fmt"
	"testing"
)

func Test_ListeningAddress(t *testing.T) {
	t.Parallel()

	unprivilegedPortStart, err := getUnprivilegedPortStart()
	if err != nil {
		t.Fatal(err)
	}

	type testCaseStruct struct {
		address    string
		uid        int
		errWrapped error
		noErrWrap  bool
		errMessage string
	}

	testCases := map[string]testCaseStruct{
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
		"valid_address": {
			address: "1.2.3.4:8000",
		},
	}

	// Only add test cases testing privileged ports if there is at least
	// one privileged port in the system.
	if unprivilegedPortStart > 1 {
		lastPrivilegedPort := unprivilegedPortStart - 1
		testCases["privileged_port_without_root"] = testCaseStruct{
			address:    fmt.Sprintf("1.2.3.4:%d", lastPrivilegedPort),
			uid:        1000,
			errWrapped: ErrPrivilegedPort,
			errMessage: fmt.Sprintf("listening on privileged port is not allowed: "+
				"port %d (user id 1000, unprivileged start port %d)",
				lastPrivilegedPort, unprivilegedPortStart),
		}
		testCases["allowed_privileged_port_without_root"] = testCaseStruct{
			address: fmt.Sprintf("1.2.3.4:%d", lastPrivilegedPort),
			uid:     1000,
		}
		testCases["privileged_port_with_root"] = testCaseStruct{
			address: fmt.Sprintf("1.2.3.4:%d", lastPrivilegedPort),
			uid:     0,
		}
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := ListeningAddress(testCase.address, testCase.uid)

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

func Test_makePrivilegedPortError(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		port                   uint16
		uid                    int
		unprivilegedPortStart  uint16
		allowedPrivilegedPorts []uint16
		errMessage             string
	}{
		"least_details": {
			port:                  1,
			uid:                   1000,
			unprivilegedPortStart: 1024,
			errMessage: "listening on privileged port is not allowed: " +
				"port 1 (user id 1000)",
		},
		"max_details": {
			port:                   1,
			uid:                    1000,
			unprivilegedPortStart:  500,
			allowedPrivilegedPorts: []uint16{2, 3},
			errMessage: "listening on privileged port is not allowed: " +
				"port 1 (user id 1000, unprivileged start port 500)",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := makePrivilegedPortError(testCase.port, testCase.uid,
				testCase.unprivilegedPortStart)

			if !errors.Is(err, ErrPrivilegedPort) {
				t.Fatalf("expected error %q to be wrapped in %s", ErrPrivilegedPort, err)
			}
			if testCase.errMessage != err.Error() {
				t.Errorf("expected error %q but got %q", testCase.errMessage, err)
			}
		})
	}
}
