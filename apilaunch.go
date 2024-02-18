package main

import (
	"log"
	"net/http"
)

func apiLaunch(data []byte) {
	// Добавление обработчика для получения данных по ID
	http.HandleFunc("/getOrphanageByName", apiGetOrphanageDataByName)

	// Установка обработчика POST запроса
	http.HandleFunc("/addNewOrphanage", apiAddNewOrphanage)

	// Обработчик POST регистрация пользователя
	http.HandleFunc("/userSignUp", apiUserSignUp)

	// Обработчик POST регистрация пользователя
	http.HandleFunc("/adminSignUp", apiAdminSignUp)

	// Обработчик POST регистрация пользователя
	http.HandleFunc("/login", apiUserLogin)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(s.Data)
	if err != nil {
		return
	}
}*/
