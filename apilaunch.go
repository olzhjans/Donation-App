package main

import (
	"awesomeProject1/auth"
	"awesomeProject1/edituser"
	"awesomeProject1/orphanage"
	"fmt"
	"log"
	"net/http"
)

func apiLaunch() {
	http.HandleFunc("/login", auth.UserLogin)

	http.HandleFunc("/userSignUp", auth.UserSignUp)
	http.HandleFunc("/adminSignUp", auth.AdminSignUp)

	http.HandleFunc("/editAdmin", edituser.EditAdmin)
	http.HandleFunc("/editUser", edituser.EditUser)

	http.HandleFunc("/addOrphanage", orphanage.AddOrphanage)
	http.HandleFunc("/editOrphanage", orphanage.EditOrphanage)

	http.HandleFunc("/getOrphanagesByRegionAndNeeds", orphanage.GetOrphanagesByRegionAndNeeds)
	//Обработчик GET запроса
	http.HandleFunc("/getOrphanage", orphanage.GetOrphanageByName) //localhost:8080/getOrphanage?name=ENTER_NAME

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
