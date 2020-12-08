package main

import (
	"net/http"
	"fmt"
	"os"
	"context"
	"errors"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

type Handler struct {
	server *http.Server
	name   string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "server %s process request.", h.name)
}

func main() {
	eg, ctx := errgroup.WithContext(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig)

	ch := make(chan error, 1)
	server := &Handler{
		server: &http.Server{},
		name:   "server",
	}

	eg.Go(func() error {
		if err := server.server.ListenAndServe(); err != nil {
			fmt.Printf("server have error, err=%+v", err)
		}
		close(ch)
		return errors.New("server have error")
	})

	eg.Go(func() error {
		var err error
		select {
		case <-sig:
			err = errors.New("signal error")
		case <-ch:
			err = errors.New("server error")
		case <-ctx.Done():
			err = errors.New("context deadline")
		}

		signal.Stop(sig)

		server.server.Close()

		return err
	})

	eg.Wait()

}
