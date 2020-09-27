package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/coreos/go-systemd/v22/activation"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world\n")
}

func main() {
	hh := new(HelloHandler)
	http.Handle("/", hh)

	listeners, err := activation.Listeners()

	if err == nil && len(listeners) >= 1 {
		wg := new(sync.WaitGroup)
		wg.Add(len(listeners))
		for _, l := range listeners {
			go func(listener net.Listener) {
				log.Println("listening in goroutine")
				log.Fatal(http.Serve(listener, nil))
				wg.Done()
			}(l)
		}
		wg.Wait()
	} else {
		log.Fatal(http.ListenAndServe(":8083", nil))
	}
}
