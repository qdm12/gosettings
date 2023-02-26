package merge

import "net"

func IP(existing, other net.IP) (result net.IP) {
	if existing != nil {
		return existing
	} else if other == nil {
		return nil
	}
	result = make(net.IP, len(other))
	copy(result, other)
	return result
}

func IPNetsSlice(a, b []net.IPNet) (result []net.IPNet) {
	if a == nil && b == nil {
		return nil
	}

	seen := make(map[string]struct{}, len(a)+len(b))
	result = make([]net.IPNet, 0, len(a)+len(b))
	for _, ipNet := range a {
		key := ipNet.String()
		if _, ok := seen[key]; ok {
			continue // duplicate
		}
		result = append(result, ipNet)
		seen[key] = struct{}{}
	}
	for _, ipNet := range b {
		key := ipNet.String()
		if _, ok := seen[key]; ok {
			continue // duplicate
		}
		result = append(result, ipNet)
		seen[key] = struct{}{}
	}
	return result
}
