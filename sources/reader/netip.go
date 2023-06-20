package reader

import (
	"net/netip"

	"github.com/qdm12/gosettings/sources/parse"
)

// NetipAddr returns a netip.Addr from the value found at the given key.
// If the value is not a valid netip.Addr string, an error is returned
// with the key in its message.
// The value is returned as the empty invalid `netip.Addr{}` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AllowEmpty option, if the
//     given key is set and its corresponding value is empty.
func (r *Reader) NetipAddr(key string, options ...Option) (
	addr netip.Addr, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.NetipAddr(r.keyToValue, key, parseOptions...)
}

// CSVNetipAddresses returns a slice of netip.Addr from a comma separated value
// found at the given key and returns an error if any value
// is not a valid netip.Addr string.
//
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
//
// The entire CSV string value may be modified depending on the
// parse default settings and the parse options given.
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func (r *Reader) CSVNetipAddresses(key string, options ...Option) (
	prefixes []netip.Addr, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVNetipAddresses(r.keyToValue, key, parseOptions...)
}

// CSVNetipPrefixes returns a slice of netip.Prefix from a comma separated value
// found at the given key and returns an error if any value
// is not a valid netip.Prefix string.
//
// The slice is returned as `nil` if:
//   - the given key is NOT set.
//   - By default and unless changed by the AcceptEmpty option,
//     if the key is set and its corresponding value is empty.
//
// The entire CSV string value may be modified depending on the
// parse default settings and the parse options given.
// The parse default settings are to:
//   - Trim line endings suffixes \r\n and \n.
//   - Trim spaces.
//   - Trim quotes.
//   - Force lowercase.
func (r *Reader) CSVNetipPrefixes(key string, options ...Option) (
	prefixes []netip.Prefix, err error) {
	parseOptions := r.makeParseOptions(options)
	return parse.CSVNetipPrefixes(r.keyToValue, key, parseOptions...)
}
