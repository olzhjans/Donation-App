package main

import (
	"awesomeProject1/auth"
	"awesomeProject1/edituser"
	"awesomeProject1/orphanage"
	"awesomeProject1/wherespent"
	"fmt"
	"log"
	"net/http"
)

func apiLaunch() {
	http.HandleFunc("/login", auth.UserLogin)
	http.HandleFunc("/userSignUp", auth.UserSignUp)
	http.HandleFunc("/adminSignUp", auth.AdminSignUp)
	http.HandleFunc("/showWaitingList", auth.ShowWaitingList)             //GET лист ожидающих подтверждения регистраций
	http.HandleFunc("/confirmRegistration", auth.ConfirmRegistrationById) //localhost:8080/confirmRegistration?id=ENTER_ID подтверждение регистрации по ID
	http.HandleFunc("/deleteWaitingList", auth.DeleteWaitingListById)     //localhost:8080/deleteWaitingList?id=ENTER_ID удалить запись по ID

	http.HandleFunc("/editAdmin", edituser.EditAdmin)
	http.HandleFunc("/editUser", edituser.EditUser)

	http.HandleFunc("/addOrphanage", orphanage.AddOrphanage)
	http.HandleFunc("/editOrphanage", orphanage.EditOrphanage)

	http.HandleFunc("/getOrphanagesByRegionAndNeeds", orphanage.GetOrphanagesByRegionAndNeeds)
	//Обработчик GET запроса
	http.HandleFunc("/getOrphanage", orphanage.GetOrphanageByName) //localhost:8080/getOrphanage?name=ENTER_NAME

	http.HandleFunc("/showWhereSpent", wherespent.ShowWhereSpent)

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
