package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/afero"

	"github.com/darkliquid/twoo/internal/fs"
	"github.com/darkliquid/twoo/pkg/twitwoo"
)

// Serve listens on bind and serves a website generated from the data in data
// or from static files in cachedir, over HTTP.
func Serve(bind, cachedir string, srcfs afero.Fs) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.StripSlashes)
	if cachedir != "" {
		r.Use(cache)
	}

	data := twitwoo.New(srcfs)
	r.Get("/", handleTweets(data))
	r.Get("/page/{page:[1-9][0-9]*}", handleTweets(data))
	r.Get("/tweet/{id:[0-9]+}", handleTweet(data))
	r.Mount("/data/{type}_media", http.FileServer(http.FS(fs.AferoFS(srcfs))))

	ctx, cncl := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cncl()

	timeout := 10 * time.Second
	srv := &http.Server{
		Addr:              bind,
		Handler:           r,
		ReadTimeout:       timeout,
		ReadHeaderTimeout: timeout,
		WriteTimeout:      timeout,
		IdleTimeout:       timeout,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	ch := make(chan error)
	go func() { ch <- srv.ListenAndServe() }()

	select {
	case err := <-ch:
		if err == nil || errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	case <-ctx.Done():
		sctx, scncl := context.WithTimeout(context.Background(), timeout)
		defer scncl()
		return srv.Shutdown(sctx)
	}
}
