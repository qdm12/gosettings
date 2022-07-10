package override

import "net"

func IP(existing, other net.IP) (result net.IP) {
	if other == nil {
		return existing
	}
	result = make(net.IP, len(other))
	copy(result, other)
	return result
}
