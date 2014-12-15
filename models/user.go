package models

import (
	"time"
)

type User struct {
	Id       int64
	Username string    `xorm:"UNIQUE NOT NULL"`
	Email    string    `xorm:"UNIQUE NOT NULL"`
	Password string    `xorm:"NOT NULL"`
	Created  time.Time `xorm:"CREATED"`
	Updated  time.Time `xorm:"UPDATED"`
}
