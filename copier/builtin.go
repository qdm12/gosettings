package copier

func Bool(original *bool) (copied *bool) {
	if original == nil {
		return nil
	}
	copied = new(bool)
	*copied = *original
	return copied
}

func Int(original *int) (copied *int) {
	if original == nil {
		return nil
	}
	copied = new(int)
	*copied = *original
	return copied
}

func Uint(original *uint) (copied *uint) {
	if original == nil {
		return nil
	}
	copied = new(uint)
	*copied = *original
	return copied
}

func Uint8(original *uint8) (copied *uint8) {
	if original == nil {
		return nil
	}
	copied = new(uint8)
	*copied = *original
	return copied
}

func Uint16(original *uint16) (copied *uint16) {
	if original == nil {
		return nil
	}
	copied = new(uint16)
	*copied = *original
	return copied
}

func Uint32(original *uint32) (copied *uint32) {
	if original == nil {
		return nil
	}
	copied = new(uint32)
	*copied = *original
	return copied
}

func String(original *string) (copied *string) {
	if original == nil {
		return nil
	}
	copied = new(string)
	*copied = *original
	return copied
}
