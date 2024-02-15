package main

type Server struct {
	Data []byte
}

type Orphanage struct {
	//ID           primitive.ObjectID `bson:"_id"`
	Name         string
	Region       string
	Address      string
	Description  string
	ChildsCount  string
	WorkingHours string
}

type Users struct {
	Password   string
	Email      string
	Region     string
	Firstname  string
	Lastname   string
	Phone      string
	Donated    string
	SignupDate string
}

type Admins struct {
	Password    string
	Email       string
	Region      string
	Firstname   string
	Lastname    string
	Phone       string
	Who         string
	Id          string
	SignupDate  string
	OrphanageId string
}
