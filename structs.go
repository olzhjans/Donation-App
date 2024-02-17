package main

type Server struct {
	Data []byte
}

type Orphanage struct {
	Name         string `json:"name"`
	Region       string `json:"region"`
	Address      string `json:"address"`
	Description  string `json:"description"`
	ChildsCount  string `json:"childs-count"`
	WorkingHours string `json:"working-hours"`
}

type Users struct {
	Password   string `json:"password"`
	Email      string `json:"email"`
	Region     string `json:"region"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Phone      string `json:"phone"`
	Donated    string `json:"donated"`
	SignupDate string `json:"signup-date"`
}

type Admins struct {
	Password    string `json:"password"`
	Email       string `json:"email"`
	Region      string `json:"region"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Phone       string `json:"phone"`
	Who         string `json:"who"`
	Id          string `json:"id"`
	SignupDate  string `json:"signup-date"`
	OrphanageId string `json:"orphanage-id"`
}
