package defaults

import "net"

func IP(existing net.IP, defaultValue net.IP) (
	result net.IP) {
	if existing != nil {
		return existing
	}
	return defaultValue
}
