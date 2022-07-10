package copier

import "net"

func IP(original net.IP) (copied net.IP) {
	if original == nil {
		return nil
	}
	copied = make(net.IP, len(original))
	copy(copied, original)
	return copied
}

func IPNet(original net.IPNet) (copied net.IPNet) {
	if original.IP != nil {
		copied.IP = make(net.IP, len(original.IP))
		copy(copied.IP, original.IP)
	}

	if original.Mask != nil {
		copied.Mask = make(net.IPMask, len(original.Mask))
		copy(copied.Mask, original.Mask)
	}

	return copied
}

func IPNetPtr(original *net.IPNet) (copied *net.IPNet) {
	if original == nil {
		return nil
	}

	copied = new(net.IPNet)
	*copied = IPNet(*original)
	return copied
}
