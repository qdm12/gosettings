package env

import (
	"fmt"
	"net/netip"
)

// NetipAddr returns a netip.Addr from an environment variable value.
// If the environment variable is not set or its value is
// the empty string, the empty invalid `netip.Addr{}` is returned.
// Otherwise, if the value is not a valid ip string,
// an error is returned with the environment variable name
// in its message.
func (e *Env) NetipAddr(envKey string, options ...Option) (
	addr netip.Addr, err error) {
	s := e.Get(envKey, options...)
	if s == nil || *s == "" {
		// note: no point accepting the empty string in this case
		return addr, nil
	}

	addr, err = netip.ParseAddr(*s)
	if err != nil {
		return addr, fmt.Errorf("environment variable %s: %w", envKey, err)
	}

	return addr, nil
}

// CSVNetipAddresses returns a slice of netip.Addr from a comma separated
// environment variable value and returns an error if any value
// is not a valid netip.Addr string.
//
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
//
// The entire CSV string value may be modified depending on the
// parse default settings and the parse options given.
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func (e *Env) CSVNetipAddresses(envKey string, options ...Option) (
	prefixes []netip.Addr, err error) {
	return csvParse(e, envKey, netip.ParseAddr, options...)
}

// CSVNetipPrefixes returns a slice of netip.Prefix from a comma separated
// environment variable value and returns an error if any value
// is not a valid netip.Prefix string.
//
// The slice is returned as `nil` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the environment variable is set and its value is empty.
//
// The entire CSV string value may be modified depending on the
// parse default settings and the parse options given.
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func (e *Env) CSVNetipPrefixes(envKey string, options ...Option) (
	prefixes []netip.Prefix, err error) {
	return csvParse(e, envKey, netip.ParsePrefix, options...)
}
