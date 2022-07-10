package defaults

func StringSlice(existing []string, defaultValue []string) (
	result []string) {
	if existing != nil {
		return existing
	}
	return defaultValue
}
