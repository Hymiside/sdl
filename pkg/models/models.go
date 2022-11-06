package models

type ConfigServer struct {
	Port string
	Host string
}

type ConfigRepository struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type School struct {
	Id          string `json:"-" db:"id"`
	Name        string `json:"name" db:"name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password_hash"`
}

type Student struct {
	Id          string `json:"-" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	MiddleName  string `json:"middle_name" db:"middle_name"`
	ClassId     string `json:"class_id" db:"class_id"`
	SchoolId    string `json:"school_id" db:"school_id"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}

type Class struct {
	SchoolId string `json:"school_id" db:"school_id"`
	LetClass string `json:"letter" db:"letter"`
	NumClass int    `json:"number" db:"number"`
}
