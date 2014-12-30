package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRole int

const (
	GUEST UserRole = iota
	REGUSER
	EDITOR
	ADMIN
	ROOT
)

type User struct {
	Id       int64
	Username string    `xorm:"varchar(25) not null unique"`
	Email    string    `xorm:"not null unique"`
	Password string    `xorm:"not null"`
	Role     UserRole  `xorm:"unique"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

func (u *User) EncodePassword() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 6)
	if err != nil {
		u.Password = fmt.Sprintf("%x", hashedPassword)
	} else {
		fmt.Errorf("EncodePassword error %v", err)
	}
}

func (u *User) ValidatePassword(password string) bool {
	newUser := &User{Password: password}
	newUser.EncodePassword()
	return u.Password == newUser.Password
}

func CreateUser(u *User) error {
	u.EncodePassword()
	return nil
}
