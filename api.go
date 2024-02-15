package main

import (
	"log"
	"net/http"
)

func apiGet(data []byte) {
	s := &Server{data}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(s.Data)
	if err != nil {
		return
	}
}
