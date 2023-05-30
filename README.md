# Gosettings

`gosettings` is a Go package providing **helper functions for working with settings**.

**Go.dev documentation:**

- [gosettings](https://pkg.go.dev/github.com/qdm12/gosettings)
- [gosettings/validate](https://pkg.go.dev/github.com/qdm12/gosettings/validate)
- [gosettings/sources/env](https://pkg.go.dev/github.com/qdm12/gosettings/sources/env)

Add it to your Go project with:

```sh
go get github.com/qdm12/gosettings
```

💁 Only compatible with Go 1.18+ since it now uses generics.

Features:

- Define settings struct methods:
  - `Copy`: use `gosettings.CopyPointer` and `gosettings.CopySlice`
  - `SetDefaults`: `gosettings.Default*` functions (see [pkg.go.dev/github.com/qdm12/gosettings](https://pkg.go.dev/github.com/qdm12/gosettings))
  - `MergeWith`: `gosettings.MergeWith*` functions (see [pkg.go.dev/github.com/qdm12/gosettings](https://pkg.go.dev/github.com/qdm12/gosettings))
  - `OverrideWith`: `gosettings.OverrideWith*` functions (see [pkg.go.dev/github.com/qdm12/gosettings](https://pkg.go.dev/github.com/qdm12/gosettings))
  - `Validate`: `validate.*` functions from [`github.com/qdm12/gosettings/validate`](https://pkg.go.dev/github.com/qdm12/gosettings/validate)
- Reading settings from sources:
  - Environment variables: `env.*` functions from [`github.com/qdm12/gosettings/sources/env`](https://pkg.go.dev/github.com/qdm12/gosettings/sources/env)
- Minor feature notes:
  - No use of `reflect` for performance
  - Near zero dependency

## Philosophy

After having worked with Go and settings from different sources for years, I have come with a design I am happy with.

### Settings struct

Each component has a settings struct, where the zero value of a field should be **meaningless**.
For example, if the value `0` is allowed for a field, then it must be an `*int` field.
On the contrary, you could have an `int` field if the zero value `0` is meaningless.
The reasoning behind this is that you want the zero Go value to be considered as 'unset field' so that the field value can be defaulted, merged with or overridden by another settings struct. See the below interface comments for more details on what this allows.

Next, each of your settings struct should *ideally* implement the following interface:

```go
type Settings interface {
 // SetDefaults sets default values for all unset fields.
 // All pointer fields must be defaulted to a non nil value.
 // Usage:
 // - Once on the base settings at the start of the program.
 // - If the user requests a reset of the settings, on an empty settings struct.
 SetDefaults()
 // Validate validates all the settings and return an error if any field value is invalid.
 // It should only be called after `SetDefaults()` is called, and therefore should assume
 // all pointer fields are set and NOT nil.
 // Usage:
 // - Validate settings early at program start
 // - Validate new settings given, after calling .Copy() + .OverrideWith(newSettings)
 Validate() (err error)
 // Copy deep copies all the settings to a new Settings object.
 // Usage:
 // - Copy settings before modifying them with OverrideWith(), to validate them with Validate() before actually using them.
 Copy() Settings
 // MergeWith sets all the unset fields of the receiver to the values of the given settings.
 // Usage:
 // - Read from different settings sources with an order of precedence
 MergeWith(other Settings)
 // OverrideWith sets all the set values of the other settings to the fields of the receiver settings.
 // Usage:
 // - Update settings at runtime
 OverrideWith(other Settings)
 // ToLinesNode returns a (tree) node with the settings as lines, for displaying settings
 // in a formatted tree, where you can nest settings node to display a full settings tree.
 ToLinesNode() *gotree.Node
 // String returns the string representation of the settings.
 // It should simply return `s.ToLinesNode().String()` to show a tree of settings.
 String() string
}
```

💁 This is my recommendation, and obviously you don't need to:

- define this interface
- have all these methods exported
- define `ToLinesNode` with [gotree](https://github.com/qdm12/gotree) if you don't want to

➡️ [**Example settings implementation**](examples/settings/settings.go)

More concrete settings implementation examples using this library are notably:

- [Gluetun](https://github.com/qdm12/gluetun/tree/master/internal/configuration)
- [qdm12/dns](https://github.com/qdm12/dns/tree/v2.0.0-beta/internal/config)

### Settings methods usage

In the following Go examples, we use the [example settings implementation](examples/settings/settings.go).

#### Read settings from multiple sources

🚧 To be completed 🚧

You can check [these 10 lines](https://github.com/qdm12/gluetun/blob/a4c80b3045e65afbf86de44c89ad18deca51a43f/internal/configuration/sources/mux/reader.go#L31-L44) from [Gluetun](https://github.com/qdm12/gluetun) for a concrete example.

#### Updating settings at runtime

🚧 To be completed 🚧

## FAQ

- Why the `github.com/qdm12/govalid` dependency? **It is used in the environment variables helpers [`sources/env`](sources/env)**.
