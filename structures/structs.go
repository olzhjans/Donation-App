package structures

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	Date        string             `bson:"date,omitempty" json:"date,omitempty"`
	SpentTo     string             `bson:"spent-to,omitempty" json:"spent-to,omitempty"`
	OrphanageId string             `bson:"orphanage-id,omitempty" json:"orphanage-id,omitempty"`
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

type LoginData struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type OrphanageFilter struct {
	Region           string `json:"region,omitempty"`
	CategoryOfDonate string `json:"category-of-donate,omitempty"`
}
