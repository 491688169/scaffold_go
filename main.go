package main

import (
	"hello/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
