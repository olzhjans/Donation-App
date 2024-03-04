package structures

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginData struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Admins struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Password    string             `json:"password"`
	Email       string             `json:"email"`
	Region      string             `json:"region"`
	Firstname   string             `json:"firstname"`
	Lastname    string             `json:"lastname"`
	Phone       string             `json:"phone"`
	Who         string             `json:"who"`
	Id          string             `json:"id"`
	SignupDate  string             `json:"signup-date"`
	OrphanageId string             `json:"orphanage-id"`
}

type Users struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Password   string             `json:"password"`
	Email      string             `json:"email"`
	Region     string             `json:"region"`
	Firstname  string             `json:"firstname"`
	Lastname   string             `json:"lastname"`
	Phone      string             `json:"phone"`
	Donated    string             `json:"donated"`
	SignupDate string             `json:"signup-date"`
}

type Orphanage struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Region       string             `bson:"region" json:"region"`
	Address      string             `bson:"address" json:"address"`
	Description  string             `bson:"description" json:"description"`
	ChildsCount  string             `bson:"childscount" json:"childs-count"`
	WorkingHours string             `bson:"workinghours" json:"working-hours"`
	Photos       []string           `bson:"photos" json:"photos"`
	Bill         int64              `bson:"bill" json:"bill"`
}

type WhereSpent struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Date        primitive.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	SpentTo     string             `bson:"spent-to,omitempty" json:"spent-to,omitempty"`
	OrphanageId string             `bson:"orphanage-id,omitempty" json:"orphanage-id,omitempty"`
}

type Need struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Amount           string             `bson:"amount,omitempty" json:"amount,omitempty"`
	Expiring         primitive.DateTime `bson:"expiring,omitempty" json:"expiring,omitempty"`
	CategoryOfDonate string             `bson:"categoryofdonate,omitempty" json:"categoryofdonate,omitempty"`
	SizeOfClothes    string             `bson:"sizeofclothes,omitempty" json:"sizeofclothes,omitempty"`
	TypeOfCount      string             `bson:"typeofcount,omitempty" json:"typeofcount,omitempty"`
	TypeOfDonate     string             `bson:"typeofdonate,omitempty" json:"typeofdonate,omitempty"`
	OrphanageId      string             `bson:"orphanageid,omitempty" json:"orphanageid,omitempty"`
	IsActive         bool               `bson:"isactive,omitempty" json:"isactive,omitempty"`
}

type Commentary struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NeedId string             `bson:"need-id" json:"need-id"`
	UserId string             `bson:"user-id" json:"user-id"`
	Text   string             `bson:"text" json:"text"`
	Date   primitive.DateTime `bson:"date,omitempty" json:"date,omitempty"`
}

type Donate struct {
	UserId      string   `bson:"user-id" json:"user-id"`
	OrphanageId []string `bson:"orphanage-id" json:"orphanage-id"`
	Sum         int      `bson:"sum" json:"sum"`
}

type DonateSubscribe struct {
	ID            primitive.ObjectID
	OrphanageId   []string
	BankDetailsId string
	Amount        int64
	WhichDay      int32
	IsActive      bool
}

type BankDetails struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	Expiring   string             `bson:"expiring" json:"expiring"`
	Cvv        string             `bson:"cvv" json:"cvv"`
	CardNumber string             `bson:"cardnumber" json:"cardnumber"`
	UserId     string             `bson:"userid" json:"userid"`
	Bill       int64              `bson:"bill" json:"bill"`
}

type DonationHistory struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserId      string             `bson:"user-id,omitempty" json:"user-id,omitempty"`
	OrphanageId []string           `bson:"orphanage-id,omitempty" json:"orphanage-id,omitempty"`
	Sum         int                `bson:"sum,omitempty" json:"sum,omitempty"`
	Date        primitive.DateTime `bson:"date,omitempty" json:"date,omitempty"`
}

type DonationSubscribe struct {
	BankDetails BankDetails `bson:"bank-details,omitempty" json:"bank-details,omitempty"`
	//BankDetailsId string `bson:"bankdetailsid,omitempty" json:"bankdetailsid,omitempty"`
	OrphanageId []string `bson:"orphanageid" json:"orphanageid"`
	Amount      int64    `bson:"amount" json:"amount"`
	WhichDay    int8     `bson:"whichday" json:"whichday"`
	IsActive    bool     `bson:"isactive" json:"isactive"`
}

type DonateDeactivation struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`
}

type OrphanageFilter struct {
	Region           string `json:"region,omitempty"`
	CategoryOfDonate string `json:"category-of-donate,omitempty"`
}

type WhereSpentFilter struct {
	OrphanageId string             `json:"orphanage-id"`
	From        primitive.DateTime `json:"from"`
	To          primitive.DateTime `json:"to"`
}

type CommentaryFilter struct {
	NeedId string             `json:"need-id"`
	From   primitive.DateTime `json:"from"`
	To     primitive.DateTime `json:"to"`
}
type DonationFilter struct {
	Id   string             `json:"id"`
	From primitive.DateTime `json:"from"`
	To   primitive.DateTime `json:"to"`
}
