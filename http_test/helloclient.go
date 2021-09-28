package main

import (
	"net/http"
)

func HelloClient(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello client"))
}
