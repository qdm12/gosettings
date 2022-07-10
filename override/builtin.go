package override

func Int(existing, other int) (result int) {
	if other == 0 {
		return existing
	}
	return other
}

func Float64(existing, other float64) (result float64) {
	if other == 0 {
		return existing
	}
	return other
}

func String(existing, other string) (result string) {
	if other == "" {
		return existing
	}
	return other
}

func Bool(existing, other *bool) (result *bool) {
	if other == nil {
		return existing
	}
	result = new(bool)
	*result = *other
	return result
}

func IntPtr(existing, other *int) (result *int) {
	if other == nil {
		return existing
	}
	result = new(int)
	*result = *other
	return result
}

func Uint8Ptr(existing, other *uint8) (result *uint8) {
	if other == nil {
		return existing
	}
	result = new(uint8)
	*result = *other
	return result
}

func Uint16Ptr(existing, other *uint16) (result *uint16) {
	if other == nil {
		return existing
	}
	result = new(uint16)
	*result = *other
	return result
}

func Uint32Ptr(existing, other *uint32) (result *uint32) {
	if other == nil {
		return existing
	}
	result = new(uint32)
	*result = *other
	return result
}

func StringPtr(existing, other *string) (result *string) {
	if other == nil {
		return existing
	}
	result = new(string)
	*result = *other
	return result
}

func Uint16Slice(existing, other []uint16) (result []uint16) {
	if other == nil {
		return existing
	}
	result = make([]uint16, len(other))
	copy(result, other)
	return result
}

func StringSlice(existing, other []string) (result []string) {
	if other == nil {
		return existing
	}
	result = make([]string, len(other))
	copy(result, other)
	return result
}
