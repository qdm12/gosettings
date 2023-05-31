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
// in the error context.
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
