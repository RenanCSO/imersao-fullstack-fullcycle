package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// User is a user of PIXCode
type User struct {
	ID        string     `json:"id" valid:"uuid"`
	Name      string     `json:"name" valid:"notnull"`
	Email     string     `json:"email" valid:"notnull"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

// NewUser is a function to create Users
func NewUser(name string, email string) (*User, error) {
	user := User {
		Name: name,
		Email: email,
	}

	user.ID = uuid.NewV4().String()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}