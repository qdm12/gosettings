package reader

func ptrTo[T any](x T) *T { return &x }

type testSource struct {
	keyValue map[string]string
}

func (t *testSource) String() string { return "test" }

func (t *testSource) Get(key string) (value string, isSet bool) {
	value, isSet = t.keyValue[key]
	return value, isSet
}

func (t *testSource) KeyTransform(key string) string { return key }
