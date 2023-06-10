package validate

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

// ListeningAddress validates a listening address string given a user ID `uid`
// and an optional list of allowed privileged ports.
func ListeningAddress(address string, uid int, allowedPrivilegedPorts ...uint16) (err error) {
	_, portStr, err := net.SplitHostPort(address)
	if err != nil {
		return fmt.Errorf("splitting host and port: %w", err)
	}

	return validatePort(portStr, uid, allowedPrivilegedPorts)
}

var (
	ErrPortNotAnInteger = errors.New("port value is not an integer")
	ErrPortTooHigh      = errors.New("port cannot be higher than 65535")
	ErrPrivilegedPort   = errors.New("cannot use privileged ports (1 to 1023) when running without root")
)

func validatePort(value string, uid int, allowedPrivilegedPorts []uint16) (err error) {
	const maxPort = 65535

	const base, bitSize = 10, 64
	portUint, err := strconv.ParseUint(value, base, bitSize)
	switch {
	case err != nil:
		return fmt.Errorf("%w: %s", ErrPortNotAnInteger, value)
	case portUint > maxPort:
		return fmt.Errorf("%w: %d", ErrPortTooHigh, portUint)
	}

	port := uint16(portUint)
	const (
		maxPrivilegedPort = 1023
		minDynamicPort    = 49151
	)
	if port == 0 || port > maxPrivilegedPort ||
		uid == -1 || uid == 0 { // windows or root uid
		return nil
	}

	for _, allowed := range allowedPrivilegedPorts {
		if allowed == port {
			return nil
		}
	}
	return fmt.Errorf("%w: %d", ErrPrivilegedPort, port)
}
