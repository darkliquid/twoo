package server

import "net/http"

func cache(next http.Handler) http.Handler {
	return next
}
