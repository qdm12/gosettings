package reader

import (
	"reflect"
	"testing"

	"github.com/qdm12/gosettings/internal/parse"
)

func Test_New(t *testing.T) {
	t.Parallel()

	testSourceA := &testSource{
		keyValue: map[string]string{},
	}
	testSourceB := &testSource{
		keyValue: map[string]string{},
	}

	readerSettings := Settings{
		Sources: []Source{testSourceA, testSourceB},
	}

	reader := New(readerSettings)

	if reader.handleDeprecatedKey == nil {
		t.Error("handleDeprecatedKey should not be nil")
	}
	reader.handleDeprecatedKey = nil

	expectedReader := &Reader{
		sources: []parse.Source{testSourceA, testSourceB},
		defaultReadSettings: settings{
			forceLowercase: ptrTo(false),
			acceptEmpty:    ptrTo(false),
		},
	}

	if !reflect.DeepEqual(expectedReader, reader) {
		t.Errorf("expected: %#v, got: %#v", expectedReader, reader)
	}
}
