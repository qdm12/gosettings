package env

type Option func(s *settings)

func ForceLowercase(lowercase bool) Option {
	return func(s *settings) {
		s.forceLowercase = &lowercase
	}
}
