package merge

import "net/http"

func HTTPHandler(existing, other http.Handler) (result http.Handler) {
	if existing != nil {
		return existing
	}
	return other
}
