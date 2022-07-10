package defaults

func Int(existing int, defaultValue int) (
	result int) {
	if existing != 0 {
		return existing
	}
	return defaultValue
}

func String(existing string, defaultValue string) (
	result string) {
	if existing != "" {
		return existing
	}
	return defaultValue
}

func Bool(existing *bool, defaultValue bool) (
	result *bool) {
	if existing != nil {
		return existing
	}
	result = new(bool)
	*result = defaultValue
	return result
}

func IntPtr(existing *int, defaultValue int) (
	result *int) {
	if existing != nil {
		return existing
	}
	result = new(int)
	*result = defaultValue
	return result
}

func Uint8Ptr(existing *uint8, defaultValue uint8) (
	result *uint8) {
	if existing != nil {
		return existing
	}
	result = new(uint8)
	*result = defaultValue
	return result
}

func Uint16Ptr(existing *uint16, defaultValue uint16) (
	result *uint16) {
	if existing != nil {
		return existing
	}
	result = new(uint16)
	*result = defaultValue
	return result
}
func Uint32Ptr(existing *uint32, defaultValue uint32) (
	result *uint32) {
	if existing != nil {
		return existing
	}
	result = new(uint32)
	*result = defaultValue
	return result
}

func StringPtr(existing *string, defaultValue string) (result *string) {
	if existing != nil {
		return existing
	}
	result = new(string)
	*result = defaultValue
	return result
}
