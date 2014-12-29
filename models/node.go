package models

import (
	"time"
)

type Node struct {
	Id      int64
	Title   string    `xorm:"varchar(50) not null"`
	Author  string    `xorm:"varchar(25) not null"`
	Content string    `xorm:"not null"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
