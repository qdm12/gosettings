package merge

func Int(existing, other int) (result int) {
	if existing != 0 {
		return existing
	}
	return other
}

func String(existing, other string) (result string) {
	if existing != "" {
		return existing
	}
	return other
}

func Float64(existing, other float64) (result float64) {
	if existing != 0 {
		return existing
	}
	return other
}

func Bool(existing, other *bool) (result *bool) {
	if existing != nil {
		return existing
	} else if other == nil {
		return nil
	}
	result = new(bool)
	*result = *other
	return result
}

func IntPtr(existing, other *int) (result *int) {
	if existing != nil {
		return existing
	} else if other == nil {
		return nil
	}
	result = new(int)
	*result = *other
	return result
}

func Uint8Ptr(existing, other *uint8) (result *uint8) {
	if existing != nil {
		return existing
	} else if other == nil {
		return nil
	}
	result = new(uint8)
	*result = *other
	return result
}

func Uint16Ptr(existing, other *uint16) (result *uint16) {
	if existing != nil {
		return existing
	} else if other == nil {
		return nil
	}
	result = new(uint16)
	*result = *other
	return result
}

func Uint32Ptr(existing, other *uint32) (result *uint32) {
	if existing != nil {
		return existing
	} else if other == nil {
		return nil
	}
	result = new(uint32)
	*result = *other
	return result
}

func StringPtr(existing, other *string) (result *string) {
	if existing != nil {
		return existing
	} else if other == nil {
		return nil
	}
	result = new(string)
	*result = *other
	return result
}

func Uint16Slice(a, b []uint16) (result []uint16) {
	if a == nil && b == nil {
		return nil
	}

	seen := make(map[uint16]struct{}, len(a)+len(b))
	result = make([]uint16, 0, len(a)+len(b))
	for _, n := range a {
		if _, ok := seen[n]; ok {
			continue // duplicate
		}
		result = append(result, n)
		seen[n] = struct{}{}
	}
	for _, n := range b {
		if _, ok := seen[n]; ok {
			continue // duplicate
		}
		result = append(result, n)
		seen[n] = struct{}{}
	}
	return result
}

func StringSlice(a, b []string) (result []string) {
	if a == nil && b == nil {
		return nil
	}

	seen := make(map[string]struct{}, len(a)+len(b))
	result = make([]string, 0, len(a)+len(b))
	for _, s := range a {
		if _, ok := seen[s]; ok {
			continue // duplicate
		}
		result = append(result, s)
		seen[s] = struct{}{}
	}
	for _, s := range b {
		if _, ok := seen[s]; ok {
			continue // duplicate
		}
		result = append(result, s)
		seen[s] = struct{}{}
	}
	return result
}
