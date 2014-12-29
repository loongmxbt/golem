package models

import (
	"time"
)

type User struct {
	Id       int64
	Username string    `xorm:"varchar(25) not null unique"`
	Email    string    `xorm:"not null unique"`
	Password string    `xorm:"not null"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}
