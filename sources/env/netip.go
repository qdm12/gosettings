package env

import (
	"net/netip"

	"github.com/qdm12/gosettings/sources/parse"
)

// NetipAddr returns a netip.Addr from an environment variable value.
// If the value is not a valid netip.Addr string, an error is returned
// with the environment variable name in its message.
// The value is returned as the empty invalid `netip.Addr{}` if:
//   - the environment variable key given is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     environment variable is set and its value is empty.
func (e *Env) NetipAddr(envKey string, options ...Option) (
	addr netip.Addr, err error) {
	parseOptions := e.makeParseOptions(options)
	return parse.NetipAddr(e.environ, envKey, parseOptions...)
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
	parseOptions := e.makeParseOptions(options)
	return parse.CSVNetipAddresses(e.environ, envKey, parseOptions...)
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
	parseOptions := e.makeParseOptions(options)
	return parse.CSVNetipPrefixes(e.environ, envKey, parseOptions...)
}
