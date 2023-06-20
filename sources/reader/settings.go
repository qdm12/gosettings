package reader

import "github.com/qdm12/gosettings/sources/parse"

type settings struct {
	forceLowercase *bool
	acceptEmpty    *bool
	retroKeys      []string
}

func (r *Reader) makeParseOptions(options []Option) (parseOptions []parse.Option) {
	var settings settings
	for _, option := range options {
		option(&settings)
	}

	const maxOptions = 3
	parseOptions = make([]parse.Option, 0, maxOptions)
	if settings.forceLowercase != nil {
		parseOption := parse.ForceLowercase(*settings.forceLowercase)
		parseOptions = append(parseOptions, parseOption)
	}
	if settings.acceptEmpty != nil {
		parseOption := parse.AcceptEmpty(*settings.acceptEmpty)
		parseOptions = append(parseOptions, parseOption)
	}
	if len(settings.retroKeys) > 0 {
		parseOption := parse.RetroKeys(r.handleDeprecatedKey, settings.retroKeys...)
		parseOptions = append(parseOptions, parseOption)
	}

	return parseOptions
}
