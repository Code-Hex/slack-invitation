package main

import (
	"net/http"

	"github.com/Code-Hex/slack-invitaion/index"
	"github.com/Code-Hex/slack-invitaion/invite"
)

func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", index.Handler)
	mux.HandleFunc("/invite", invite.Handler)
	http.ListenAndServe(":3000", mux)
}
