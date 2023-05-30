package merger

import (
	"errors"
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	t.Parallel()

	sourceA := &sourceA{}
	sourceB := &sourceB{}

	merger := New[settings](sourceA, sourceB)

	expectedMerger := &Merger[settings]{
		sources: []Source[settings]{sourceA, sourceB},
	}
	if !reflect.DeepEqual(merger, expectedMerger) {
		t.Fatalf("expected merger to be %v, got %v", expectedMerger, merger)
	}
}

func Test_Merger_String(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		merger   *Merger[settings]
		expected string
	}{
		"no_sources": {
			merger:   &Merger[settings]{},
			expected: "merger",
		},
		"one_source": {
			merger: &Merger[settings]{
				sources: []Source[settings]{&sourceA{}},
			},
			expected: "merger of A",
		},
		"two_sources": {
			merger: &Merger[settings]{
				sources: []Source[settings]{&sourceA{}, &sourceB{}},
			},
			expected: "merger of A, B",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := testCase.merger.String()

			if actual != testCase.expected {
				t.Fatalf("expected merger string to be %q, got %q", testCase.expected, actual)
			}
		})
	}
}

func Test_Merger_Read(t *testing.T) {
	t.Parallel()

	errTest := errors.New("test error")

	testCases := map[string]struct {
		merger   *Merger[settings]
		settings settings
		err      error
	}{
		"no_sources": {
			merger: &Merger[settings]{},
		},
		"one_source": {
			merger: &Merger[settings]{
				sources: []Source[settings]{&sourceA{}},
			},
		},
		"two_sources": {
			merger: &Merger[settings]{
				sources: []Source[settings]{&sourceA{}, &sourceB{}},
			},
			settings: settings{x: 1},
		},
		"two_sources_error": {
			merger: &Merger[settings]{
				sources: []Source[settings]{&sourceError{err: errTest}, &sourceB{}},
			},
			err: errTest,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			settings, err := testCase.merger.Read()

			if !reflect.DeepEqual(settings, testCase.settings) {
				t.Fatalf("expected merged settings to be %v, got %v", testCase.settings, settings)
			}
			if !errors.Is(err, testCase.err) {
				t.Fatalf("expected error to be %v, got %v", testCase.err, err)
			}
		})
	}
}
