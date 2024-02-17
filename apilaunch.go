package main

import (
	"log"
	"net/http"
)

func apiLaunch(data []byte) {
	s := &Server{data}
	http.Handle("/getOrphanageData", s)

	// Добавление обработчика для получения данных по ID
	http.HandleFunc("/getOrphanageByName", apiGetByName)

	// Установка обработчика POST запроса
	http.HandleFunc("/addNewOrphanage", addNewOrphanage)

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
