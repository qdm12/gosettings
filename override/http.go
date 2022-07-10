package override

import "net/http"

func WithHTTPHandler(existing, other http.Handler) (result http.Handler) {
	if other != nil {
		return other
	}
	return existing
}
