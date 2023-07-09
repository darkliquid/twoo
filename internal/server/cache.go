package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"

	"github.com/spf13/afero"
)

const (
	cacheDirMode = 0755
)

func cache(cachedir string) func(http.Handler) http.Handler {
	fs := afero.NewBasePathFs(afero.NewOsFs(), cachedir)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get cache file name
			name := path.Join(r.URL.Path, "index.html")

			// Check if cached file exists and is not a directory
			if fi, err := fs.Stat(name); err == nil && !fi.IsDir() {
				w.WriteHeader(http.StatusOK)
				f, ferr := fs.Open(name)
				if ferr != nil {
					http.Error(w, ferr.Error(), http.StatusInternalServerError)
					return
				}
				defer f.Close()

				io.Copy(w, f) //nolint:errcheck // nothing we can do about this
				return
			}

			// Record the response so we can cache it
			rec := httptest.NewRecorder()
			next.ServeHTTP(rec, r)

			result := rec.Result() //nolint:bodyclose // this is a recorder
			statusCode := result.StatusCode
			value := rec.Body.Bytes()

			// As long we get a successful response, cache it
			if statusCode < http.StatusBadRequest {
				if err := fs.MkdirAll(r.URL.Path, cacheDirMode); err != nil {
					writeResponse(w, result.Header, statusCode, value)
					return
				}

				f, err := fs.Create(name)
				if err != nil {
					writeResponse(w, result.Header, statusCode, value)
					return
				}
				defer f.Close()

				f.Write(value) //nolint:errcheck // nothing we can do about this
			}
			writeResponse(w, result.Header, statusCode, value)
		})
	}
}

func writeResponse(w http.ResponseWriter, h http.Header, statusCode int, body []byte) {
	for k, v := range h {
		w.Header().Set(k, strings.Join(v, ","))
	}
	w.WriteHeader(statusCode)
	w.Write(body) //nolint:errcheck // nothing we can do about this
}
