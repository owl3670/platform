package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() {
	print("Serve")
	http.ListenAndServe(":4000", mux.NewRouter())
}
