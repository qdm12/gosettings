# Gosettings

`gosettings` is a Go package providing **helper functions for working with settings**.

Add it to your Go project with:

```sh
go get github.com/qdm12/gosettings
```

Features:

- Define settings struct methods:
  - `Copy`: [`github.com/qdm12/gosettings/copier`](https://pkg.go.dev/github.com/qdm12/gosettings/copier)
  - `SetDefaults`: [`github.com/qdm12/gosettings/defaults`](https://pkg.go.dev/github.com/qdm12/gosettings/defaults)
  - `MergeWith`: [`github.com/qdm12/gosettings/merge`](https://pkg.go.dev/github.com/qdm12/gosettings/merge)
  - `OverrideWith`: [`github.com/qdm12/gosettings/override`](https://pkg.go.dev/github.com/qdm12/gosettings/override)
  - `Validate`: [`github.com/qdm12/gosettings/validate`](https://pkg.go.dev/github.com/qdm12/gosettings/validate)
- Reading settings from sources:
  - Environment variables: [`github.com/qdm12/gosettings/sources/env`](https://pkg.go.dev/github.com/qdm12/gosettings/sources/env)

## Philosophy

After having worked with Go and settings from different sources for years, I have come with a design I am happy with.

### Settings struct

Each component has a settings struct, where the zero value of a field should be **meaningless**.
For example, if the value `0` is allowed for a field, then it must be an `*int` field.
On the contrary, you could have an `int` field if the zero value `0` is not valid.
The reasoning behind this is that you want the zero Go value to be considered as 'unset field' so that the field value can be defaulted, merged with or overridden by another settings struct.

Next, each of your settings struct should implement the following interface:

```go
type Settings interface {
 // SetDefaults sets default values for all unset fields.
 // All nilable fields must be set to something not nil.
 // Usage:
 // - Once on the base settings at the start of the program.
 // - If the user requests a reset of the settings, on a zeroed settings struct.
 SetDefaults()
 // Validate validates all the settings and return an error if any field value is invalid.
 // It should only be called after SetDefaults() is called.
 // Usage:
 // - Validate settings early at program start
 // - Validate new settings given, after calling .Copy() + .OverrideWith(newSettings)
 Validate() (err error)
 // Copy deep copies all the settings to a new Settings object.
 // Usage:
 // - Copy settings before modifying them with OverrideWith() or MergeWith(), to validate them with Validate().
 Copy() Settings
 // MergeWith sets all the unset fields of the receiver to the values of the given settings.
 // Usage:
 // - Read from different settings source with an order of precedence
 MergeWith(other Settings)
 // OverrideWith sets all the set values of the other settings to the fields of the receiver settings.
 // Usage:
 // - Write to different settings destination with an order of precedence
 OverrideWith(other Settings)
 // ToLinesNode returns a (tree) node with the settings as lines, for displaying settings
 // in a formatted tree, where you can nest settings node to display a full settings tree.
 ToLinesNode() *gotree.Node
 // String returns the string representation of the settings.
 // It should return `s.ToLinesNode().String()`.
 String() string
}
```

💁 you don't need to define this interface

💁 you don't need to have all these methods exported

💁 you don't need to define `ToLinesNode` with [gotree](https://github.com/qdm12/gotree) if you don't want to

➡️ [**Example settings implementation**](examples/settings/settings.go)

### Settings methods usage

In the following Go examples, we use the [example settings implementation](examples/settings/settings.go).

#### Read settings from multiple sources

🚧 To be completed 🚧

You can check [these 10 lines](https://github.com/qdm12/gluetun/blob/a4c80b3045e65afbf86de44c89ad18deca51a43f/internal/configuration/sources/mux/reader.go#L31-L44) from [Gluetun](https://github.com/qdm12/gluetun) for a concrete example.

#### Updating settings at runtime

🚧 To be completed 🚧

## FAQ

- Why the `github.com/qdm12/govalid` dependency? **It is used in the environment variables helpers [`sources/envhelpers`](sources/envhelpers)**.
- Why is it still Go 1.17 💤?? **Because many linters still don't support Go 1.18**, so several projects are still using Go 1.17.
