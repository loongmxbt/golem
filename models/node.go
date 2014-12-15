package models

import (
	"time"
)

type Node struct {
	Id      int64
	Title   string    `xorm:"NOT NULL"`
	Author  string    `xorm:"NOT NULL"`
	Content string    `xorm:"NOT NULL"`
	Created time.Time `xorm:"CREATED"`
	Updated time.Time `xorm:"UPDATED"`
}
