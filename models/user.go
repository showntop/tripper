package models

import (
	// "strings"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Username       string `json:"username"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Mobile         string `json:"mobile"`
	Password       string `json:"password" sql:"-"`
	HashedPassword string `json:"password" sql:"hashed_password"`
}

func (u *User) Validate() error {
	if len(u.Username) > 10 {
		return errors.New("用户名长度超限")
	}
	return nil
}

func (u *User) EncryptPassword() error {
	// u.HashedPassword
	p, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.HashedPassword = string(p)
	return nil
}

func (u *User) Authenticate() error {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(u.Password))
}
