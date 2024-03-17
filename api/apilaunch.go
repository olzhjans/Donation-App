package api

import (
	auth2 "awesomeProject1/adminrights/adminauth"
	donationsubscribe2 "awesomeProject1/adminrights/donationsubscribe"
	"awesomeProject1/adminrights/editadmin"
	needs2 "awesomeProject1/adminrights/needs"
	orphanage2 "awesomeProject1/adminrights/orphanage"
	"awesomeProject1/auth"
	"awesomeProject1/comments"
	"awesomeProject1/donation"
	"awesomeProject1/donationsubscribe"
	"awesomeProject1/edituser"
	"awesomeProject1/needs"
	"awesomeProject1/wherespent"
	"fmt"
	"github.com/golang/glog"
	"net/http"
)

func LaunchApiServer() {
	//USER RIGHTS
	http.HandleFunc("/login", auth.UserLogin)
	http.HandleFunc("/userSignUp", auth.UserSignUp)
	http.HandleFunc("/showWhereSpent", wherespent.ShowWhereSpent)
	http.HandleFunc("/showNeeds", needs.ShowNeeds) //localhost:8080/showNeeds?orphanageid=ENTER_ID
	http.HandleFunc("/getNeedsByRegionAndType", needs.GetNeedsByRegionAndType)
	http.HandleFunc("/getComments", comments.GetComments)
	http.HandleFunc("/addComment", comments.AddComment)
	http.HandleFunc("/deleteComment", comments.DeleteComment) //DELETE ЗАПРОС localhost:8080/deleteComment?_id=ENTER_ID
	http.HandleFunc("/addDonate", donation.AddDonate)
	http.HandleFunc("/getTotalDonatedByOrphanageIdAndPeriod", donation.GetTotalDonatedByOrphanageIdAndPeriod)
	http.HandleFunc("/getTotalDonatedByUserIdAndPeriod", donation.GetTotalDonatedByUserIdAndPeriod)
	http.HandleFunc("/getDonationHistoryByUserId", donation.GetDonationHistoryByUserId) //GET localhost:8080/getDonationHistoryByUserId?userid=ENTER_ID
	http.HandleFunc("/addDonationSubscribe", donationsubscribe.SubscribeToDonation)
	http.HandleFunc("/getDonationSubscribeByUserId", donationsubscribe.GetDonationSubscribeByUserId) //GET localhost:8080/getDonationSubscribeByUserId?userid=ENTER_ID

	//ADMIN RIGHTS
	http.HandleFunc("/editAdmin", editadmin.EditAdmin)
	http.HandleFunc("/editUser", edituser.EditUser)
	http.HandleFunc("/addOrphanage", orphanage2.AddOrphanage)
	http.HandleFunc("/editOrphanage", orphanage2.EditOrphanage)
	http.HandleFunc("/getOrphanage", orphanage2.GetOrphanageByName) //localhost:8080/getOrphanage?name=ENTER_NAME
	http.HandleFunc("/addNeed", needs2.AddNeeds)
	http.HandleFunc("/activateNeed", needs2.ActivateNeedByNeedId)     //localhost:8080/activateNeed?needid=ENTER_ID
	http.HandleFunc("/deactivateNeed", needs2.DeactivateNeedByNeedId) //localhost:8080/deactivateNeed?needid=ENTER_ID
	http.HandleFunc("/adminSignUp", auth2.AdminSignUp)
	http.HandleFunc("/showWaitingList", auth2.ShowWaitingList)                                        //GET лист ожидающих подтверждения регистраций
	http.HandleFunc("/confirmRegistration", auth2.ConfirmRegistrationById)                            //localhost:8080/confirmRegistration?_id=ENTER_ID подтверждение регистрации по ID
	http.HandleFunc("/deleteWaitingList", auth2.DeleteWaitingListById)                                //localhost:8080/deleteWaitingList?_id=ENTER_ID удалить запись по ID
	http.HandleFunc("/activateDonateSubscription", donationsubscribe2.ActivateDonateSubscription)     //GET localhost:8080/activateDonateSubscription?_id=ENTER_ID
	http.HandleFunc("/deactivateDonateSubscription", donationsubscribe2.DeactivateDonateSubscription) //GET localhost:8080/deactivateDonateSubscription?_id=ENTER_ID

	//LAUNCH SERVER
	fmt.Println("Starting API server...")
	go func() {
		glog.Fatal(http.ListenAndServe(":8080", nil))
	}()
}
