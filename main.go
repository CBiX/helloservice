package main

import (
	"io"
	"log"
	"net/http"

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
		// only supports one listener for now, for multiple listeners try Accept=true in the .service
		log.Fatal(http.Serve(listeners[0], nil))
	} else {
		log.Fatal(http.ListenAndServe(":8083", nil))
	}
}
