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

type User struct {
	Id          string `json:"-" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	MiddleName  string `json:"middle_name" db:"middle_name"`
	Class       string `json:"class" db:"class"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}
