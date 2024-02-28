package structures

import "go.mongodb.org/mongo-driver/bson/primitive"

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
}

type Commentary struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NeedId string             `bson:"need-id" json:"need-id"`
	UserId string             `bson:"user-id" json:"user-id"`
	Text   string             `bson:"text" json:"text"`
	Date   primitive.DateTime `bson:"date,omitempty" json:"date,omitempty"`
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
