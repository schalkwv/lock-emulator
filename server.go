package main

import (
	"fmt"
	"net/http"
)

func LockServer(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()
	router.Handle("/nuki", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
		w.WriteHeader(http.StatusOK)
	}))
}
