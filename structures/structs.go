package structures

import "go.mongodb.org/mongo-driver/bson/primitive"

type Orphanage struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string             `json:"name"`
	Region       string             `json:"region"`
	Address      string             `json:"address"`
	Description  string             `json:"description"`
	ChildsCount  string             `json:"childs-count"`
	WorkingHours string             `json:"working-hours"`
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
