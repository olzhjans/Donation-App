package main

import (
	"awesomeProject1/auth"
	"awesomeProject1/comments"
	"awesomeProject1/edituser"
	"awesomeProject1/needs"
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
	http.HandleFunc("/confirmRegistration", auth.ConfirmRegistrationById) //localhost:8080/confirmRegistration?_id=ENTER_ID подтверждение регистрации по ID
	http.HandleFunc("/deleteWaitingList", auth.DeleteWaitingListById)     //localhost:8080/deleteWaitingList?_id=ENTER_ID удалить запись по ID

	http.HandleFunc("/editAdmin", edituser.EditAdmin)
	http.HandleFunc("/editUser", edituser.EditUser)

	http.HandleFunc("/addOrphanage", orphanage.AddOrphanage)
	http.HandleFunc("/editOrphanage", orphanage.EditOrphanage)

	http.HandleFunc("/getOrphanagesByRegionAndNeeds", orphanage.GetOrphanagesByRegionAndNeeds)
	//Обработчик GET запроса
	http.HandleFunc("/getOrphanage", orphanage.GetOrphanageByName) //localhost:8080/getOrphanage?name=ENTER_NAME

	http.HandleFunc("/showWhereSpent", wherespent.ShowWhereSpent)
	http.HandleFunc("/showNeeds", needs.ShowNeeds) //localhost:8080/showNeeds?orphanageid=ENTER_ID
	http.HandleFunc("/addNeeds", needs.AddNeeds)   //localhost:8080/showNeeds?orphanageid=ENTER_ID

	http.HandleFunc("/getComments", comments.GetComments)
	http.HandleFunc("/addComment", comments.AddComment)
	http.HandleFunc("/deleteComment", comments.DeleteComment) //DELETE ЗАПРОС localhost:8080/deleteComment?_id=ENTER_ID

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
