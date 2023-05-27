package gosettings

func ptrTo[T any](t T) *T { return &t }

type testInterface interface {
	F()
}

type testInterfaceImplA struct{}

func (testInterfaceImplA) F() {}

type testInterfaceImplB struct{}

func (testInterfaceImplB) F() {}
