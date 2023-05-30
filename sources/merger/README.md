# Merger settings source

The `Merger` settings source is a generic settings source that can be used to merge multiple settings sources together, with an order of precedence.

If you have two settings sources `env` (environment variables) and `flags` (cli flags), you can use the `Merger` settings source to merge the settings read from them together, with the `Flags` settings source having a higher precedence than the `Env` settings source:

```go
package main

import (
  "github.com/qdm12/gosettings/sources/merger"
)

type Settings struct {
  // ...
}

// ... and Settings methods

func main() {
  env := NewEnv()
  flags := NewFlags()
  merger := merger.New[Settings](flags, env)
  settings, err := merger.Read()
  if err != nil {
    panic(err)
  }
  settings.SetDefaults()

  // ...
}
```

💁 A runnable example is available at [examples/merger/main.go](../../examples/merger/main.go)
