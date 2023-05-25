//go:build go1.18
// +build go1.18

package gosettings

import "golang.org/x/exp/constraints"

// Number represents a number type, it can notably be
// an integer, a float, and any type aliases of them such
// as `time.Duration`.
type Number interface {
	constraints.Integer | constraints.Float
}
