package reader

func ptrTo[T any](x T) *T { return &x }
