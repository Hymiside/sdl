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
	Id          string `json:"-,omitempty" db:"id"`
	Name        string `json:"name,omitempty" db:"name"`
	PhoneNumber string `json:"phone_number,omitempty" db:"phone_number"`
	Email       string `json:"email,omitempty" db:"email"`
	Password    string `json:"password,omitempty" db:"password_hash"`
}

type Student struct {
	Id          string `json:"-,omitempty" db:"id"`
	FirstName   string `json:"first_name,omitempty" db:"first_name"`
	LastName    string `json:"last_name,omitempty" db:"last_name"`
	MiddleName  string `json:"middle_name,omitempty" db:"middle_name"`
	ClassLet    string `json:"class_let,omitempty" db:"letter"`
	ClassNum    int    `json:"class_num,omitempty" db:"number"`
	SchoolId    string `json:"school_id,omitempty" db:"school_id"`
	Email       string `json:"email,omitempty" db:"email"`
	PhoneNumber string `json:"phone_number,omitempty" db:"phone_number"`
}

type Class struct {
	SchoolId string `json:"school_id,omitempty" db:"school_id"`
	LetClass string `json:"letter,omitempty" db:"letter"`
	NumClass int    `json:"number,omitempty" db:"number"`
}
