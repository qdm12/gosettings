package main

import (
	"fmt"

	"github.com/qdm12/gosettings/reader"
	"github.com/qdm12/gosettings/reader/sources/env"
)

func main() {
	sourceA := env.New([]string{"KEY1=A1"})
	sourceB := env.New([]string{"KEY1=B1", "KEY2=2"})
	reader := reader.New(reader.Settings{
		Sources: []reader.Source{sourceA, sourceB},
	})

	value := reader.String("KEY1")
	fmt.Println(value) // A1 - source A takes precedence

	n, err := reader.Int("KEY2")
	if err != nil {
		panic(err)
	}
	fmt.Println(n) // 2 - source A has no value, so source B is used.
}
