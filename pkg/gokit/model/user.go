package model

import "github.com/asaskevich/govalidator"

type User struct {
	ID        string `json:"userId"`
	FirstName string `json:"firstName" valid:"required"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" valid:"required"`
	Contact   string `json:"contact"`
	Password  string `json:"password" valid:"required"`
}

func ValidateUser(user *User) (bool, error) {
	result, errValidate := govalidator.ValidateStruct(user)
	if errValidate != nil {
		return false, errValidate
	}
	return result, nil
}
