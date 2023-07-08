package server

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/darkliquid/twoo/internal/website"
	"github.com/darkliquid/twoo/pkg/twitwoo"
)

const pageSize = 20

func handleTweets(data *twitwoo.Data) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := chi.URLParam(r, "page")
		page, _ := strconv.ParseInt(pageStr, 10, 64)
		if err := website.Index(data, page, pageSize, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func handleTweet(data *twitwoo.Data) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, _ := strconv.ParseInt(idStr, 10, 64)
		if err := website.Page(data, id, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
