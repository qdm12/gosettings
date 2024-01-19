package validate

import (
	"testing"
)

func Test_getUnprivilegedPortStart(t *testing.T) {
	t.Parallel()

	_, err := getUnprivilegedPortStart()
	if err != nil {
		t.Fatal(err)
	}
}
