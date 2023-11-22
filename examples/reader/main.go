package main

import (
	"fmt"

	"github.com/qdm12/gosettings/reader"
	"github.com/qdm12/gosettings/reader/sources/env"
	"github.com/qdm12/gosettings/reader/sources/flag"
)

func main() {
	flagSource := flag.New([]string{"program", "--key1=A"})
	envSource := env.New([]string{"KEY1=B", "KEY2=2"})
	reader := reader.New(reader.Settings{
		Sources: []reader.Source{flagSource, envSource},
	})

	value := reader.String("KEY1")
	// flag source takes precedence
	fmt.Println(value) // Prints "A"

	n, err := reader.Int("KEY2")
	if err != nil {
		panic(err)
	}
	// flag source has no value, so the environment
	// variable source is used.
	fmt.Println(n) // Prints "2"
}
