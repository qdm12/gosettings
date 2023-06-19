package parse

import (
	"net/netip"
)

// NetipAddr returns a netip.Addr from the value found at the given
// key in the keyValues map.
// If the value is not a valid netip.Addr string, an error is returned
// with the key in its message.
// The value is returned as the empty invalid `netip.Addr{}` if:
//   - the key given is NOT set in the keyValues map.
//   - By default and unless changed by the AllowEmpty option, if the
//     key is set and its corresponding value is empty.
func NetipAddr(keyValues map[string]string, key string,
	options ...Option) (addr netip.Addr, err error) {
	return GetParse(keyValues, key, netip.ParseAddr, options...)
}

// CSVNetipAddresses returns a slice of netip.Addr from a comma separated
// string value found at the given key in the given keyValues map, and
// returns an error if any value is not a valid netip.Addr string.
//
// The slice is returned as `nil` if:
//   - the key given is NOT set in the keyValues map.
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
func CSVNetipAddresses(keyValues map[string]string, key string,
	options ...Option) (prefixes []netip.Addr, err error) {
	return CSVParse(keyValues, key, netip.ParseAddr, options...)
}

// CSVNetipPrefixes returns a slice of netip.Prefix from a comma separated
// string value found at the given key in the given keyValues map, and
// returns an error if any value is not a valid netip.Prefix string.
//
// The slice is returned as `nil` if:
//   - the key given is NOT set in the keyValues map.
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
func CSVNetipPrefixes(keyValues map[string]string, key string,
	options ...Option) (prefixes []netip.Prefix, err error) {
	return CSVParse(keyValues, key, netip.ParsePrefix, options...)
}
