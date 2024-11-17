package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id          uint   `json:"user_id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    []byte `json:"-"`
	Role        string `json:"user_role"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
