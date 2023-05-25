package gosettings

type testInterface interface {
	F()
}

type testInterfaceImplA struct{}

func (testInterfaceImplA) F() {}

type testInterfaceImplB struct{}

func (testInterfaceImplB) F() {}
